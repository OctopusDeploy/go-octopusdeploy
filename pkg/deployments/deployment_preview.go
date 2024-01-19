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

type DeploymentPreviewRequestBody struct {
	EnvironmentId string `json:"EnvironmentId"`
	TenantId      string `json:"TenantId"`
}

// For review time - Should DeploymentPreviewRequestBody be an array of pointers to an object?
type DeploymentPreviewsBody struct {
	DeploymentPreviews   []DeploymentPreviewRequestBody `json:"DeploymentPreviews"`
	IncludeDisabledSteps bool                           `json:"IncludeDisabledSteps"`
	ReleaseId            string                         `json:"ReleaseId"`
	SpaceId              string                         `json:"SpaceId"`
}
