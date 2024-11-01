package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type RunbookSnapshotTemplate struct {
	NextNameIncrement string                                `json:"NextNameIncrement,omitempty"`
	Packages          []*releases.ReleaseTemplatePackage    `json:"Packages"`
	GitResources      []releases.ReleaseTemplateGitResource `json:"GitResources,omitempty"`
	RunbookID         string                                `json:"RunbookId,omitempty"`
	RunbookProcessID  string                                `json:"RunbookProcessId,omitempty"`

	resources.Resource
}
