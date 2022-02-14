package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// CommunityActionTemplate represents a community action template in Octopus
// Deploy.
type CommunityActionTemplate struct {
	ActionType  string                    `json:"ActionType" validate:"required,notblank"`
	Author      string                    `json:"Author,omitempty"`
	Category    string                    `json:"Category,omitempty"`
	Description string                    `json:"Description,omitempty"`
	ExternalId  *uuid.UUID                `json:"ExternalId,omitempty"`
	HistoryURL  string                    `json:"HistoryUrl,omitempty"`
	Name        string                    `json:"Name" validate:"required"`
	Packages    []PackageReference        `json:"Packages,omitempty"`
	Parameters  []ActionTemplateParameter `json:"Parameters"`
	Properties  map[string]PropertyValue  `json:"Properties,omitempty"`
	Type        string                    `json:"Type,omitempty"`
	Version     int32                     `json:"Version,omitempty"`
	Website     string                    `json:"Website,omitempty"`

	Resource
}

// NewCommunityActionTemplate initializes a community action template.
func NewCommunityActionTemplate(name string, actionType string) *CommunityActionTemplate {
	return &CommunityActionTemplate{
		ActionType: actionType,
		Name:       name,
		Packages:   []PackageReference{},
		Parameters: []ActionTemplateParameter{},
		Properties: map[string]PropertyValue{},
		Resource:   *NewResource(),
	}
}

// Validate checks the state of this CommunityActionTemplate and returns an
// error if invalid.
func (c *CommunityActionTemplate) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(c)
}