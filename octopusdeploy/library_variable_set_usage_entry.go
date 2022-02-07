package octopusdeploy

type LibraryVariableSetUsageEntry struct {
	LibraryVariableSetID   string `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string `json:"LibraryVariableSetName,omitempty"`

	Resource
}

func NewLibraryVariableSetUsageEntry() *LibraryVariableSetUsageEntry {
	return &LibraryVariableSetUsageEntry{
		Resource: *newResource(),
	}
}
