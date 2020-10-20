package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type workerPoolService struct {
	dynamicWorkerTypesPath string
	sortOrderPath          string
	summaryPath            string
	supportedTypesPath     string

	canDeleteService
}

func newWorkerPoolService(sling *sling.Sling, uriTemplate string, dynamicWorkerTypesPath string, sortOrderPath string, summaryPath string, supportedTypesPath string) *workerPoolService {
	workerPoolService := &workerPoolService{
		dynamicWorkerTypesPath: dynamicWorkerTypesPath,
		sortOrderPath:          sortOrderPath,
		summaryPath:            summaryPath,
		supportedTypesPath:     supportedTypesPath,
	}
	workerPoolService.service = newService(serviceWorkerPoolService, sling, uriTemplate, new(WorkerPool))

	return workerPoolService
}

func toWorkerPoolArray(workerPools []IWorkerPool) []IWorkerPool {
	items := []IWorkerPool{}
	for _, workerPool := range workerPools {
		items = append(items, workerPool)
	}
	return items
}

func (s workerPoolService) getPagedResponse(path string) ([]*WorkerPool, error) {
	resources := []*WorkerPool{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(WorkerPools), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*WorkerPools)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new worker pool.
func (s workerPoolService) Add(workerPool IWorkerPool) (IWorkerPool, error) {
	if workerPool == nil {
		return nil, createInvalidParameterError("Add", parameterWorkerPool)
	}

	path, err := getAddPath(s, workerPool)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), workerPool, new(WorkerPool), path)
	if err != nil {
		return nil, err
	}

	return resp.(*WorkerPool), nil
}

// GetAll returns all worker pools. If none can be found or an error occurs, it
// returns an empty collection.
func (s workerPoolService) GetAll() ([]*WorkerPool, error) {
	items := []*WorkerPool{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the worker pool that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s workerPoolService) GetByID(id string) (*WorkerPool, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*WorkerPool), nil
}

// GetByIDs returns the worker pools that match the input IDs.
func (s workerPoolService) GetByIDs(ids []string) ([]*WorkerPool, error) {
	if len(ids) == 0 {
		return []*WorkerPool{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*WorkerPool{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName returns the worker pools with a matching partial name.
func (s workerPoolService) GetByName(name string) ([]*WorkerPool, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*WorkerPool{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns worker pools with a matching
// partial name.
func (s workerPoolService) GetByPartialName(name string) ([]*WorkerPool, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*WorkerPool{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a worker pool based on the one provided as input.
func (s workerPoolService) Update(workerPool *WorkerPool) (*WorkerPool, error) {
	if workerPool == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterWorkerPool)
	}

	path, err := getUpdatePath(s, workerPool)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), workerPool, new(WorkerPool), path)
	if err != nil {
		return nil, err
	}

	return resp.(*WorkerPool), nil
}
