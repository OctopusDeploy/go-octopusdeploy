package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
)

type RunPreview struct {
	StepsToExecute                []*deployments.DeploymentTemplateStep `json:"StepsToExecute,omitempty"`
	UseGuidedFailureModeByDefault bool                                  `json:"UseGuidedFailureModeByDefault"`
}
