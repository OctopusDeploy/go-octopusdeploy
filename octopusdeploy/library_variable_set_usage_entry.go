package octopusdeploy

type LibraryVariableSetUsageEntry struct {
	LibraryVariableSetID   string `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string `json:"LibraryVariableSetName,omitempty"`

	resource
}

func NewLibraryVariableSetUsageEntry() *LibraryVariableSetUsageEntry {
	return &LibraryVariableSetUsageEntry{
		resource: *newResource(),
	}
}
