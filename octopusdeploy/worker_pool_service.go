package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type workerPoolService struct {
	dynamicWorkerTypesPath string
	sortOrderPath          string
	summaryPath            string
	supportedTypesPath     string

	services.canDeleteService
}

func newWorkerPoolService(sling *sling.Sling, uriTemplate string, dynamicWorkerTypesPath string, sortOrderPath string, summaryPath string, supportedTypesPath string) *workerPoolService {
	workerPoolService := &workerPoolService{
		dynamicWorkerTypesPath: dynamicWorkerTypesPath,
		sortOrderPath:          sortOrderPath,
		summaryPath:            summaryPath,
		supportedTypesPath:     supportedTypesPath,
	}
	workerPoolService.service = services.newService(ServiceWorkerPoolService, sling, uriTemplate)

	return workerPoolService
}

// Add creates a new worker pool.
func (s workerPoolService) Add(workerPool IWorkerPool) (IWorkerPool, error) {
	if workerPool == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterWorkerPool)
	}

	workerPoolResource, err := ToWorkerPoolResource(workerPool)
	if err != nil {
		return nil, err
	}

	response, err := apiAdd(s.getClient(), workerPoolResource, new(WorkerPoolResource), s.BasePath)
	if err != nil {
		return nil, err
	}

	return ToWorkerPool(response.(*WorkerPoolResource))
}

// Get returns a collection of worker pools based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s workerPoolService) Get(workerPoolsQuery WorkerPoolsQuery) (*WorkerPools, error) {
	v, _ := query.Values(workerPoolsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := apiGet(s.getClient(), new(WorkerPoolResources), path)
	if err != nil {
		return &WorkerPools{}, err
	}

	return ToWorkerPools(response.(*WorkerPoolResources)), nil
}

// GetAll returns all worker pools. If none can be found or an error occurs, it
// returns an empty collection.
func (s *workerPoolService) GetAll() ([]IWorkerPool, error) {
	items := []*WorkerPoolResource{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
	return ToWorkerPoolArray(items), err
}

// GetByID returns the worker pool that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s workerPoolService) GetByID(id string) (IWorkerPool, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(WorkerPoolResource), path)
	if err != nil {
		return nil, err
	}

	return ToWorkerPool(resp.(*WorkerPoolResource))
}

// Update modifies a worker pool based on the one provided as input.
func (s workerPoolService) Update(workerPool IWorkerPool) (IWorkerPool, error) {
	if workerPool == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterWorkerPool)
	}

	path, err := getUpdatePath(s, workerPool)
	if err != nil {
		return nil, err
	}

	workerPoolResource, err := ToWorkerPoolResource(workerPool)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), workerPoolResource, new(WorkerPoolResource), path)
	if err != nil {
		return nil, err
	}

	return ToWorkerPool(resp.(*WorkerPoolResource))
}
