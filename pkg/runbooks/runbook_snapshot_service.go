package runbooks

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
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
	return s.addOrPublish(runbookSnapshot, false)
}

// Publish a new runbook snapshot
func (s *RunbookSnapshotService) Publish(runbookSnapshot *RunbookSnapshot) (*RunbookSnapshot, error) {
	return s.addOrPublish(runbookSnapshot, true)
}

func (s *RunbookSnapshotService) addOrPublish(runbookSnapshot *RunbookSnapshot, publish bool) (*RunbookSnapshot, error) {
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

	if publish {
		path = fmt.Sprintf("%s?publish=true", path)
	}

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

// ----- Experimental ---------------------------------------------------------
// SnapshotVariablesByName requires octopus feature toggle partial-updates-on-variables = true
func SnapshotVariablesByName(client newclient.Client, runbookSnapshot *RunbookSnapshot, variables []core.VariableIdentifier, concurrencyToken *string) (*RunbookSnapshot, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("SnapshotVariablesByName", "client")
	}
	if runbookSnapshot == nil {
		return nil, internal.CreateInvalidParameterError("SnapshotVariablesByName", "runbookSnapshot")
	}
	if len(variables) == 0 {
		return nil, internal.CreateInvalidParameterError("SnapshotVariablesByName", "variables")
	}

	for i, v := range variables {
		if v.Name == "" {
			return nil, internal.CreateInvalidParameterError("SnapshotVariablesByName", fmt.Sprintf("variables[%d].Name", i))
		}
		if v.OwnerID == "" {
			return nil, internal.CreateInvalidParameterError("SnapshotVariablesByName", fmt.Sprintf("variables[%d].OwnerId", i))
		}
	}

	spaceId, err := internal.GetSpaceID(runbookSnapshot.SpaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbookSnapshotVariablesByName, map[string]any{
		"spaceId":           spaceId,
		"runbookSnapshotId": runbookSnapshot.ID,
	})
	if err != nil {
		return nil, err
	}

	command := core.SnapshotVariablesByNameCommand{Variables: variables, VariableSnapshotConcurrencyToken: concurrencyToken}
	return newclient.Post[RunbookSnapshot](client.HttpSession(), expandedUri, command)
}
