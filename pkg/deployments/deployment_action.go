package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type DeploymentAction struct {
	ActionType                    string                        `json:"ActionType" validate:"required,notblank"`
	CanBeUsedForProjectVersioning bool                          `json:"CanBeUsedForProjectVersioning,omitempty"`
	Channels                      []string                      `json:"Channels,omitempty"`
	Condition                     string                        `json:"Condition,omitempty"`
	Container                     *DeploymentActionContainer    `json:"Container,omitempty"`
	Environments                  []string                      `json:"Environments,omitempty"`
	ExcludedEnvironments          []string                      `json:"ExcludedEnvironments,omitempty"`
	IsDisabled                    bool                          `json:"IsDisabled,omitempty"`
	IsRequired                    bool                          `json:"IsRequired,omitempty"`
	Name                          string                        `json:"Name" validate:"required,notblank"`
	Notes                         string                        `json:"Notes,omitempty"`
	Packages                      []*packages.PackageReference  `json:"Packages,omitempty"`
	Properties                    map[string]core.PropertyValue `json:"Properties,omitempty"`
	StepPackageVersion            string                        `json:"StepPackageVersion,omitempty"`
	TenantTags                    []string                      `json:"TenantTags,omitempty"`
	WorkerPool                    string                        `json:"WorkerPoolId,omitempty"`
	WorkerPoolVariable            string                        `json:"WorkerPoolVariable,omitempty"`

	resources.Resource
}

// NewDeploymentAction initializes a DeploymentAction with a name.
func NewDeploymentAction(name string, actionType string) *DeploymentAction {
	return &DeploymentAction{
		ActionType: actionType,
		Name:       name,
		Properties: map[string]core.PropertyValue{},
		Resource:   *resources.NewResource(),
	}
}

// Validate checks the state of the channel and returns an error if invalid.
func (d DeploymentAction) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(d)
}
