package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type runbookProcessService struct {
	service
}

func newRunbookProcessService(sling *sling.Sling, uriTemplate string) *runbookProcessService {
	return &runbookProcessService{
		service: newService(serviceRunbookProcessService, sling, uriTemplate, nil),
	}
}

func (s runbookProcessService) getPagedResponse(path string) ([]*model.RunbookProcess, error) {
	resources := []*model.RunbookProcess{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.RunbookProcesses), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.RunbookProcesses)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetAll returns all runbook processes. If none can be found or an error
// occurs, it returns an empty collection.
func (s runbookProcessService) GetAll() ([]*model.RunbookProcess, error) {
	path, err := getPath(s)
	if err != nil {
		return []*model.RunbookProcess{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the runbook process that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s runbookProcessService) GetByID(id string) (*model.RunbookProcess, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.RunbookProcess), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*model.RunbookProcess), nil
}
