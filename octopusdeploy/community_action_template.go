package octopusdeploy

import "github.com/go-playground/validator/v10"

type CommunityActionTemplates struct {
	Items []*CommunityActionTemplate `json:"Items"`
	PagedResults
}

type CommunityActionTemplate struct {
	Author      string                           `json:"Author,omitempty"`
	Description string                           `json:"Description,omitempty"`
	HistoryURL  string                           `json:"HistoryUrl,omitempty"`
	Name        string                           `json:"Name" validate:"required"`
	Parameters  []ActionTemplateParameter        `json:"Parameters"`
	Properties  map[string]PropertyValueResource `json:"Properties,omitempty"`
	Type        string                           `json:"Type,omitempty"`
	Version     int32                            `json:"Version,omitempty"`
	Website     string                           `json:"Website,omitempty"`

	resource
}

// NewCommunityActionTemplate initializes a community action template.
func NewCommunityActionTemplate(name string) *CommunityActionTemplate {
	return &CommunityActionTemplate{
		Name:       name,
		Parameters: []ActionTemplateParameter{},
		Properties: map[string]PropertyValueResource{},
		resource:   *newResource(),
	}
}

// Validate checks the state of the lifecycle and returns an error if invalid.
func (c *CommunityActionTemplate) Validate() error {
	return validator.New().Struct(c)
}
