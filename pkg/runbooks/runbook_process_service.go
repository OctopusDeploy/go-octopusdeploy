package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type RunbookProcessService struct {
	services.Service
}

func NewRunbookProcessService(sling *sling.Sling, uriTemplate string) *RunbookProcessService {
	return &RunbookProcessService{
		Service: services.NewService(constants.ServiceRunbookProcessService, sling, uriTemplate),
	}
}

// GetByID returns the runbook process that matches the input ID. If one cannot
// be found, it returns nil and an error.
//
// Deprecated: use runbookprocess.GetByID
func (s *RunbookProcessService) GetByID(id string) (*RunbookProcess, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(RunbookProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookProcess), nil
}

// Update modifies a runbook process based on the one provided as input.
//
// Deprecated: use runbookprocess.Update
func (s *RunbookProcessService) Update(runbook *RunbookProcess) (*RunbookProcess, error) {
	if runbook == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterRunbook)
	}

	path, err := services.GetUpdatePath(s, runbook)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), runbook, new(RunbookProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookProcess), nil
}
