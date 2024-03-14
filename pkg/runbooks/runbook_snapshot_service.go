package runbooks

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type RunbookSnapshotService struct {
	services.CanDeleteService
}

func NewRunbookSnapshotService(sling *sling.Sling, uriTemplate string) *RunbookSnapshotService {
	return &RunbookSnapshotService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceRunbookSnapshotService, sling, uriTemplate),
		},
	}
}

// Add creates a new runbook snapshot.
func (s *RunbookSnapshotService) Add(runbookSnapshot *RunbookSnapshot) (*RunbookSnapshot, error) {
	if IsNil(runbookSnapshot) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRunbookSnapshot)
	}

	if err := runbookSnapshot.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	path, err := services.GetAddPath(s, runbookSnapshot)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiAdd(s.GetClient(), runbookSnapshot, new(RunbookSnapshot), path)
	if err != nil {
		return nil, err
	}

	return response.(*RunbookSnapshot), nil
}

// Publishes a runbook snapshot
func (s *RunbookSnapshotService) Publish(runbookSnapshot *RunbookSnapshot) (*RunbookSnapshot, error) {
	if IsNil(runbookSnapshot) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRunbookSnapshot)
	}

	if err := runbookSnapshot.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	path, err := services.GetAddPath(s, runbookSnapshot)
	if err != nil {
		return nil, err
	}

	path = fmt.Sprintf("%s?publish=true", path)

	response, err := services.ApiAdd(s.GetClient(), runbookSnapshot, new(RunbookSnapshot), path)
	if err != nil {
		return nil, err
	}

	return response.(*RunbookSnapshot), nil
}

// GetByID returns the release that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *RunbookSnapshotService) GetByID(id string) (*RunbookSnapshot, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(RunbookSnapshot), path)
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookSnapshot), nil
}
