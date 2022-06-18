package releases

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/core"

type ReleasePackageVersionBuildInformation struct {
	Branch           string                `json:"Branch,omitempty"`
	BuildEnvironment string                `json:"BuildEnvironment,omitempty"`
	BuildNumber      string                `json:"BuildNumber,omitempty"`
	BuildURL         string                `json:"BuildUrl,omitempty"`
	Commits          []*core.CommitDetails `json:"Commits"`
	PackageID        string                `json:"PackageId,omitempty"`
	VcsCommitNumber  string                `json:"VcsCommitNumber,omitempty"`
	VcsCommitURL     string                `json:"VcsCommitUrl,omitempty"`
	VcsRoot          string                `json:"VcsRoot,omitempty"`
	VcsType          string                `json:"VcsType,omitempty"`
	Version          string                `json:"Version,omitempty"`
	WorkItems        []*core.WorkItemLink  `json:"WorkItems"`
}
