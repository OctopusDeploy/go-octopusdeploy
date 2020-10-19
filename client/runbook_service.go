package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type runbookService struct {
	service
}

func newRunbookService(sling *sling.Sling, uriTemplate string) *runbookService {
	runbookService := &runbookService{}
	runbookService.service = newService(serviceRunbookService, sling, uriTemplate, new(model.Runbook))

	return runbookService
}

// GetAll returns all runbooks. If none can be found or an error occurs, it
// returns an empty collection.
func (s runbookService) GetAll() ([]*model.Runbook, error) {
	items := []*model.Runbook{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the runbook that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s runbookService) GetByID(id string) (*model.Runbook, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Runbook), path)
	if err != nil {
		return nil, createResourceNotFoundError("runbook", "ID", id)
	}

	return resp.(*model.Runbook), nil
}

// Add returns the runbook that matches the input ID.
func (s runbookService) Add(runbook *model.Runbook) (*model.Runbook, error) {
	if runbook == nil {
		return nil, createInvalidParameterError("Add", "runbook")
	}

	path, err := getAddPath(s, runbook)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), runbook, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Runbook), nil
}
