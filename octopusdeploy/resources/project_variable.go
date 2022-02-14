package resources

type ProjectVariable struct {
	Links       map[string]string                   `json:"Links,omitempty"`
	ProjectID   string                              `json:"ProjectId,omitempty"`
	ProjectName string                              `json:"ProjectName,omitempty"`
	Templates   []*ActionTemplateParameter          `json:"Templates"`
	Variables   map[string]map[string]PropertyValue `json:"Variables,omitempty"`
}
