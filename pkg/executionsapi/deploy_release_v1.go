package executionsapi

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type CreateExecutionAbstractCommandV1 struct {
	SpaceIDOrName        string            `json:"spaceIdOrName"`
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
	Tenant                   []string `json:"tenants,omitempty"`
	TenantTags               []string `json:"tenantTags,omitempty"`
	ForcePackageRedeployment string   `json:"forcePackageRedeployment,omitempty"`
	UpdateVariableSnapshot   bool     `json:"updateVariableSnapshot,omitempty"`
	CreateExecutionAbstractCommandV1
}

func NewCreateDeploymentTenantedCommandV1(spaceIDOrName string, projectIDOrName string) *CreateDeploymentTenantedCommandV1 {
	return &CreateDeploymentTenantedCommandV1{
		CreateExecutionAbstractCommandV1: CreateExecutionAbstractCommandV1{
			SpaceIDOrName:   spaceIDOrName,
			ProjectIDOrName: projectIDOrName,
		},
	}
}

// MarshalJSON adds the redundant 'spaceId' parameter which is required by the server
func (c *CreateDeploymentTenantedCommandV1) MarshalJSON() ([]byte, error) {
	converted := struct {
		SpaceID string `json:"spaceId"`
		CreateDeploymentTenantedCommandV1
	}{
		SpaceID:                           c.SpaceIDOrName,
		CreateDeploymentTenantedCommandV1: *c,
	}

	return json.Marshal(converted)
}

func CreateDeploymentTenantedV1(client newclient.Client, command *CreateDeploymentTenantedCommandV1) (*CreateDeploymentResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("CreateDeploymentTenantedV1", "command")
	}
	if client.SpaceID() == "" {
		return nil, internal.CreateInvalidClientStateError("CreateDeploymentTenantedV1")
	}
	if command.ReleaseVersion == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.ReleaseVersion")
	}
	if command.EnvironmentName == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.EnvironmentName")
	}

	// Note: command has a SpaceIDOrName field in it, which carries the space, however, we can't use it
	// as the server's route URL *requires* a space **ID**, not a name. In fact, the client's spaceID should always win.
	command.SpaceIDOrName = client.SpaceID()
	url, err := client.URITemplateCache().Expand(uritemplates.CreateDeploymentTenantedCommandV1, map[string]any{"spaceId": client.SpaceID()})
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
	ForcePackageRedeployment string   `json:"forcePackageRedeployment,omitempty"`
	UpdateVariableSnapshot   bool     `json:"updateVariableSnapshot,omitempty"`
	CreateExecutionAbstractCommandV1
}

func NewCreateDeploymentUntenantedCommandV1(spaceIDOrName string, projectIDOrName string) *CreateDeploymentUntenantedCommandV1 {
	return &CreateDeploymentUntenantedCommandV1{
		CreateExecutionAbstractCommandV1: CreateExecutionAbstractCommandV1{
			SpaceIDOrName:   spaceIDOrName,
			ProjectIDOrName: projectIDOrName,
		},
	}
}

// MarshalJSON adds the redundant 'spaceId' parameter which is required by the server
func (c *CreateDeploymentUntenantedCommandV1) MarshalJSON() ([]byte, error) {
	converted := struct {
		SpaceID string `json:"spaceId"`
		CreateDeploymentUntenantedCommandV1
	}{
		SpaceID:                             c.SpaceIDOrName,
		CreateDeploymentUntenantedCommandV1: *c,
	}

	return json.Marshal(converted)
}

func CreateDeploymentUntenantedV1(client newclient.Client, command *CreateDeploymentUntenantedCommandV1) (*CreateDeploymentResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("CreateDeploymentUntenantedV1", "command")
	}
	if client.SpaceID() == "" {
		return nil, internal.CreateInvalidClientStateError("CreateDeploymentUntenantedV1")
	}
	if command.ReleaseVersion == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.ReleaseVersion")
	}
	if len(command.EnvironmentNames) == 0 {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("command.EnvironmentNames")
	}

	// Note: command has a SpaceIDOrName field in it, which carries the space, however, we can't use it
	// as the server's route URL *requires* a space **ID**, not a name. In fact, the client's spaceID should always win.
	command.SpaceIDOrName = client.SpaceID()
	url, err := client.URITemplateCache().Expand(uritemplates.CreateDeploymentUntenantedCommandV1, map[string]any{"spaceId": client.SpaceID()})
	if err != nil {
		return nil, err
	}
	resp, err := services.ApiPost(client.Sling(), command, new(CreateDeploymentResponseV1), url)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateDeploymentResponseV1), nil
}
