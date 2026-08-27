package core

// VariableIdentifier identifies a single variable within a variable snapshot.
type VariableIdentifier struct {
	Name    string `json:"Name"`
	OwnerID string `json:"OwnerId"`
}

// SnapshotVariablesByNameCommand is the request body for refreshing named variables within a variable snapshot.
type SnapshotVariablesByNameCommand struct {
	Variables []VariableIdentifier `json:"Variables"`

	// VariableSnapshotConcurrencyToken is the token read from the release or runbook snapshot. When supplied,
	// the update fails if the snapshot has changed since. Omit to skip the check.
	VariableSnapshotConcurrencyToken *string `json:"VariableSnapshotConcurrencyToken,omitempty"`
}
