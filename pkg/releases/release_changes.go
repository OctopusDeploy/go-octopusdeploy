package releases

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"

type ReleaseChanges struct {
	BuildInformation []*ReleasePackageVersionBuildInformation `json:"BuildInformation"`
	Commits          []*core.CommitDetails                    `json:"Commits"`
	ReleaseNotes     string                                   `json:"ReleaseNotes,omitempty"`
	Version          string                                   `json:"Version,omitempty"`
	WorkItems        []*core.WorkItemLink                     `json:"WorkItems"`
}
