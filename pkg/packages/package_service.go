package packages

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
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

func Upload(client newclient.Client, command *PackageUploadCommand) (*PackageUploadResponse, bool, error) {
	multipartWriter := NewMultipartFileStreamingReader(command.FileName, command.FileReader)

	path, err := client.URITemplateCache().Expand(uritemplates.PackageUpload, command)
	if err != nil {
		return nil, false, err
	}

	req, err := http.NewRequest(http.MethodPost, path, multipartWriter)
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
	if resp.StatusCode == 201 || resp.StatusCode == 200 {
		outputResponseBody := new(PackageUploadResponse)
		err = bodyDecoder.Decode(outputResponseBody)
		if err != nil {
			return nil, false, err
		}
		// the server returns 201 if it created a new file, 200 if it ignored an existing file
		createdNewFile := resp.StatusCode == 201
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
