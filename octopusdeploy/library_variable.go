package octopusdeploy

type LibraryVariable struct {
	LibraryVariableSetID   string                     `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string                     `json:"LibraryVariableSetName,omitempty"`
	Links                  map[string]string          `json:"Links,omitempty"`
	Templates              []*ActionTemplateParameter `json:"Templates"`
	Variables              map[string]PropertyValue   `json:"Variables,omitempty"`
}

func NewLibraryVariable() *LibraryVariable {
	return &LibraryVariable{
		Links:     map[string]string{},
		Variables: map[string]PropertyValue{},
	}
}
