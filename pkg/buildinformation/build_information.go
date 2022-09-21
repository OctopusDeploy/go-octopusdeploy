package buildinformation

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/issuetrackers"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type BuildInformationCommand struct {
	PackageID               string                  `json:"PackageId"`
	SpaceID                 string                  `json:"-" uri:"spaceId"`
	OctopusBuildInformation OctopusBuildInformation `json:"OctopusBuildInformation"`
	OverwriteMode           OverwriteMode           `json:"-" uri:"overwriteMode,omitempty"`
	Version                 string                  `json:"Version"`
}

func NewBuildInformationCommand(spaceID string, packageID string, version string) *BuildInformationCommand {
	return &BuildInformationCommand{
		PackageID:     packageID,
		SpaceID:       spaceID,
		OverwriteMode: OverwriteModeFailIfExists,
		Version:       version,
	}
}

type BuildInformationResponse struct {
	Branch                string                         `json:"Branch,omitempty"`
	BuildEnvironment      string                         `json:"BuildEnvironment,omitempty"`
	BuildNumber           string                         `json:"BuildNumber,omitempty"`
	BuildURL              string                         `json:"BuildUrl,omitempty"`
	Commits               []*issuetrackers.CommitDetails `json:"Commits"`
	Created               *time.Time                     `json:"Created"`
	IncompleteDataWarning string                         `json:"IncompleteDataWarning"`
	IssueTrackerName      string                         `json:"IssueTrackerName"`
	PackageID             string                         `json:"PackageId"`
	VcsCommitNumber       string                         `json:"VcsCommitNumber,omitempty"`
	VcsCommitURL          string                         `json:"VcsCommitUrl,omitempty"`
	VcsRoot               string                         `json:"VcsRoot,omitempty"`
	VcsType               string                         `json:"VcsType,omitempty"`
	Version               string                         `json:"Version"`
	WorkItems             []core.WorkItemLink            `json:"WorkItems"`

	resources.Resource
}
