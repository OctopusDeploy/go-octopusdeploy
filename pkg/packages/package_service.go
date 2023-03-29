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
	"regexp"
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
	items := []*Package{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
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
func Upload(client newclient.Client, spaceID string, fileName string, reader io.Reader, overwriteMode OverwriteMode) (*PackageUploadResponse, bool, error) {
	// directly uploading a file only requires a forward-moving `io.Reader`.
	//
	// UploadV2 requires `io.ReadSeeker` because delta upload needs it.
	// If we pass useDeltaCompression: false to UploadV2 then it promises not to
	// call Seek, so we can preserve our existing `reader` contract here by faking it out

	return UploadV2(client, spaceID, fileName, &fakeReadSeeker{reader: reader}, overwriteMode, false)
}

func UploadV2(client newclient.Client, spaceID string, fileName string, reader io.ReadSeeker, overwriteMode OverwriteMode, useDeltaCompression bool) (*PackageUploadResponse, bool, error) {
	if client == nil {
		return nil, false, internal.CreateRequiredParameterIsEmptyOrNilError("client")
	}
	if spaceID == "" {
		return nil, false, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if fileName == "" {
		return nil, false, internal.CreateRequiredParameterIsEmptyOrNilError("fileName")
	}
	if reader == nil {
		return nil, false, internal.CreateRequiredParameterIsEmptyOrNilError("reader")
	}

	if useDeltaCompression {
		fileLength, err := reader.Seek(0, io.SeekEnd)
		if err != nil {
			return nil, false, err
		}
		_, err = reader.Seek(0, io.SeekStart)
		if err != nil {
			return nil, false, err
		}
		response, createdNewPackage, err, errIsRecoverable := attemptDeltaPush(client, spaceID, fileName, reader, fileLength, overwriteMode)

		// If the delta upload was good, we are all done here
		if err == nil {
			return response, createdNewPackage, err
		}

		if !errIsRecoverable {
			// we should log the error, but we don't have access to a logger
			return nil, false, err
		}

		// at this point we lose the error from attemptDeltaPush, but we don't do anything with it anyway

		// we can recover by pushing the full file
		// need to seek back to the start or the regular forward-read will fail
		_, err = reader.Seek(0, io.SeekStart)
		if err != nil {
			return nil, false, err
		}
	}

	// push complete package to server
	params := map[string]any{
		"spaceId": spaceID,
	}
	if overwriteMode != "" {
		params["overwriteMode"] = overwriteMode
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.PackageUpload, params)
	if err != nil {
		return nil, false, err
	}

	return httpUploadPackageFile(client, expandedUri, fileName, reader)
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

// ParsePackageIDAndVersion ported from OctopusServer's PackageIdentity class
// See PackageIdentityParser in the C# Octopus Client SDK
func ParsePackageIDAndVersion(fileName string) (packageID string, version string, err error) {
	pattern := regexp.MustCompile(
		"^" + // start of line
			"(?P<packageId>(\\w+([_.-]\\w+)*?))" + // Package ID
			"\\." + // version separator
			"(?P<semanticVersion>(\\d+(\\.\\d+){0,3}" + // Major Minor Patch
			"(-[0-9A-Za-z-]+(\\.[0-9A-Za-z-]+)*)?)" + // Pre-release identifiers
			"(\\+[0-9A-Za-z-]+(\\.[0-9A-Za-z-]+)*)?)" + // Build Metadata
			"$") // EOL

	match := pattern.FindStringSubmatch(fileName)
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
