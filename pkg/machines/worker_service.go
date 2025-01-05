package machines

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
	"strings"
)

type WorkerService struct {
	discoverWorkerPath   string
	operatingSystemsPath string
	shellsPath           string

	services.CanDeleteService
}

func NewWorkerService(sling *sling.Sling, uriTemplate string, discoverWorkerPath string, operatingSystemsPath string, shellsPath string) *WorkerService {
	return &WorkerService{
		discoverWorkerPath:   discoverWorkerPath,
		operatingSystemsPath: operatingSystemsPath,
		shellsPath:           shellsPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceWorkerService, sling, uriTemplate),
		},
	}
}

// Get returns a collection of machines based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *WorkerService) Get(workersQuery WorkersQuery) (*resources.Resources[*Worker], error) {
	path, err := s.GetURITemplate().Expand(workersQuery)
	if err != nil {
		return &resources.Resources[*Worker]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Worker]), path)
	if err != nil {
		return &resources.Resources[*Worker]{}, err
	}

	return response.(*resources.Resources[*Worker]), nil
}

// Add creates a new worker.
//
// Deprecated: use workers.Add
func (s *WorkerService) Add(worker *Worker) (*Worker, error) {
	if IsNil(worker) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterWorker)
	}

	path, err := services.GetAddPath(s, worker)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), worker, new(Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Worker), nil
}

// GetAll returns all workers. If none can be found or an error occurs, it
// returns an empty collection.
//
// Deprecated: use workers.GetAll
func (s *WorkerService) GetAll() ([]*Worker, error) {
	items := []*Worker{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the worker that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: use workers.GetByID
func (s *WorkerService) GetByID(id string) (*Worker, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Worker), nil
}

// GetByIDs returns the workers that match the input IDs.
func (s *WorkerService) GetByIDs(ids []string) ([]*Worker, error) {
	if len(ids) == 0 {
		return []*Worker{}, nil
	}

	path, err := services.GetByIDsPath(s, ids)
	if err != nil {
		return []*Worker{}, err
	}

	return services.GetPagedResponse[Worker](s, path)
}

// GetByName returns the workers with a matching partial name.
func (s *WorkerService) GetByName(name string) ([]*Worker, error) {
	if internal.IsEmpty(name) {
		return []*Worker{}, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName)
	}

	path, err := services.GetByNamePath(s, name)
	if err != nil {
		return []*Worker{}, err
	}

	return services.GetPagedResponse[Worker](s, path)
}

// GetByIdentifier returns the worker with a matching ID or name.
func (s *WorkerService) GetByIdentifier(identifier string) (*Worker, error) {
	worker, err := s.GetByID(identifier)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if worker != nil {
			return worker, nil
		}
	}

	possibleWorkers, err := s.GetByName(identifier)
	if err != nil {
		return nil, err
	}

	for _, w := range possibleWorkers {
		if strings.EqualFold(identifier, w.Name) {
			return w, nil
		}
	}

	return nil, fmt.Errorf("cannot find worker with name or ID of '%s'", identifier)
}

// GetByPartialName performs a lookup and returns environments with a matching
// partial name.
func (s *WorkerService) GetByPartialName(partialName string) ([]*Worker, error) {
	if internal.IsEmpty(partialName) {
		return []*Worker{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*Worker{}, err
	}

	return services.GetPagedResponse[Worker](s, path)
}

// Update modifies a worker based on the one provided as input.
//
// Deprecated: use workers.Update
func (s *WorkerService) Update(worker *Worker) (*Worker, error) {
	if worker == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterWorker)
	}

	path, err := services.GetUpdatePath(s, worker)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), worker, new(Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Worker), nil
}

func (s *WorkerService) GetWorkerOperatingSystems() ([]string, error) {
	resp, err := api.ApiGet(s.GetClient(), new([]string), s.operatingSystemsPath)
	if err != nil {
		return nil, err
	}

	response := resp.(*[]string)
	return *response, nil
}

func (s *WorkerService) GetWorkerShells() ([]string, error) {
	resp, err := api.ApiGet(s.GetClient(), new([]string), s.shellsPath)
	if err != nil {
		return nil, err
	}

	response := resp.(*[]string)
	return *response, nil
}
