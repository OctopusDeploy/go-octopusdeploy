package releases

// VariableIdentifier identifies a single variable within a variable snapshot.
type VariableIdentifier struct {
	Name    string `json:"Name"`
	OwnerID string `json:"OwnerId"`
}
