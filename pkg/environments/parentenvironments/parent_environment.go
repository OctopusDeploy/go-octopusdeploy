package parentenvironments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

// AutomaticDeprovisioningRule defines the rule for automatic deprovisioning.
type AutomaticDeprovisioningRule struct {
	ExpiryDays  int `json:"ExpiryDays,omitempty"`
	ExpiryHours int `json:"ExpiryHours,omitempty"`
}

// ParentEnvironment represents a parent environment in Octopus Deploy.
type ParentEnvironment struct {
	Name                        string                       `json:"Name" validate:"required"`
	SpaceID                     string                       `json:"SpaceId" validate:"required"`
	Description                 string                       `json:"Description,omitempty"`
	Slug                        string                       `json:"Slug,omitempty"`
	UseGuidedFailure            bool                         `json:"UseGuidedFailure"`
	AutomaticDeprovisioningRule *AutomaticDeprovisioningRule `json:"AutomaticDeprovisioningRule,omitempty"`

	resources.Resource
}

// NewParentEnvironment creates and initializes a parent environment.
func NewParentEnvironment(name string) *ParentEnvironment {
	return &ParentEnvironment{
		Name:     name,
		Resource: *resources.NewResource(),
	}
}

// Validate checks the state of the parent environment and returns an error if
// invalid.
func (p *ParentEnvironment) Validate() error {
	return validator.New().Struct(p)
}
