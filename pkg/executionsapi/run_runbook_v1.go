package executionsapi

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/uritemplates"
)

type RunbookRunCommandV1 struct {
	RunbookName      string   `json:"runbookName"` // required
	EnvironmentNames []string `json:"environmentNames,omitempty"`
	TenantNames      []string `json:"tenants,omitempty"`
	TenantTags       []string `json:"tenantTags,omitempty"`
	Snapshot         []string `json:"snapshot,omitempty"`
	CreateExecutionAbstractCommandV1
}

type RunbookRunResponseV1 struct {
	DeploymentServerTasks []*DeploymentServerTask `json:"DeploymentServerTasks,omitempty"`
}

func NewRunbookRunCommandV1(spaceIDOrName string, projectIDOrName string) *RunbookRunCommandV1 {
	return &RunbookRunCommandV1{
		CreateExecutionAbstractCommandV1: CreateExecutionAbstractCommandV1{
			SpaceIDOrName:   spaceIDOrName,
			ProjectIDOrName: projectIDOrName,
		},
	}
}

// MarshalJSON adds the redundant 'spaceId' parameter which is required by the server
func (r *RunbookRunCommandV1) MarshalJSON() ([]byte, error) {
	command := struct {
		SpaceID string `json:"spaceId"`
		RunbookRunCommandV1
	}{
		SpaceID:             r.SpaceIDOrName,
		RunbookRunCommandV1: *r,
	}
	return json.Marshal(command)
}

func RunbookRunV1(client *client.Client, command *RunbookRunCommandV1) (*RunbookRunResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("RunbookRunV1", "command")
	}
	if client.SpaceID == "" {
		return nil, internal.CreateInvalidClientStateError("RunbookRunV1")
	}

	// Note: command has a SpaceIDOrName field in it, which carries the space, however, we can't use it
	// as the server's route URL *requires* a space **ID**, not a name. In fact, the client's spaceID should always win.
	command.SpaceIDOrName = client.SpaceID
	url, err := client.URITemplateCache.Expand(uritemplates.CreateRunRunbookCommand, map[string]any{"spaceId": client.SpaceID})
	if err != nil {
		return nil, err
	}
	resp, err := services.ApiPost(client.Root.GetClient(), command, new(RunbookRunResponseV1), url)
	if err != nil {
		return nil, err
	}
	return resp.(*RunbookRunResponseV1), nil
}
