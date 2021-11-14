package octopusdeploy

import (
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
	"github.com/jinzhu/copier"
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
	workerPoolService.service = newService(ServiceWorkerPoolService, sling, uriTemplate)

	return workerPoolService
}

func toWorkerPool(workerPoolResource *WorkerPoolResource) (IWorkerPool, error) {
	if isNil(workerPoolResource) {
		return nil, createInvalidParameterError("toWorkerPool", ParameterWorkerPoolResource)
	}

	var workerPool IWorkerPool
	var err error
	switch workerPoolResource.GetWorkerPoolType() {
	case WorkerPoolTypeDynamic:
		workerPool, err = NewDynamicWorkerPool(workerPoolResource.GetName(), workerPoolResource.WorkerType)
		if err != nil {
			return nil, err
		}
	case WorkerPoolTypeStatic:
		workerPool, err = NewStaticWorkerPool(workerPoolResource.GetName())
		if err != nil {
			return nil, err
		}
	}

	err = copier.Copy(workerPool, workerPoolResource)
	if err != nil {
		return nil, err
	}

	return workerPool, nil
}

func toWorkerPools(workerPoolResources *WorkerPoolResources) *WorkerPools {
	return &WorkerPools{
		Items:        toWorkerPoolArray(workerPoolResources.Items),
		PagedResults: workerPoolResources.PagedResults,
	}
}

func toWorkerPoolArray(workerPoolResources []*WorkerPoolResource) []IWorkerPool {
	items := []IWorkerPool{}
	for _, workerPoolResource := range workerPoolResources {
		workerPool, err := toWorkerPool(workerPoolResource)
		if err != nil {
			return nil
		}
		items = append(items, workerPool)
	}
	return items
}

func toWorkerPoolResource(workerPool IWorkerPool) (*WorkerPoolResource, error) {
	if isNil(workerPool) {
		return nil, createInvalidParameterError("toWorkerPoolResource", ParameterWorkerPool)
	}

	workerPoolResource := newWorkerPoolResource(workerPool.GetName())
	workerPoolResource.WorkerPoolType = workerPool.GetWorkerPoolType()

	err := copier.Copy(&workerPoolResource, workerPool)
	if err != nil {
		return nil, err
	}

	return workerPoolResource, nil
}

// Add creates a new worker pool.
func (s workerPoolService) Add(workerPool IWorkerPool) (IWorkerPool, error) {
	if workerPool == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterWorkerPool)
	}

	workerPoolResource, err := toWorkerPoolResource(workerPool)
	if err != nil {
		return nil, err
	}

	response, err := apiAdd(s.getClient(), workerPoolResource, new(WorkerPoolResource), s.BasePath)
	if err != nil {
		return nil, err
	}

	return toWorkerPool(response.(*WorkerPoolResource))
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

	return toWorkerPools(response.(*WorkerPoolResources)), nil
}

// GetAll returns all worker pools. If none can be found or an error occurs, it
// returns an empty collection.
func (s *workerPoolService) GetAll() ([]IWorkerPool, error) {
	items := []*WorkerPoolResource{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
	return toWorkerPoolArray(items), err
}

// GetByID returns the worker pool that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s workerPoolService) GetByID(id string) (IWorkerPool, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError(OperationGetByID, ParameterID)
	}

	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(WorkerPoolResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(IWorkerPool), nil
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

	var workerPoolResource interface{}
	switch workerPool.GetWorkerPoolType() {
	case "DynamicWorkerPool":
		workerPoolResource = new(DynamicWorkerPool)
	case "StaticWorkerPool":
		workerPoolResource = new(StaticWorkerPool)
	}

	resp, err := apiUpdate(s.getClient(), workerPool, workerPoolResource, path)
	if err != nil {
		return nil, err
	}

	return resp.(IWorkerPool), nil
}
