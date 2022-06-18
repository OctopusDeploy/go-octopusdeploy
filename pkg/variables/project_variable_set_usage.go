package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/runbooks"
)

type ProjectVariableSetUsage struct {
	IsCurrentlyBeingUsedInProject bool                                  `json:"IsCurrentlyBeingUsedInProject"`
	ProjectID                     string                                `json:"ProjectId,omitempty"`
	ProjectName                   string                                `json:"ProjectName,omitempty"`
	ProjectSlug                   string                                `json:"ProjectSlug,omitempty"`
	Releases                      []*releases.ReleaseUsageEntry         `json:"Releases"`
	RunbookSnapshots              []*runbooks.RunbookSnapshotUsageEntry `json:"RunbookSnapshots"`
}
