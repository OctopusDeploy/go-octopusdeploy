package octopusdeploy

import "time"

// RunbookSnapshots defines a collection of runbook snapshots with built-in
// support for paged results from the API.
type RunbookSnapshots struct {
	Items []*RunbookSnapshot `json:"Items"`
	PagedResults
}

// RunbookSnapshot represents a runbook snapshot.
type RunbookSnapshot struct {
	Assembled                     *time.Time         `json:"Assembled,omitempty"`
	FrozenProjectVariableSetID    string             `json:"FrozenProjectVariableSetId,omitempty"`
	FrozenRunbookProcessID        string             `json:"FrozenRunbookProcessId,omitempty"`
	LibraryVariableSetSnapshotIDs []string           `json:"LibraryVariableSetSnapshotIds"`
	Name                          string             `json:"Name,omitempty"`
	Notes                         string             `json:"Notes,omitempty"`
	ProjectID                     string             `json:"ProjectId,omitempty"`
	ProjectVariableSetSnapshotID  string             `json:"ProjectVariableSetSnapshotId,omitempty"`
	RunbookID                     string             `json:"RunbookId,omitempty"`
	SelectedPackages              []*SelectedPackage `json:"SelectedPackages"`
	SpaceID                       string             `json:"SpaceId,omitempty"`

	resource
}

// NewRunbookSnapshot creates and initializes a runbook snapshot.
func NewRunbookSnapshot(name string, projectID string) *RunbookSnapshot {
	return &RunbookSnapshot{
		Name:      name,
		ProjectID: projectID,
		resource:  *newResource(),
	}
}
