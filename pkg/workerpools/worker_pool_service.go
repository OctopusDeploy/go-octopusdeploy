package workerpools

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type WorkerPoolService struct {
	dynamicWorkerTypesPath string
	sortOrderPath          string
	summaryPath            string
	supportedTypesPath     string

	services.CanDeleteService
}

func NewWorkerPoolService(sling *sling.Sling, uriTemplate string, dynamicWorkerTypesPath string, sortOrderPath string, summaryPath string, supportedTypesPath string) *WorkerPoolService {
	return &WorkerPoolService{
		dynamicWorkerTypesPath: dynamicWorkerTypesPath,
		sortOrderPath:          sortOrderPath,
		summaryPath:            summaryPath,
		supportedTypesPath:     supportedTypesPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceWorkerPoolService, sling, uriTemplate),
		},
	}
}

// Add creates a new worker pool.
func (s *WorkerPoolService) Add(workerPool IWorkerPool) (IWorkerPool, error) {
	if IsNil(workerPool) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterWorkerPool)
	}

	workerPoolResource, err := ToWorkerPoolResource(workerPool)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiAdd(s.GetClient(), workerPoolResource, new(WorkerPoolResource), s.BasePath)
	if err != nil {
		return nil, err
	}

	return ToWorkerPool(response.(*WorkerPoolResource))
}

// Get returns a collection of worker pools based on the criteria defined by
// its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func (s *WorkerPoolService) Get(workerPoolsQuery WorkerPoolsQuery) (*WorkerPools, error) {
	v, _ := query.Values(workerPoolsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := services.ApiGet(s.GetClient(), new(WorkerPoolResources), path)
	if err != nil {
		return &WorkerPools{}, err
	}

	return ToWorkerPools(response.(*WorkerPoolResources)), nil
}

// GetAll returns all worker pools. If none can be found or an error occurs, it
// returns an empty collection.
func (s *WorkerPoolService) GetAll() ([]IWorkerPool, error) {
	items := []*WorkerPoolResource{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return ToWorkerPoolArray(items), err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return ToWorkerPoolArray(items), err
}

// GetByID returns the worker pool that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *WorkerPoolService) GetByID(id string) (IWorkerPool, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(WorkerPoolResource), path)
	if err != nil {
		return nil, err
	}

	return ToWorkerPool(resp.(*WorkerPoolResource))
}

// Update modifies a worker pool based on the one provided as input.
func (s *WorkerPoolService) Update(workerPool IWorkerPool) (IWorkerPool, error) {
	if workerPool == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterWorkerPool)
	}

	path, err := services.GetUpdatePath(s, workerPool)
	if err != nil {
		return nil, err
	}

	workerPoolResource, err := ToWorkerPoolResource(workerPool)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), workerPoolResource, new(WorkerPoolResource), path)
	if err != nil {
		return nil, err
	}

	return ToWorkerPool(resp.(*WorkerPoolResource))
}
