package runbooks

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type RunbookRunCommandV1 struct {
	RunbookName      string   `json:"runbookName"` // required
	EnvironmentNames []string `json:"environmentNames,omitempty"`
	TenantNames      []string `json:"tenants,omitempty"`
	TenantTags       []string `json:"tenantTags,omitempty"`
	Snapshot         []string `json:"snapshot,omitempty"`
	deployments.CreateExecutionAbstractCommandV1
}

type RunbookRunResponseV1 struct {
	DeploymentServerTasks []*deployments.DeploymentServerTask `json:"DeploymentServerTasks,omitempty"`
}

func NewRunbookRunCommandV1(spaceID string, projectIDOrName string) *RunbookRunCommandV1 {
	return &RunbookRunCommandV1{
		CreateExecutionAbstractCommandV1: deployments.CreateExecutionAbstractCommandV1{
			SpaceID:         spaceID,
			ProjectIDOrName: projectIDOrName,
		},
	}
}

// MarshalJSON adds the redundant 'spaceIdOrName' parameter which is required by the server
func (r *RunbookRunCommandV1) MarshalJSON() ([]byte, error) {
	command := struct {
		SpaceIDOrName string `json:"spaceIdOrName"`
		RunbookRunCommandV1
	}{
		SpaceIDOrName:       r.SpaceID,
		RunbookRunCommandV1: *r,
	}
	return json.Marshal(command)
}

func RunbookRunV1(client newclient.Client, command *RunbookRunCommandV1) (*RunbookRunResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("RunbookRunV1", "command")
	}
	if command.SpaceID == "" {
		return nil, internal.CreateInvalidParameterError("RunbookRunV1", "command.SpaceID")
	}

	url, err := client.URITemplateCache().Expand(uritemplates.CreateRunRunbookCommand, map[string]any{"spaceId": command.SpaceID})
	if err != nil {
		return nil, err
	}
	resp, err := services.ApiPost(client.Sling(), command, new(RunbookRunResponseV1), url)
	if err != nil {
		return nil, err
	}
	return resp.(*RunbookRunResponseV1), nil
}
