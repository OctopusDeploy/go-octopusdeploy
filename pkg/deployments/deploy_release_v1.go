package deployments

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type CreateExecutionAbstractCommandV1 struct {
	// also has awkward SpaceIDOrName; see CreateReleaseV1 for explanation
	SpaceID              string            `json:"spaceId"`
	ProjectIDOrName      string            `json:"projectName"`
	ForcePackageDownload bool              `json:"forcePackageDownload,omitempty"`
	SpecificMachineNames []string          `json:"specificMachineNames,omitempty"`
	ExcludedMachineNames []string          `json:"excludedMachineNames,omitempty"`
	SkipStepNames        []string          `json:"skipStepNames,omitempty"`
	UseGuidedFailure     *bool             `json:"useGuidedFailure,omitempty"` // note: nil is valid, meaning 'use default'
	RunAt                string            `json:"runAt,omitempty"`            // contains a datetimeOffset-parseable value
	NoRunAfter           string            `json:"noRunAfter,omitempty"`       // contains a datetimeOffset-parseable value
	Variables            map[string]string `json:"variables,omitempty"`
}

type DeploymentServerTask struct {
	DeploymentID string `json:"DeploymentId"`
	ServerTaskID string `json:"ServerTaskId"`
}

type CreateDeploymentResponseV1 struct {
	DeploymentServerTasks []*DeploymentServerTask `json:"DeploymentServerTasks,omitempty"`
}

// ----- Tenanted -----------------------------------------------

type CreateDeploymentTenantedCommandV1 struct {
	ReleaseVersion           string   `json:"releaseVersion"`  // required
	EnvironmentName          string   `json:"environmentName"` // required
	Tenants                  []string `json:"tenants,omitempty"`
	TenantTags               []string `json:"tenantTags,omitempty"`
	ForcePackageRedeployment bool     `json:"forcePackageRedeployment,omitempty"`
	UpdateVariableSnapshot   bool     `json:"updateVariableSnapshot,omitempty"`
	CreateExecutionAbstractCommandV1
}

func NewCreateDeploymentTenantedCommandV1(spaceID string, projectIDOrName string) *CreateDeploymentTenantedCommandV1 {
	return &CreateDeploymentTenantedCommandV1{
		CreateExecutionAbstractCommandV1: CreateExecutionAbstractCommandV1{
			SpaceID:         spaceID,
			ProjectIDOrName: projectIDOrName,
		},
	}
}

// MarshalJSON adds the redundant 'spaceIdOrName' parameter which is required by the server
func (c *CreateDeploymentTenantedCommandV1) MarshalJSON() ([]byte, error) {
	converted := struct {
		SpaceIDOrName string `json:"spaceIdOrName"`
		CreateDeploymentTenantedCommandV1
	}{
		SpaceIDOrName:                     c.SpaceID,
		CreateDeploymentTenantedCommandV1: *c,
	}

	return json.Marshal(converted)
}

func CreateDeploymentTenantedV1(client newclient.Client, command *CreateDeploymentTenantedCommandV1) (*CreateDeploymentResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("CreateDeploymentTenantedV1", "command")
	}
	if command.SpaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.SpaceID")
	}
	if command.ReleaseVersion == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.ReleaseVersion")
	}
	if command.EnvironmentName == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.EnvironmentName")
	}

	url, err := client.URITemplateCache().Expand(uritemplates.CreateDeploymentTenantedCommandV1, map[string]any{"spaceId": command.SpaceID})
	if err != nil {
		return nil, err
	}
	resp, err := services.ApiPost(client.Sling(), command, new(CreateDeploymentResponseV1), url)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateDeploymentResponseV1), nil
}

// ----- Untenanted -----------------------------------------------

type CreateDeploymentUntenantedCommandV1 struct {
	ReleaseVersion           string   `json:"releaseVersion"`   // required
	EnvironmentNames         []string `json:"environmentNames"` // required
	ForcePackageRedeployment bool     `json:"forcePackageRedeployment,omitempty"`
	UpdateVariableSnapshot   bool     `json:"updateVariableSnapshot,omitempty"`
	CreateExecutionAbstractCommandV1
}

func NewCreateDeploymentUntenantedCommandV1(spaceID string, projectIDOrName string) *CreateDeploymentUntenantedCommandV1 {
	return &CreateDeploymentUntenantedCommandV1{
		CreateExecutionAbstractCommandV1: CreateExecutionAbstractCommandV1{
			SpaceID:         spaceID,
			ProjectIDOrName: projectIDOrName,
		},
	}
}

// MarshalJSON adds the redundant 'spaceIdOrName' parameter which is required by the server
func (c *CreateDeploymentUntenantedCommandV1) MarshalJSON() ([]byte, error) {
	converted := struct {
		SpaceIDOrName string `json:"spaceIdOrName"`
		CreateDeploymentUntenantedCommandV1
	}{
		SpaceIDOrName:                       c.SpaceID,
		CreateDeploymentUntenantedCommandV1: *c,
	}
	return json.Marshal(converted)
}

func CreateDeploymentUntenantedV1(client newclient.Client, command *CreateDeploymentUntenantedCommandV1) (*CreateDeploymentResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("CreateDeploymentUntenantedV1", "command")
	}
	if command.SpaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.SpaceID")
	}
	if command.ReleaseVersion == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.ReleaseVersion")
	}
	if len(command.EnvironmentNames) == 0 {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.EnvironmentNames")
	}

	url, err := client.URITemplateCache().Expand(uritemplates.CreateDeploymentUntenantedCommandV1, map[string]any{"spaceId": command.SpaceID})
	if err != nil {
		return nil, err
	}
	resp, err := services.ApiPost(client.Sling(), command, new(CreateDeploymentResponseV1), url)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateDeploymentResponseV1), nil
}
