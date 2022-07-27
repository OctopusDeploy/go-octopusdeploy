package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type RunbookService struct {
	services.CanDeleteService
}

func NewRunbookService(sling *sling.Sling, uriTemplate string) *RunbookService {
	return &RunbookService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceRunbookService, sling, uriTemplate),
		},
	}
}

// Add returns the runbook that matches the input ID.
func (s *RunbookService) Add(runbook *Runbook) (*Runbook, error) {
	if IsNil(runbook) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRunbook)
	}

	if err := runbook.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	path, err := services.GetAddPath(s, runbook)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), runbook, new(Runbook), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Runbook), nil
}

// GetAll returns all runbooks. If none can be found or an error occurs, it
// returns an empty collection.
func (s *RunbookService) GetAll() ([]*Runbook, error) {
	items := []*Runbook{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the runbook that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *RunbookService) GetByID(id string) (*Runbook, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Runbook), path)
	if err != nil {
		return nil, internal.CreateResourceNotFoundError("runbook", "ID", id)
	}

	return resp.(*Runbook), nil
}

func (s *RunbookService) GetRunbookSnapshotTemplate(runbook *Runbook) (*RunbookSnapshotTemplate, error) {
	resp, err := api.ApiGet(s.GetClient(), new(RunbookSnapshotTemplate), runbook.Links["RunbookSnapshotTemplate"])
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookSnapshotTemplate), nil
}

// Update modifies a runbook based on the one provided as input.
func (s *RunbookService) Update(runbook *Runbook) (*Runbook, error) {
	if runbook == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterRunbook)
	}

	path, err := services.GetUpdatePath(s, runbook)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), runbook, new(Runbook), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Runbook), nil
}
