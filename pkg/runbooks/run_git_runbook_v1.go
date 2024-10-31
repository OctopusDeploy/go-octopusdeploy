package runbooks

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type GitRunbookRunCommandV1 struct {
	RunbookName      string   `json:"runbookName"` // required
	EnvironmentNames []string `json:"environmentNames,omitempty"`
	Tenants          []string `json:"tenants,omitempty"`
	TenantTags       []string `json:"tenantTags,omitempty"`
	GitRef           string   `json:"gitRef"` // required
	PackageVersion   string   `json:"packageVersion,omitempty"`
	Packages         []string `json:"packages,omitempty"`
	GitResources     []string `json:"gitResources,omitempty"`
	deployments.CreateExecutionAbstractCommandV1
}

type GitRunbookRunServerTask struct {
	RunbookRunID string `json:"RunbookRunId"`
	ServerTaskID string `json:"ServerTaskId"`
}

type GitRunbookRunResponseV1 struct {
	RunbookRunServerTasks []*RunbookRunServerTask `json:"RunbookRunServerTasks,omitempty"`
}

func NewGitRunbookRunCommandV1(spaceID string, projectIDOrName string) *GitRunbookRunCommandV1 {
	return &GitRunbookRunCommandV1{
		CreateExecutionAbstractCommandV1: deployments.CreateExecutionAbstractCommandV1{
			SpaceID:         spaceID,
			ProjectIDOrName: projectIDOrName,
		},
	}
}

// MarshalJSON adds the redundant 'spaceIdOrName' parameter which is required by the server
func (r *GitRunbookRunCommandV1) MarshalJSON() ([]byte, error) {
	command := struct {
		SpaceIDOrName string `json:"spaceIdOrName"`
		GitRunbookRunCommandV1
	}{
		SpaceIDOrName:          r.SpaceID,
		GitRunbookRunCommandV1: *r,
	}
	return json.Marshal(command)
}

func GitRunbookRunV1(client newclient.Client, command *GitRunbookRunCommandV1) (*GitRunbookRunResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("RunbookRunV1", "command")
	}
	if command.SpaceID == "" {
		return nil, internal.CreateInvalidParameterError("RunbookRunV1", "command.SpaceID")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.CreateRunGitRunbookCommand, map[string]any{"spaceId": command.SpaceID})
	if err != nil {
		return nil, err
	}
	return newclient.Post[GitRunbookRunResponseV1](client.HttpSession(), expandedUri, command)
}
