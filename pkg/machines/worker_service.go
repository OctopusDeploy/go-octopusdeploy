package machines

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
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

func (s *WorkerService) getPagedResponse(path string) ([]*Worker, error) {
	resources := []*Worker{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(Workers), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Workers)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new worker.
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

// TODO: validation implementation

// func (s workerService) DiscoverWorker() ([]string, error) {
// 	resp, err := services.ApiGet(s.GetClient(), new([]string), s.discoverWorkerPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := resp.(*[]string)
// 	return *response, nil
// }

// GetAll returns all workers. If none can be found or an error occurs, it
// returns an empty collection.
func (s *WorkerService) GetAll() ([]*Worker, error) {
	items := []*Worker{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the worker that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *WorkerService) GetByID(id string) (*Worker, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(Worker), path)
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

	return s.getPagedResponse(path)
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

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns enironments with a matching
// partial name.
func (s *WorkerService) GetByPartialName(partialName string) ([]*Worker, error) {
	if internal.IsEmpty(partialName) {
		return []*Worker{}, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName)
	}

	path, err := services.GetByPartialNamePath(s, partialName)
	if err != nil {
		return []*Worker{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an worker based on the one provided as input.
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
	resp, err := services.ApiGet(s.GetClient(), new([]string), s.operatingSystemsPath)
	if err != nil {
		return nil, err
	}

	response := resp.(*[]string)
	return *response, nil
}

func (s *WorkerService) GetWorkerShells() ([]string, error) {
	resp, err := services.ApiGet(s.GetClient(), new([]string), s.shellsPath)
	if err != nil {
		return nil, err
	}

	response := resp.(*[]string)
	return *response, nil
}
