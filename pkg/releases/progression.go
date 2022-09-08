package releases

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tasks"
)

// LifecycleProgression represents the progression and deployment status for a single release
type LifecycleProgression struct {
	Phases                         []*LifecycleProgressionPhase `json:"Phases"`
	NextDeployments                []string                     `json:"NextDeployments"`
	NextDeploymentsMinimumRequired int                          `json:"NextDeploymentsMinimumRequired"`

	resources.Resource
}

type LifecycleProgressionPhase struct {
	ID                                 *string            `json:"Id"` // this seems to always be null in the server response
	Name                               string             `json:"Name"`
	Blocked                            bool               `json:"Blocked"`
	Progress                           PhaseProgress      `json:"Progress"`
	Deployments                        []*PhaseDeployment `json:"Deployments"`
	AutomaticDeploymentTargets         []string           `json:"AutomaticDeploymentTargets"`
	OptionalDeploymentTargets          []string           `json:"OptionalDeploymentTargets"`
	MinimumEnvironmentsBeforePromotion int                `json:"MinimumEnvironmentsBeforePromotion"`
	IsOptionalPhase                    bool               `json:"IsOptionalPhase"`

	resources.Resource
}

// PhaseDeployment represents a deployment as part of a progression phase
type PhaseDeployment struct {
	Task *tasks.Task `json:"Task"`
	// If we uncomment this line we get an import cycle releases -> deployments -> releases
	// Deployment *Deployment `json:"Deployment"`
}

type PhaseProgress string

const (
	PhaseProgressPending  = PhaseProgress("Pending")
	PhaseProgressCurrent  = PhaseProgress("Current")
	PhaseProgressComplete = PhaseProgress("Complete")
)
