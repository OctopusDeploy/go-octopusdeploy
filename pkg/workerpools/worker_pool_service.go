package workerpools

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
	"strings"
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

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*WorkerPoolResource]), path)
	if err != nil {
		return &WorkerPools{}, err
	}

	return ToWorkerPools(response.(*resources.Resources[*WorkerPoolResource])), nil
}

// GetAll returns all worker pools. If none can be found or an error occurs, it
// returns an empty collection.
func (s *WorkerPoolService) GetAll() ([]*WorkerPoolListResult, error) {
	items := []*WorkerPoolListResult{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return nil, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
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

	resp, err := api.ApiGet(s.GetClient(), new(WorkerPoolResource), path)
	if err != nil {
		return nil, err
	}

	return ToWorkerPool(resp.(*WorkerPoolResource))
}

func (s *WorkerPoolService) GetByName(name string) ([]IWorkerPool, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	path, err := services.GetByNamePath(s, name)
	if err != nil {
		return []IWorkerPool{}, err
	}

	response, err := services.GetPagedResponse[WorkerPoolResource](s, path)
	return ToWorkerPoolArray(response), nil
}

func (s *WorkerPoolService) GetByIdentifier(identifier string) (IWorkerPool, error) {
	workerPool, err := s.GetByID(identifier)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if workerPool != nil {
			return workerPool, nil
		}
	}

	possibleWorkerPools, err := s.GetByName(identifier)
	if err != nil {
		return nil, err
	}

	for _, w := range possibleWorkerPools {
		if strings.EqualFold(identifier, w.GetName()) {
			return w, nil
		}
	}

	// this is a workaround as the api does not support querying via the worker pool slug
	allWorkerPools, err := s.GetAll()
	for _, pool := range allWorkerPools {
		if strings.EqualFold(identifier, pool.Slug) {
			workerPool, err := s.GetByID(pool.ID)
			if err != nil {
				return nil, err
			}
			return workerPool, nil
		}
	}

	return nil, fmt.Errorf("cannot find worker pool with identifier of '%s'", identifier)
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

func (s *WorkerPoolService) GetWorkers(workerPool IWorkerPool) ([]*machines.Worker, error) {
	if workerPool == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterWorkerPool)
	}

	uriTemplate, err := uritemplates.Parse(workerPool.GetLinks()[constants.LinkWorkers])
	if err != nil {
		return nil, err
	}

	values := make(map[string]interface{})
	path, err := uriTemplate.Expand(values)
	if err != nil {
		return nil, err
	}

	return services.GetPagedResponse[machines.Worker](s, path)
}
