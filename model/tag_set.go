package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type TagSets struct {
	Items []TagSet `json:"Items"`
	PagedResults
}

type TagSet struct {
	Name string `json:"Name"`
	Tags []Tag  `json:"Tags,omitempty"`

	Resource
}

// NewTagSet initializes a TagSet with a name.
func NewTagSet(name string) *TagSet {
	return &TagSet{
		Name: name,
	}
}

// GetID returns the ID value of the TagSet.
func (resource TagSet) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this TagSet.
func (resource TagSet) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this TagSet was changed.
func (resource TagSet) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this TagSet.
func (resource TagSet) GetLinks() map[string]string {
	return resource.Links
}

func (resource TagSet) SetID(id string) {
	resource.ID = id
}

func (resource TagSet) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource TagSet) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the TagSet and returns an error if invalid.
func (resource TagSet) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &TagSet{}
