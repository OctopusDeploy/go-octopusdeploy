package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type workerService struct {
	discoverWorkerPath   string
	operatingSystemsPath string
	shellsPath           string

	service
}

func newWorkerService(sling *sling.Sling, uriTemplate string, discoverWorkerPath string, operatingSystemsPath string, shellsPath string) *workerService {
	workerService := &workerService{
		discoverWorkerPath:   discoverWorkerPath,
		operatingSystemsPath: operatingSystemsPath,
		shellsPath:           shellsPath,
	}
	workerService.service = newService(serviceWorkerService, sling, uriTemplate, new(model.Worker))

	return workerService
}

func (s workerService) getPagedResponse(path string) ([]*model.Worker, error) {
	resources := []*model.Worker{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Workers), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Workers)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new worker.
func (s workerService) Add(worker *model.Worker) (*model.Worker, error) {
	if worker == nil {
		return nil, createInvalidParameterError("Add", parameterWorker)
	}

	path, err := getAddPath(s, worker)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), worker, new(model.Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Worker), nil
}

// GetAll returns all workers. If none can be found or an error occurs, it
// returns an empty collection.
func (s workerService) GetAll() ([]*model.Worker, error) {
	items := []*model.Worker{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the worker that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s workerService) GetByID(id string) (*model.Worker, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Worker), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.Worker), nil
}

// GetByIDs returns the workers that match the input IDs.
func (s workerService) GetByIDs(ids []string) ([]*model.Worker, error) {
	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*model.Worker{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName returns the workers with a matching partial name.
func (s workerService) GetByName(name string) ([]*model.Worker, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*model.Worker{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns enironments with a matching
// partial name.
func (s workerService) GetByPartialName(name string) ([]*model.Worker, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.Worker{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an worker based on the one provided as input.
func (s workerService) Update(worker *model.Worker) (*model.Worker, error) {
	if worker == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterWorker)
	}

	path, err := getUpdatePath(s, worker)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), worker, new(model.Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Worker), nil
}
