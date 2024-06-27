package packages

import (
	"encoding/json"
	"errors"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type PackageService struct {
	bulkPath           string
	deltaSignaturePath string
	deltaUploadPath    string
	notesListPath      string
	uploadPath         string

	services.CanDeleteService
}

func NewPackageService(sling *sling.Sling, uriTemplate string, deltaSignaturePath string, deltaUploadPath string, notesListPath string, bulkPath string, uploadPath string) *PackageService {
	return &PackageService{
		bulkPath:           bulkPath,
		deltaSignaturePath: deltaSignaturePath,
		deltaUploadPath:    deltaUploadPath,
		notesListPath:      notesListPath,
		uploadPath:         uploadPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServicePackageService, sling, uriTemplate),
		},
	}
}

// GetAll returns all packages. If none can be found or an error occurs, it
// returns an empty collection.
func (s *PackageService) GetAll() ([]*Package, error) {
	path := s.GetBasePath() // to get all packages we just hit /api/Spaces-NN/packages

	var packages []*Package

	loadNextPage := true
	for loadNextPage {
		resp, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Package]), path)

		if err != nil {
			return packages, err
		}

		r := resp.(*resources.Resources[*Package])
		packages = append(packages, r.Items...)
		path, loadNextPage = services.LoadNextPage(r.PagedResults)
	}

	return packages, nil
}

// GetByID returns the package that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *PackageService) GetByID(id string) (*Package, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Package), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Package), nil
}

// Update modifies a package based on the one provided as input.
func (s *PackageService) Update(octopusPackage *Package) (*Package, error) {
	if octopusPackage == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "octopusPackage")
	}

	path, err := services.GetUpdatePath(s, octopusPackage)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), octopusPackage, new(Package), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Package), nil
}

// ----------------------

// Upload uploads a package to the octopus server's builtin package feed.
// Parameters:
// - client: The API client reference
// - spaceID: ID of the octopus space to work within
// - fileName: The string which we tell the server to use for the file name (may not necessarily be an actual filename on disk)
// - reader: io.Reader which provides the binary file data to upload
// - overwriteMode: Instructs the server what to do in the case that the package already exists.
//
// Return values:
// - PackageUploadResponse: The server's response to the upload request
// - bool: True if the server created a new file, false if it ignored an existing file
// - error: Any error that occurred during the upload process
func Upload(client newclient.Client, spaceID string, fileName string, reader io.Reader, overwriteMode OverwriteMode) (*PackageUploadResponse, bool, error) {
	// directly uploading a file only requires a forward-moving `io.Reader`.
	//
	// UploadV2 requires `io.ReadSeeker` because delta upload needs it.
	// If we pass useDeltaCompression: false to UploadV2 then it promises not to
	// call Seek, so we can preserve our existing `reader` contract here by faking it out

	v2Response, err := UploadV2(client, spaceID, fileName, &fakeReadSeeker{reader: reader}, overwriteMode, false)
	if err != nil {
		return nil, false, err
	}
	return &v2Response.PackageUploadResponse, v2Response.CreatedNewFile, nil
}

func UploadV2(client newclient.Client, spaceID string, fileName string, reader io.ReadSeeker, overwriteMode OverwriteMode, useDeltaCompression bool) (*PackageUploadResponseV2, error) {
	if client == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("client")
	}
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if fileName == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("fileName")
	}
	if reader == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("reader")
	}

	if useDeltaCompression {
		fileLength, err := reader.Seek(0, io.SeekEnd)
		if err != nil {
			return nil, err
		}
		_, err = reader.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}
		return uploadDelta(client, spaceID, fileName, reader, fileLength, overwriteMode)
	} else {
		// push complete package to server
		params := map[string]any{
			"spaceId": spaceID,
		}
		if overwriteMode != "" {
			params["overwriteMode"] = overwriteMode
		}

		expandedUri, err := client.URITemplateCache().Expand(uritemplates.PackageUpload, params)
		if err != nil {
			return nil, err
		}

		stdResponse, createdNewFile, err := httpUploadPackageFile(client, expandedUri, fileName, reader)
		if err != nil {
			return nil, err
		}

		return &PackageUploadResponseV2{
			CreatedNewFile:        createdNewFile,
			UploadMethod:          UploadMethodStandard,
			PackageUploadResponse: *stdResponse,
			UploadInfo:            nil,
		}, nil

	}
}

