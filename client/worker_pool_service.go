package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type workerPoolService struct {
	dynamicWorkerTypesPath string
	sortOrderPath          string
	summaryPath            string
	supportedTypesPath     string

	service
}

func newWorkerPoolService(sling *sling.Sling, uriTemplate string, dynamicWorkerTypesPath string, sortOrderPath string, summaryPath string, supportedTypesPath string) *workerPoolService {
	workerPoolService := &workerPoolService{
		dynamicWorkerTypesPath: dynamicWorkerTypesPath,
		sortOrderPath:          sortOrderPath,
		summaryPath:            summaryPath,
		supportedTypesPath:     supportedTypesPath,
	}
	workerPoolService.service = newService(serviceWorkerPoolService, sling, uriTemplate, new(model.WorkerPool))

	return workerPoolService
}

func toWorkerPoolArray(workerPools []model.IWorkerPool) []model.IWorkerPool {
	items := []model.IWorkerPool{}
	for _, workerPool := range workerPools {
		items = append(items, workerPool)
	}
	return items
}

func (s workerPoolService) getPagedResponse(path string) ([]*model.WorkerPool, error) {
	resources := []*model.WorkerPool{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.WorkerPools), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.WorkerPools)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new worker pool.
func (s workerPoolService) Add(workerPool model.IWorkerPool) (model.IWorkerPool, error) {
	if workerPool == nil {
		return nil, createInvalidParameterError("Add", parameterWorkerPool)
	}

	path, err := getAddPath(s, workerPool)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), workerPool, new(model.WorkerPool), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.WorkerPool), nil
}

// GetAll returns all worker pools. If none can be found or an error occurs, it
// returns an empty collection.
func (s workerPoolService) GetAll() ([]*model.WorkerPool, error) {
	items := []*model.WorkerPool{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the worker pool that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s workerPoolService) GetByID(id string) (*model.WorkerPool, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.WorkerPool), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.WorkerPool), nil
}

// GetByIDs returns the worker pools that match the input IDs.
func (s workerPoolService) GetByIDs(ids []string) ([]*model.WorkerPool, error) {
	if len(ids) == 0 {
		return []*model.WorkerPool{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*model.WorkerPool{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName returns the worker pools with a matching partial name.
func (s workerPoolService) GetByName(name string) ([]*model.WorkerPool, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*model.WorkerPool{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns worker pools with a matching
// partial name.
func (s workerPoolService) GetByPartialName(name string) ([]*model.WorkerPool, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*model.WorkerPool{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a worker pool based on the one provided as input.
func (s workerPoolService) Update(workerPool *model.WorkerPool) (*model.WorkerPool, error) {
	if workerPool == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterWorkerPool)
	}

	path, err := getUpdatePath(s, workerPool)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), workerPool, new(model.WorkerPool), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.WorkerPool), nil
}
