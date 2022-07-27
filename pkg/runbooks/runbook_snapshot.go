package runbooks

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// RunbookSnapshot represents a runbook snapshot.
type RunbookSnapshot struct {
	Assembled                     *time.Time                  `json:"Assembled,omitempty"`
	FrozenProjectVariableSetID    string                      `json:"FrozenProjectVariableSetId,omitempty"`
	FrozenRunbookProcessID        string                      `json:"FrozenRunbookProcessId,omitempty"`
	LibraryVariableSetSnapshotIDs []string                    `json:"LibraryVariableSetSnapshotIds"`
	Name                          string                      `json:"Name,omitempty"`
	Notes                         string                      `json:"Notes,omitempty"`
	ProjectID                     string                      `json:"ProjectId" validate:"required,notblank"`
	ProjectVariableSetSnapshotID  string                      `json:"ProjectVariableSetSnapshotId,omitempty"`
	RunbookID                     string                      `json:"RunbookId" validate:"required,notblank"`
	SelectedPackages              []*packages.SelectedPackage `json:"SelectedPackages"`
	SpaceID                       string                      `json:"SpaceId,omitempty"`

	resources.Resource
}

// NewRunbookSnapshot creates and initializes a runbook snapshot.
func NewRunbookSnapshot(name string, projectID string, runbookID string) *RunbookSnapshot {
	return &RunbookSnapshot{
		LibraryVariableSetSnapshotIDs: []string{},
		Name:                          name,
		ProjectID:                     projectID,
		RunbookID:                     runbookID,
		SelectedPackages:              []*packages.SelectedPackage{},
		Resource:                      *resources.NewResource(),
	}
}

// Validate checks the state of the runbook snapshot and returns an error if
// invalid.
func (c RunbookSnapshot) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(c)
}
