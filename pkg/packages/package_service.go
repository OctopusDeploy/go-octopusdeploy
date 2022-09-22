package packages

import (
	"bytes"
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
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

func Upload(client newclient.Client, command *PackageUploadCommand) (*PackageUploadResponse, error) {
	// TODO implement upload streaming once tests are in place
	file, _ := os.Open(command.FileName)
	defer func() { _ = file.Close() }()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer func() { _ = writer.Close() }()
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	_, err := io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	path, err := client.URITemplateCache().Expand(uritemplates.PackageUpload, command)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := client.HttpSession().DoRawRequest(req)
	if err != nil {
		return nil, err
	}
	defer newclient.CloseResponse(resp)

	bodyDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode == 201 {
		outputResponseBody := new(PackageUploadResponse)
		err = bodyDecoder.Decode(outputResponseBody)
		if err != nil {
			return nil, err
		}
		return outputResponseBody, nil
	} else {
		outputResponseError := new(core.APIError)
		err = bodyDecoder.Decode(outputResponseError)
		if err != nil {
			return nil, err
		}
		return nil, outputResponseError
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
