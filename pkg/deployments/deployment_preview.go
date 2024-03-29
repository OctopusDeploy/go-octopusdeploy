package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type MachineDeploymentPreview struct {
	ID                string `json:"Id,omitempty"`
	Name              string `json:"Name,omitempty"`
	HasLatestCalamari bool   `json:"HasLatestCalamari"`
	HealthStatus      string `json:"HealthStatus,omitempty"` // machines.HealthStatus validate:"omitempty,oneof=HasWarnings Healthy Unavailable Unhealthy Unknown"`
}

type DeploymentTemplateStep struct {
	ActionID                string                         `json:"ActionId,omitempty"`
	ActionName              string                         `json:"ActionName,omitempty"`
	ActionNumber            string                         `json:"ActionNumber,omitempty"`
	Roles                   []string                       `json:"Roles,omitempty"`
	MachineNames            []string                       `json:"MachineNames,omitempty"`
	Machines                []*MachineDeploymentPreview    `json:"Machines,omitempty"`
	CanBeSkipped            bool                           `json:"CanBeSkipped"`
	IsDisabled              bool                           `json:"IsDisabled"`
	HasNoApplicableMachines bool                           `json:"HasNoApplicableMachines"`
	UnavailableMachines     []*resources.ReferenceDataItem `json:"UnavailableMachines,omitempty"`
	ExcludedMachines        []*resources.ReferenceDataItem `json:"ExcludedMachines,omitempty"`
}

type Control struct {
	Type            string                     `json:"Type"`
	Name            string                     `json:"Name"`
	Label           string                     `json:"Label"`
	Description     string                     `json:"Description"`
	Required        bool                       `json:"Required"`
	DisplaySettings *resources.DisplaySettings `json:"DisplaySettings"`
}

type Element struct {
	Name            string   `json:"Name"`
	Control         *Control `json:"Control"`
	IsValueRequired bool     `json:"IsValueRequired"`
}

type Form struct {
	Values   map[string]string `json:"Values"`
	Elements []*Element        `json:"Elements"`
}

type DeploymentPreview struct {
	// Changes []*ReleaseChanges // we don't use this at the moment, and it is large+expensive, so don't de-serialize for now
	// ChangesMarkdown string  // we don't use this at the moment, and it is large+expensive, so don't de-serialize for now
	Form                          *Form                     `json:"Form,omitempty"`
	StepsToExecute                []*DeploymentTemplateStep `json:"StepsToExecute,omitempty"`
	UseGuidedFailureModeByDefault bool                      `json:"UseGuidedFailureModeByDefault"`
}

type DeploymentPreviewRequest struct {
	EnvironmentId string `json:"EnvironmentId"`
	TenantId      string `json:"TenantId"`
}

// For review time - Should DeploymentPreviewRequestBody be an array of pointers to an object?
type DeploymentPreviewsBody struct {
	DeploymentPreviews   []DeploymentPreviewRequest `json:"DeploymentPreviews"`
	IncludeDisabledSteps bool                       `json:"IncludeDisabledSteps"`
	ReleaseId            string                     `json:"ReleaseId"`
	SpaceId              string                     `json:"SpaceId"`
}

func NewEmptyDeploymentPreview() *DeploymentPreview {
	return &DeploymentPreview{
		Form:                          &Form{Values: map[string]string{}, Elements: []*Element{}},
		StepsToExecute:                []*DeploymentTemplateStep{},
		UseGuidedFailureModeByDefault: false,
	}
}

func NewFormWithValuesAndElements(values map[string]string, elements []*Element) *Form {
	return &Form{
		Values:   values,
		Elements: elements,
	}
}

func NewElement(name string, control *Control, isValueRequired bool) *Element {
	return &Element{
		Name:            name,
		Control:         control,
		IsValueRequired: isValueRequired,
	}
}

func NewControl(controlType, name, label, description string, required bool, displaySettings *resources.DisplaySettings) *Control {
	return &Control{
		Type:            controlType,
		Name:            name,
		Label:           label,
		Description:     description,
		Required:        required,
		DisplaySettings: displaySettings,
	}
}
