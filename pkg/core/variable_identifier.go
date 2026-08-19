package core

// VariableIdentifier identifies a single variable within a variable snapshot.
//
// OwnerID is required because a variable name is not unique across a snapshot: the snapshot
// merges the project's own variables with a copy of every library variable set the project
// uses, so the same name can appear more than once. OwnerID is either the project
// (e.g. Projects-1) or a library variable set (e.g. LibraryVariableSets-1).
type VariableIdentifier struct {
	Name    string `json:"Name"`
	OwnerID string `json:"OwnerId"`
}
