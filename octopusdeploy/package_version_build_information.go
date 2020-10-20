package octopusdeploy

import "time"

type PackageVersionBuildInformation struct {
	Branch                string           `json:"Branch,omitempty"`
	BuildEnvironment      string           `json:"BuildEnvironment,omitempty"`
	BuildNumber           string           `json:"BuildNumber,omitempty"`
	BuildURL              string           `json:"BuildUrl,omitempty"`
	Commits               []*CommitDetails `json:"Commits"`
	Created               time.Time        `json:"Created,omitempty"`
	IncompleteDataWarning string           `json:"IncompleteDataWarning,omitempty"`
	IssueTrackerName      string           `json:"IssueTrackerName,omitempty"`
	PackageID             string           `json:"PackageId,omitempty"`
	VcsCommitNumber       string           `json:"VcsCommitNumber,omitempty"`
	VcsCommitURL          string           `json:"VcsCommitUrl,omitempty"`
	VcsRoot               string           `json:"VcsRoot,omitempty"`
	VcsType               string           `json:"VcsType,omitempty"`
	Version               string           `json:"Version,omitempty"`
	WorkItems             []*WorkItemLink  `json:"WorkItems"`

	Resource
}

func NewPackageVersionBuildInformation() *PackageVersionBuildInformation {
	return &PackageVersionBuildInformation{
		Resource: *newResource(),
	}
}
