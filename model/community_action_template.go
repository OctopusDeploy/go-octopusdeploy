package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CommunityActionTemplates struct {
	Items []CommunityActionTemplate `json:"Items"`
	PagedResults
}

type CommunityActionTemplate struct {
	Author      string                           `json:"Author,omitempty"`
	Description string                           `json:"Description,omitempty"`
	HistoryURL  string                           `json:"HistoryUrl,omitempty"`
	Name        string                           `json:"Name,omitempty"`
	Parameters  []*ActionTemplateParameter       `json:"Parameters"`
	Properties  map[string]PropertyValueResource `json:"Properties,omitempty"`
	Type        string                           `json:"Type,omitempty"`
	Version     int32                            `json:"Version,omitempty"`
	Website     string                           `json:"Website,omitempty"`

	Resource
}

// NewCommunityActionTemplate initializes a community action template. If any
// input parameters are invalid, it will return nil and an error.
func NewCommunityActionTemplate(name string) (*CommunityActionTemplate, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("CommunityActionTemplate", "name")
	}

	return &CommunityActionTemplate{
		Name: name,
	}, nil
}

// GetID returns the ID value of the CommunityActionTemplate.
func (resource CommunityActionTemplate) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this CommunityActionTemplate.
func (resource CommunityActionTemplate) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this CommunityActionTemplate was changed.
func (resource CommunityActionTemplate) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this CommunityActionTemplate.
func (resource CommunityActionTemplate) GetLinks() map[string]string {
	return resource.Links
}

// SetID
func (resource CommunityActionTemplate) SetID(id string) {
	resource.ID = id
}

// SetLastModifiedBy
func (resource CommunityActionTemplate) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

// SetLastModifiedOn
func (resource CommunityActionTemplate) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the CommunityActionTemplate and returns an error if invalid.
func (resource CommunityActionTemplate) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &CommunityActionTemplate{}
