package model

import (
	"github.com/go-playground/validator/v10"
)

type Lifecycles struct {
	Items []Lifecycle `json:"Items"`
	PagedResults
}

type Lifecycle struct {
	ID                      string          `json:"Id,omitempty"`
	Name                    string          `json:"Name" validate:"required"`
	Description             string          `json:"Description,omitempty"`
	ReleaseRetentionPolicy  RetentionPeriod `json:"ReleaseRetentionPolicy,omitempty"`
	TentacleRetentionPolicy RetentionPeriod `json:"TentacleRetentionPolicy,omitempty"`
	Phases                  []Phase         `json:"Phases"`
}

const (
	RetentionUnitDays  = RetentionUnit("Days")
	RetentionUnitItems = RetentionUnit("Items")
)

func NewLifecycle(name string) *Lifecycle {
	return &Lifecycle{
		Name:   name,
		Phases: []Phase{},
		TentacleRetentionPolicy: RetentionPeriod{
			Unit: RetentionUnitDays,
		},
		ReleaseRetentionPolicy: RetentionPeriod{
			Unit: RetentionUnitDays,
		},
	}
}

// ValidateLifecycleValues checks the values of a Lifecycle object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating lifecycles.
func ValidateLifecycleValues(Lifecycle *Lifecycle) error {
	validate := validator.New()

	err := validate.Struct(Lifecycle)

	if err != nil {
		return err
	}

	if Lifecycle.Phases != nil {
		for _, phase := range Lifecycle.Phases {
			phaseErr := validate.Struct(phase)

			if phaseErr != nil {
				return phaseErr
			}
		}
	}

	return nil
}
