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

// GetAll returns all runbook processes. If none can be found or an error
// occurs, it returns an empty collection.
func (s runbookProcessService) GetAll() ([]*RunbookProcess, error) {
	items := []*RunbookProcess{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
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
