package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type workerService struct {
	discoverWorkerPath   string
	operatingSystemsPath string
	shellsPath           string

	canDeleteService
}

func newWorkerService(sling *sling.Sling, uriTemplate string, discoverWorkerPath string, operatingSystemsPath string, shellsPath string) *workerService {
	workerService := &workerService{
		discoverWorkerPath:   discoverWorkerPath,
		operatingSystemsPath: operatingSystemsPath,
		shellsPath:           shellsPath,
	}
	workerService.service = newService(ServiceWorkerService, sling, uriTemplate)

	return workerService
}

func (s workerService) getPagedResponse(path string) ([]*Worker, error) {
	resources := []*Worker{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Workers), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Workers)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new worker.
func (s workerService) Add(worker *Worker) (*Worker, error) {
	if worker == nil {
		return nil, createInvalidParameterError("Add", ParameterWorker)
	}

	path, err := getAddPath(s, worker)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), worker, new(Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Worker), nil
}

// TODO: validation implementation

// func (s workerService) DiscoverWorker() ([]string, error) {
// 	resp, err := apiGet(s.getClient(), new([]string), s.discoverWorkerPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := resp.(*[]string)
// 	return *response, nil
// }

// GetAll returns all workers. If none can be found or an error occurs, it
// returns an empty collection.
func (s workerService) GetAll() ([]*Worker, error) {
	items := []*Worker{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the worker that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s workerService) GetByID(id string) (*Worker, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Worker), nil
}

// GetByIDs returns the workers that match the input IDs.
func (s workerService) GetByIDs(ids []string) ([]*Worker, error) {
	if len(ids) == 0 {
		return []*Worker{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*Worker{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName returns the workers with a matching partial name.
func (s workerService) GetByName(name string) ([]*Worker, error) {
	path, err := getByNamePath(s, name)
	if err != nil {
		return []*Worker{}, err
	}

	return s.getPagedResponse(path)
}

// GetByPartialName performs a lookup and returns enironments with a matching
// partial name.
func (s workerService) GetByPartialName(name string) ([]*Worker, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Worker{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies an worker based on the one provided as input.
func (s workerService) Update(worker *Worker) (*Worker, error) {
	if worker == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterWorker)
	}

	path, err := getUpdatePath(s, worker)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), worker, new(Worker), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Worker), nil
}

func (s workerService) GetWorkerOperatingSystems() ([]string, error) {
	resp, err := apiGet(s.getClient(), new([]string), s.operatingSystemsPath)
	if err != nil {
		return nil, err
	}

	response := resp.(*[]string)
	return *response, nil
}

func (s workerService) GetWorkerShells() ([]string, error) {
	resp, err := apiGet(s.getClient(), new([]string), s.shellsPath)
	if err != nil {
		return nil, err
	}

	response := resp.(*[]string)
	return *response, nil
}
