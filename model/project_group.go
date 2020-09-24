package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ProjectGroups struct {
	Items []ProjectGroup `json:"Items"`
	PagedResults
}

type ProjectGroup struct {
	Description       string   `json:"Description,omitempty"`
	EnvironmentIds    []string `json:"EnvironmentIds"`
	Name              string   `json:"Name,omitempty" validate:"required"`
	RetentionPolicyID string   `json:"RetentionPolicyId,omitempty"`

	Resource
}

// GetID returns the ID value of the ProjectGroup.
func (resource ProjectGroup) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this ProjectGroup.
func (resource ProjectGroup) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this ProjectGroup was changed.
func (resource ProjectGroup) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this ProjectGroup.
func (resource ProjectGroup) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the ProjectGroup and returns an error if invalid.
func (resource ProjectGroup) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

func NewProjectGroup(name string) *ProjectGroup {
	return &ProjectGroup{
		Name: name,
	}
}

var _ ResourceInterface = &ProjectGroup{}
