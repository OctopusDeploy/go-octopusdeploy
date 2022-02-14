package services

import (
	"github.com/dghubble/sling"
)

type packageService struct {
	bulkPath           string
	deltaSignaturePath string
	deltaUploadPath    string
	notesListPath      string
	uploadPath         string

	canDeleteService
}

func newPackageService(sling *sling.Sling, uriTemplate string, deltaSignaturePath string, deltaUploadPath string, notesListPath string, bulkPath string, uploadPath string) *packageService {
	packageService := &packageService{
		bulkPath:           bulkPath,
		deltaSignaturePath: deltaSignaturePath,
		deltaUploadPath:    deltaUploadPath,
		notesListPath:      notesListPath,
		uploadPath:         uploadPath,
	}
	packageService.service = newService(ServicePackageService, sling, uriTemplate)

	return packageService
}

func (s packageService) getPagedResponse(path string) ([]*Package, error) {
	resources := []*Package{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Packages), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Packages)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetAll returns all packages. If none can be found or an error occurs, it
// returns an empty collection.
func (s packageService) GetAll() ([]*Package, error) {
	path, err := getPath(s)
	if err != nil {
		return []*Package{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the package that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s packageService) GetByID(id string) (*Package, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Package), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Package), nil
}

// Update modifies a package based on the one provided as input.
func (s packageService) Update(octopusPackage *Package) (*Package, error) {
	if octopusPackage == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterPackage)
	}

	path, err := getUpdatePath(s, octopusPackage)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), octopusPackage, new(Package), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Package), nil
}
