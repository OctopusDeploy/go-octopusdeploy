package releases

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/dashboard"
)

type ReleaseProgression struct {
	Channel                 *channels.Channel                    `json:"Channel,omitempty"`
	Deployments             map[string][]dashboard.DashboardItem `json:"Deployments,omitempty"`
	HasUnresolvedDefect     bool                                 `json:"HasUnresolvedDefect,omitempty"`
	NextDeployments         []string                             `json:"NextDeployments"`
	Release                 *Release                             `json:"Release,omitempty"`
	ReleaseRetentionPeriod  *core.RetentionPeriod                `json:"ReleaseRetentionPeriod,omitempty"`
	TentacleRetentionPeriod *core.RetentionPeriod                `json:"TentacleRetentionPeriod,omitempty"`
}
