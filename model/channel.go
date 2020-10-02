package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Channels defines a collection of channels with built-in support for paged results.
type Channels struct {
	Items []Channel `json:"Items"`
	PagedResults
}

type Channel struct {
	Description string        `json:"Description"`
	IsDefault   bool          `json:"IsDefault"`
	LifecycleID string        `json:"LifecycleId"`
	Name        string        `json:"Name"`
	ProjectID   string        `json:"ProjectId"`
	Rules       []ChannelRule `json:"Rules,omitempty"`
	TenantTags  []string      `json:"TenantedDeploymentMode,omitempty"`

	Resource
}

func NewChannel(name, description, projectID string) (*Channel, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewChannel", "name")
	}

	return &Channel{
		Name:        name,
		ProjectID:   projectID,
		Description: description,
	}, nil
}

// GetID returns the ID value of the Channel.
func (resource Channel) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Channel.
func (resource Channel) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Channel was changed.
func (resource Channel) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Channel.
func (resource Channel) GetLinks() map[string]string {
	return resource.Links
}

// SetID
func (resource Channel) SetID(id string) {
	resource.ID = id
}

// SetLastModifiedBy
func (resource Channel) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

// SetLastModifiedOn
func (resource Channel) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the Channel and returns an error if invalid.
func (resource Channel) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return ValidateMultipleProperties([]error{
		ValidateRequiredPropertyValue("Name", resource.Name),
		ValidateRequiredPropertyValue("ProjectID", resource.ProjectID),
	})
}

var _ ResourceInterface = &Channel{}
