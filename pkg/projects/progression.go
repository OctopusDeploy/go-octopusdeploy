package projects

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// Progression represents the project level (dashboard-style) progression and current deployed releases for the project
type Progression struct {
	LifecycleEnvironments map[string][]resources.ReferenceDataItem `json:"LifecycleEnvironments,omitempty"`
	Environments          []*resources.ReferenceDataItem           `json:"Environments"`
	ChannelEnvironments   map[string][]resources.ReferenceDataItem `json:"ChannelEnvironments,omitempty"`
	Releases              []*ReleaseProgression                    `json:"Releases"`

	resources.Resource
}

// ReleaseProgression represents information about a release within the context of a Project Progression
// Mirrors ReleaseProgressionResource in the server
type ReleaseProgression struct {
	Release                 *releases.Release           `json:"Release"`
	Channel                 *channels.Channel           `json:"Channel"`
	Deployments             map[string][]*DashboardItem `json:"Deployments"`
	NextDeployments         []string                    `json:"NextDeployments"`
	HasUnresolvedDefect     bool                        `json:"HasUnresolvedDefect"`
	ReleaseRetentionPeriod  *core.RetentionPeriod       `json:"ReleaseRetentionPeriod"`
	TentacleRetentionPeriod *core.RetentionPeriod       `json:"TentacleRetentionPeriod"`
}

type DashboardItem struct {
	ProjectID               string `json:"ProjectId"`
	DeploymentEnvironmentID string `json:"DeploymentEnvironmentId"`
	ReleaseID               string `json:"ReleaseId"`
	DeploymentID            string `json:"DeploymentId"`
	TaskID                  string `json:"TaskId"`
	// TODO more fields if we need them
}
