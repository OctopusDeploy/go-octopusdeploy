package octopusdeploy

type Library struct {
	LibraryVariableSetID   string                     `json:"LibraryVariableSetId,omitempty"`
	LibraryVariableSetName string                     `json:"LibraryVariableSetName,omitempty"`
	Links                  map[string]string          `json:"Links,omitempty"`
	Templates              []*ActionTemplateParameter `json:"Templates"`
	Variables              map[string]PropertyValue   `json:"Variables,omitempty"`
}
