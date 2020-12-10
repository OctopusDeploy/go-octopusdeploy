package octopusdeploy

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type Lifecycles struct {
	Items []*Lifecycle `json:"Items"`
	PagedResults
}

type Lifecycle struct {
	Description             string          `json:"Description,omitempty"`
	Name                    string          `json:"Name" validate:"required"`
	Phases                  []Phase         `json:"Phases,omitempty"`
	ReleaseRetentionPolicy  RetentionPeriod `json:"ReleaseRetentionPolicy,omitempty"`
	SpaceID                 string          `json:"SpaceId,omitempty"`
	TentacleRetentionPolicy RetentionPeriod `json:"TentacleRetentionPolicy,omitempty"`

	resource
}

const (
	RetentionUnitDays  string = "Days"
	RetentionUnitItems string = "Items"
)

func NewLifecycle(name string) *Lifecycle {
	return &Lifecycle{
		Name:   strings.TrimSpace(name),
		Phases: []Phase{},
		ReleaseRetentionPolicy: RetentionPeriod{
			Unit:           "Days",
			QuantityToKeep: 30,
		},
		TentacleRetentionPolicy: RetentionPeriod{
			Unit:           "Days",
			QuantityToKeep: 30,
		},
		resource: *newResource(),
	}
}

// Validate checks the state of the lifecycle and returns an error if invalid.
func (l *Lifecycle) Validate() error {
	validate := validator.New()
	err := validate.Struct(l)

	if err != nil {
		return err
	}

	if l.Phases != nil {
		for _, phase := range l.Phases {
			phaseErr := validate.Struct(phase)

			if phaseErr != nil {
				return phaseErr
			}
		}
	}

	return nil
}
