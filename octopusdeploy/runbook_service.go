package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type runbookService struct {
	canDeleteService
}

func newRunbookService(sling *sling.Sling, uriTemplate string) *runbookService {
	runbookService := &runbookService{}
	runbookService.service = newService(serviceRunbookService, sling, uriTemplate, new(Runbook))

	return runbookService
}

// Add returns the runbook that matches the input ID.
func (s runbookService) Add(runbook *Runbook) (*Runbook, error) {
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

	return resp.(*Runbook), nil
}

// GetAll returns all runbooks. If none can be found or an error occurs, it
// returns an empty collection.
func (s runbookService) GetAll() ([]*Runbook, error) {
	items := []*Runbook{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the runbook that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s runbookService) GetByID(id string) (*Runbook, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Runbook), path)
	if err != nil {
		return nil, createResourceNotFoundError("runbook", "ID", id)
	}

	return resp.(*Runbook), nil
}

// Update modifies a runbook based on the one provided as input.
func (s runbookService) Update(runbook *Runbook) (*Runbook, error) {
	if runbook == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterRunbook)
	}

	path, err := getUpdatePath(s, runbook)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), runbook, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Runbook), nil
}
