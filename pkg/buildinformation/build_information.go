package buildinformation

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/issuetrackers"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type BuildInformation struct {
	Branch                string                         `json:"Branch,omitempty"`
	BuildEnvironment      string                         `json:"BuildEnvironment,omitempty"`
	BuildNumber           string                         `json:"BuildNumber,omitempty"`
	BuildURL              string                         `json:"BuildUrl,omitempty"`
	Commits               []*issuetrackers.CommitDetails `json:"Commits"`
	Created               time.Time                      `json:"Created,omitempty"`
	IncompleteDataWarning string                         `json:"IncompleteDataWarning,omitempty"`
	IssueTrackerName      string                         `json:"IssueTrackerName,omitempty"`
	PackageID             string                         `json:"PackageId,omitempty"`
	VcsCommitNumber       string                         `json:"VcsCommitNumber,omitempty"`
	VcsCommitURL          string                         `json:"VcsCommitUrl,omitempty"`
	VcsRoot               string                         `json:"VcsRoot,omitempty"`
	VcsType               string                         `json:"VcsType,omitempty"`
	Version               string                         `json:"Version,omitempty"`
	WorkItems             []*core.WorkItemLink           `json:"WorkItems"`

	resources.Resource
}

func NewBuildInformation() *BuildInformation {
	return &BuildInformation{
		Resource: *resources.NewResource(),
	}
}

type Commit struct {
	Id      string `json:"Id,omitempty"`
	Comment string `json:"Comment,omitempty"`
}

type OctopusBuildInformation struct {
	BuildEnvironment string `json:"BuildEnvironment,omitempty"`
	BuildNumber      string `json:"BuildNumber,omitempty"`
	BuildUrl         string `json:"BuildUrl,omitempty"`
	Branch           string `json:"Branch,omitempty"`
	VcsType          string `json:"VcsType,omitempty"`
	VcsRoot          string `json:"VcsRoot,omitempty"`
	VcsCommitNumber  string `json:"VcsCommitNumber,omitempty"`

	Commits []*Commit `json:"Commits"`
}

type CreateBuildInformationCommand struct {
	SpaceId                 string                   `json:"SpaceId,omitempty"`
	PackageId               string                   `json:"PackageId,omitempty"`
	Version                 string                   `json:"Version,omitempty"`
	OctopusBuildInformation *OctopusBuildInformation `json:"OctopusBuildInformation,omitempty"`
	OverwriteMode           OverwriteMode            `json:"OverwriteMode,omitempty"`
}

func NewCreateBuildInformationCommand(spaceId string, packageId string, version string, buildInformation OctopusBuildInformation) *CreateBuildInformationCommand {
	return &CreateBuildInformationCommand{
		SpaceId:                 spaceId,
		PackageId:               packageId,
		Version:                 version,
		OctopusBuildInformation: &buildInformation,
	}
}
