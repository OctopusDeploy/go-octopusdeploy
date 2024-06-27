package channels

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/gitdependencies"

type ChannelGitResourceRule struct {
	Id                   string                                          `json:"Id,omitempty"`
	GitDependencyActions []gitdependencies.DeploymentActionGitDependency `json:"GitDependencyActions,omitempty"`
	Rules                []string                                        `json:"Rules,omitempty"`
}
