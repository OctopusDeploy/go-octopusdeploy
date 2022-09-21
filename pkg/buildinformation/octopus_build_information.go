package buildinformation

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/issuetrackers"
)

type OctopusBuildInformation struct {
	Branch           string                  `json:"Branch,omitempty"`
	BuildEnvironment string                  `json:"BuildEnvironment,omitempty"`
	BuildNumber      string                  `json:"BuildNumber,omitempty"`
	BuildURL         string                  `json:"BuildUrl,omitempty"`
	Commits          []*issuetrackers.Commit `json:"Commits"`
	VcsCommitNumber  string                  `json:"VcsCommitNumber,omitempty"`
	VcsCommitURL     string                  `json:"VcsCommitUrl,omitempty"`
	VcsRoot          string                  `json:"VcsRoot,omitempty"`
	VcsType          string                  `json:"VcsType,omitempty"`
}

func NewOctopusBuildInformation() *OctopusBuildInformation {
	return &OctopusBuildInformation{
		Commits: []*issuetrackers.Commit{},
	}
}
