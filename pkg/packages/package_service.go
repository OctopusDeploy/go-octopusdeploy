package packages

import (
	"encoding/json"
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

	multipartWriter := NewMultipartFileStreamingReader(fileName, reader)

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

	req, err := http.NewRequest(http.MethodPost, expandedUri, multipartWriter)
	if err != nil {
		return nil, false, err
	}
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())

	resp, err := client.HttpSession().DoRawRequest(req)
	if err != nil {
		return nil, false, err
	}
	defer newclient.CloseResponse(resp)

	bodyDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		outputResponseBody := new(PackageUploadResponse)
		err = bodyDecoder.Decode(outputResponseBody)
		if err != nil {
			return nil, false, err
		}
		// the server returns 201 if it created a new file, 200 if it ignored an existing file
		createdNewFile := resp.StatusCode == http.StatusCreated
		return outputResponseBody, createdNewFile, nil
	} else {
		outputResponseError := new(core.APIError)
		err = bodyDecoder.Decode(outputResponseError)
		if err != nil {
			return nil, false, err
		}
		return nil, false, outputResponseError
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
