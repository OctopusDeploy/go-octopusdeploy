package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Lifecycles struct {
	Items []Lifecycle `json:"Items"`
	PagedResults
}

type Lifecycle struct {
	Name                    string          `json:"Name" validate:"required"`
	Description             string          `json:"Description,omitempty"`
	ReleaseRetentionPolicy  RetentionPeriod `json:"ReleaseRetentionPolicy,omitempty"`
	TentacleRetentionPolicy RetentionPeriod `json:"TentacleRetentionPolicy,omitempty"`
	Phases                  []Phase         `json:"Phases"`

	Resource
}

const (
	RetentionUnitDays  = RetentionUnit("Days")
	RetentionUnitItems = RetentionUnit("Items")
)

func NewLifecycle(name string) (*Lifecycle, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewLifecycle", "name")
	}

	return &Lifecycle{
		Name:   name,
		Phases: []Phase{},
		TentacleRetentionPolicy: RetentionPeriod{
			Unit:           RetentionUnitDays,
			QuantityToKeep: 30,
		},
		ReleaseRetentionPolicy: RetentionPeriod{
			Unit:           RetentionUnitDays,
			QuantityToKeep: 30,
		},
	}, nil
}

// GetID returns the ID value of the Lifecycle.
func (resource Lifecycle) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Lifecycle.
func (resource Lifecycle) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Lifecycle was changed.
func (resource Lifecycle) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Lifecycle.
func (resource Lifecycle) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Lifecycle and returns an error if invalid.
func (resource Lifecycle) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	if resource.Phases != nil {
		for _, phase := range resource.Phases {
			phaseErr := validate.Struct(phase)

			if phaseErr != nil {
				return phaseErr
			}
		}
	}

	return nil
}

var _ ResourceInterface = &Lifecycle{}
