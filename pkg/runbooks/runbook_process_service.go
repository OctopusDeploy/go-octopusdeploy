package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
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
func (s *RunbookProcessService) GetByID(id string) (*RunbookProcess, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(RunbookProcess), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookProcess), nil
}
