package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
)

type RunbookSnapshotTemplate struct {
	NextNameIncrement string                             `json:"NextNameIncrement,omitempty"`
	Packages          []*releases.ReleaseTemplatePackage `json:"Packages"`
	RunbookID         string                             `json:"RunbookId,omitempty"`
	RunbookProcessID  string                             `json:"RunbookProcessId,omitempty"`

	resources.Resource
}
