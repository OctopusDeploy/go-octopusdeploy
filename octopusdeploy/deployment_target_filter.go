package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
)

type DeploymentTargetFilter struct {
	Environments    []string `json:"EnvironmentIds,omitempty"`
	EventCategories []string `json:"EventCategories,omitempty"`
	EventGroups     []string `json:"EventGroups,omitempty"`
	Roles           []string `json:"Roles,omitempty"`

	triggerFilter
}

func NewDeploymentTargetFilter(environments []string, eventCategories []string, eventGroups []string, roles []string) *DeploymentTargetFilter {
	deploymentTargetFilter := &DeploymentTargetFilter{
		Environments:    environments,
		EventCategories: eventCategories,
		EventGroups:     eventGroups,
		Roles:           roles,
		triggerFilter:   *newTriggerFilter(MachineFilter),
	}

	return deploymentTargetFilter
}

func (f *DeploymentTargetFilter) GetFilterType() FilterType {
	return f.Type
}

func (f *DeploymentTargetFilter) SetFilterType(filterType FilterType) {
	f.Type = filterType
}

func (f *DeploymentTargetFilter) Validate() error {
	return validator.New().Struct(f)
}

var _ ITriggerFilter = &DeploymentTargetFilter{}
