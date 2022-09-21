package releases

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/issuetrackers"
)

type ReleaseChanges struct {
	BuildInformation []*ReleasePackageVersionBuildInformation `json:"BuildInformation"`
	Commits          []*issuetrackers.CommitDetails           `json:"Commits"`
	ReleaseNotes     string                                   `json:"ReleaseNotes,omitempty"`
	Version          string                                   `json:"Version,omitempty"`
	WorkItems        []*core.WorkItemLink                     `json:"WorkItems"`
}
