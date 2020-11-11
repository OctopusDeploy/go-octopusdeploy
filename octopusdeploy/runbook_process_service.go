package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type runbookProcessService struct {
	service
}

func newRunbookProcessService(sling *sling.Sling, uriTemplate string) *runbookProcessService {
	return &runbookProcessService{
		service: newService(ServiceRunbookProcessService, sling, uriTemplate),
	}
}

func (s runbookProcessService) getPagedResponse(path string) ([]*RunbookProcess, error) {
	resources := []*RunbookProcess{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(RunbookProcesses), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*RunbookProcesses)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetAll returns all runbook processes. If none can be found or an error
// occurs, it returns an empty collection.
func (s runbookProcessService) GetAll() ([]*RunbookProcess, error) {
	path, err := getPath(s)
	if err != nil {
		return []*RunbookProcess{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the runbook process that matches the input ID. If one cannot
// be found, it returns nil and an error.
func (s runbookProcessService) GetByID(id string) (*RunbookProcess, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(RunbookProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookProcess), nil
}
