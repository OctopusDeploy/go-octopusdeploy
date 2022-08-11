package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// DeploymentSettings are loaded by the ProjectService because deployment settings are a subresource of Projects
type DeploymentSettings struct {
	ChangeDescription               string                       `json:"ChangeDescription,omitempty"`
	ConnectivityPolicy              *core.ConnectivityPolicy     `json:"ProjectConnectivityPolicy,omitempty"`
	DefaultGuidedFailureMode        string                       `json:"DefaultGuidedFailureMode,omitempty"`
	DefaultToSkipIfAlreadyInstalled bool                         `json:"DefaultToSkipIfAlreadyInstalled,omitempty"`
	DeploymentChangesTemplate       string                       `json:"DeploymentChangesTemplate,omitempty"`
	ProjectID                       string                       `json:"ProjectId"`
	ReleaseNotesTemplate            string                       `json:"ReleaseNotesTemplate,omitempty"`
	SpaceID                         string                       `json:"SpaceId"`
	VersioningStrategy              *projects.VersioningStrategy `json:"VersioningStrategy,omitempty"`

	resources.Resource
}

func NewDeploymentSettings() *DeploymentSettings {
	return &DeploymentSettings{
		Resource: *resources.NewResource(),
	}
}
