package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type DeploymentAction struct {
	ActionType                    string                     `json:"ActionType" validate:"required,notblank"`
	CanBeUsedForProjectVersioning bool                       `json:"CanBeUsedForProjectVersioning,omitempty"`
	Channels                      []string                   `json:"Channels,omitempty"`
	Condition                     string                     `json:"Condition,omitempty"`
	Container                     *DeploymentActionContainer `json:"Container,omitempty"`
	Environments                  []string                   `json:"Environments,omitempty"`
	ExcludedEnvironments          []string                   `json:"ExcludedEnvironments,omitempty"`
	IsDisabled                    bool                       `json:"IsDisabled,omitempty"`
	IsRequired                    bool                       `json:"IsRequired,omitempty"`
	Name                          string                     `json:"Name" validate:"required,notblank"`
	Notes                         string                     `json:"Notes,omitempty"`
	Packages                      []*PackageReference        `json:"Packages,omitempty"`
	Properties                    map[string]PropertyValue   `json:"Properties,omitempty"`
	TenantTags                    []string                   `json:"TenantTags,omitempty"`
	WorkerPoolID                  string                     `json:"WorkerPoolId,omitempty"`
	WorkerPoolVariable            string                     `json:"WorkerPoolVariable,omitempty"`

	Resource
}

// NewDeploymentAction initializes a DeploymentAction with a name.
func NewDeploymentAction(name string, actionType string) *DeploymentAction {
	return &DeploymentAction{
		ActionType: actionType,
		Name:       name,
		Properties: map[string]PropertyValue{},
		Resource:   *NewResource(),
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
