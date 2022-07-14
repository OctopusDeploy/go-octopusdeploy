package variables

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/runbooks"
)

type ProjectVariableSetUsage struct {
	IsCurrentlyBeingUsedInProject bool                                  `json:"IsCurrentlyBeingUsedInProject"`
	ProjectID                     string                                `json:"ProjectId,omitempty"`
	ProjectName                   string                                `json:"ProjectName,omitempty"`
	ProjectSlug                   string                                `json:"ProjectSlug,omitempty"`
	Releases                      []*releases.ReleaseUsageEntry         `json:"Releases"`
	RunbookSnapshots              []*runbooks.RunbookSnapshotUsageEntry `json:"RunbookSnapshots"`
}
