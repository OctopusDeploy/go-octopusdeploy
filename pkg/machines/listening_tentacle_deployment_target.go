package machines

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
)

type IDeploymentTarget interface {
	GetEndpoint() IEndpoint
	GetName() string
	GetHealthStatus() string
	GetIsDisabled() bool
}

type ListeningTentacleDeploymentTarget struct {
	Endpoint               *ListeningTentacleEndpoint  `json:"Endpoint" validate:"required"`
	EnvironmentIDs         []string                    `json:"EnvironmentIds,omitempty"`
	HasLatestCalamari      bool                        `json:"HasLatestCalamari,omitempty"`
	HealthStatus           string                      `json:"HealthStatus,omitempty" validate:"omitempty,oneof=HasWarnings Healthy Unavailable Unhealthy Unknown"`
	IsDisabled             bool                        `json:"IsDisabled,omitempty"`
	IsInProcess            bool                        `json:"IsInProcess,omitempty"`
	MachinePolicyID        string                      `json:"MachinePolicyId,omitempty"`
	Name                   string                      `json:"Name,omitempty"`
	OperatingSystem        string                      `json:"OperatingSystem,omitempty"`
	Roles                  []string                    `json:"Roles,omitempty"`
	ShellName              string                      `json:"ShellName,omitempty"`
	ShellVersion           string                      `json:"ShellVersion,omitempty"`
	SpaceID                string                      `json:"SpaceId,omitempty"`
	Status                 string                      `json:"Status,omitempty" validate:"omitempty,oneof=CalamariNeedsUpgrade Disabled NeedsUpgrade Offline Online Unknown"`
	StatusSummary          string                      `json:"StatusSummary,omitempty"`
	TenantedDeploymentMode core.TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantIDs              []string                    `json:"TenantIds,omitempty"`
	TenantTags             []string                    `json:"TenantTags,omitempty"`
	Thumbprint             string                      `json:"Thumbprint,omitempty"`
	URI                    string                      `json:"Uri,omitempty" validate:"omitempty,uri"`

	resources.Resource
}

func NewListeningTentacleDeploymentTarget(name string, endpoint *ListeningTentacleEndpoint, environmentIDs []string, roles []string) *ListeningTentacleDeploymentTarget {
	return &ListeningTentacleDeploymentTarget{
		Endpoint:               endpoint,
		EnvironmentIDs:         environmentIDs,
		Name:                   name,
		OperatingSystem:        "Unknown",
		Roles:                  roles,
		ShellName:              "Unknown",
		ShellVersion:           "Unknown",
		TenantIDs:              []string{},
		TenantTags:             []string{},
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),

		Resource: *resources.NewResource(),
	}
}

func (l *ListeningTentacleDeploymentTarget) GetEndpoint() IEndpoint {
	return l.Endpoint
}

func (l *ListeningTentacleDeploymentTarget) GetHealthStatus() string {
	return l.HealthStatus
}

func (l *ListeningTentacleDeploymentTarget) GetIsDisabled() bool {
	return l.IsDisabled
}

func (l *ListeningTentacleDeploymentTarget) GetName() string {
	return l.Name
}

var _ IDeploymentTarget = &ListeningTentacleDeploymentTarget{}