// httpUploadFile creates a multipart POST to upload a file and sends it to the server for either a full package upload
// or a delta upload. Both have the same response type so we deserialize that too
func httpUploadPackageFile(client newclient.Client, url string, fileName string, reader io.Reader) (response *PackageUploadResponse, createdNewFile bool, err error) {

	multipartWriter := NewMultipartFileStreamingReader(fileName, reader)

	req, err := http.NewRequest(http.MethodPost, url, multipartWriter)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())

	resp, err := client.HttpSession().DoRawRequest(req)
	if err != nil {
		return
	}
	defer newclient.CloseResponse(resp)

	bodyDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		outputResponseBody := new(PackageUploadResponse)
		err = bodyDecoder.Decode(outputResponseBody)
		if err != nil {
			return
		}
		// the server returns 201 if it created a new file, 200 if it ignored an existing file
		createdNewFile = resp.StatusCode == http.StatusCreated
		response = outputResponseBody
		return
	} else {
		outputResponseError := new(core.APIError)
		err = bodyDecoder.Decode(outputResponseError)
		if err != nil {
			return
		}
		err = outputResponseError
		return
	}
}

// List returns a list of packages from the server, in a standard Octopus paginated result structure.
// If you don't specify --limit the server will use a default limit (typically 30)
func List(client newclient.Client, spaceID string, filter string, limit int) (*resources.Resources[*Package], error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	templateParams := map[string]any{"spaceId": spaceID}
	if filter != "" {
		templateParams["filter"] = filter
	}
	if limit > 0 {
		templateParams["take"] = limit
	}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.Packages, templateParams)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return newclient.Get[resources.Resources[*Package]](client.HttpSession(), expandedUri)
}

// ---- internal helpers -----

func withoutLastExtension(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[:i]
		}
	}
	return path
}

// conventional wisdom in go is that the file extension is always the thing after the last .
// but this mishandles .tar.gz files, so we need a special case to handle them (and .tar.bz2 etc)
func removeExtension(fileName string) string {
	candidate := withoutLastExtension(fileName)
	if len(candidate) > 3 && strings.EqualFold(filepath.Ext(candidate), ".tar") {
		return withoutLastExtension(candidate)
	}
	return candidate
}

// ParsePackageIDAndVersion ported from OctopusServer's PackageIdentity class
// See PackageIdentityParser in the C# Octopus Client SDK
// Note: Unlike in C#, fileName here includes the file extension
func ParsePackageIDAndVersion(fileName string) (packageID string, version string, err error) {
	pattern := regexp.MustCompile(
		"^" + // start of line
			"(?P<packageId>(\\w+([_.-]\\w+)*?))" + // Package ID
			"\\." + // version separator
			"(?P<semanticVersion>(\\d+(\\.\\d+){0,3}" + // Major Minor Patch
			"(-[0-9A-Za-z-]+(\\.[0-9A-Za-z-]+)*)?)" + // Pre-release identifiers
			"(\\+[0-9A-Za-z-]+(\\.[0-9A-Za-z-]+)*)?)" + // Build Metadata
			"$") // EOL

	match := pattern.FindStringSubmatch(removeExtension(fileName))
	if match == nil {
		err = errors.New("could not determine the package ID and/or version based on the supplied filename")
		return
	}
	for i, name := range pattern.SubexpNames() {
		if name == "packageId" {
			packageID = match[i]
		} else if name == "semanticVersion" {
			version = match[i]
		}
	}
	return
}

type fakeReadSeeker struct {
	reader io.Reader
}

func (f *fakeReadSeeker) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *fakeReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, errors.New("seek is not supported for fakeReadSeeker")
}
