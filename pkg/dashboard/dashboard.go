package dashboard

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/interruptions"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type Dashboard struct {
	Projects      []DashboardProject      `json:"Projects,omitempty"`
	ProjectGroups []DashboardProjectGroup `json:"ProjectGroups,omitempty"`
	Environments  []DashboardEnvironment  `json:"Environments,omitempty"`
	Tenants       []DashboardTenant       `json:"Tenants,omitempty"`
	Items         []DashboardItem         `json:"Items,omitempty"`
	PreviousItems []DashboardItem         `json:"PreviousItems,omitempty"`
	ProjectLimit  *int                    `json:"ProjectLimit,omitempty"`
	IsFiltered    bool                    `json:"IsFiltered"`

	resources.Resource
}

type DashboardProject struct {
	Name                           string   `json:"Name,omitempty"`
	IsDisabled                     bool     `json:"IsDisabled"`
	Slug                           string   `json:"Slug,omitempty"`
	ProjectGroupID                 string   `json:"ProjectGroupId,omitempty"`
	EnvironmentIDs                 []string `json:"EnvironmentIds,omitempty"`
	TenantedDeploymentMode         string   `json:"TenantedDeploymentMode,omitempty"`
	CanPerformUntenantedDeployment bool     `json:"CanPerformUntenantedDeployment"`

	resources.Resource
}

type DashboardProjectGroup struct {
	Name           string   `json:"Name,omitempty"`
	EnvironmentIDs []string `json:"EnvironmentIds,omitempty"`

	resources.Resource
}

type DashboardEnvironment struct {
	Name string `json:"Name,omitempty"`

	resources.Resource
}

type DashboardTenant struct {
	Name                string              `json:"Name,omitempty"`
	ProjectEnvironments map[string][]string `json:"ProjectEnvironments,omitempty"`
	TenantTags          []string            `json:"TenantTags,omitempty"`
	IsDisabled          bool                `json:"IsDisabled"`

	resources.Resource
}

type DashboardItem struct {
	ProjectID                string                           `json:"ProjectId,omitempty"`
	EnvironmentID            string                           `json:"EnvironmentId,omitempty"`
	ReleaseID                string                           `json:"ReleaseId,omitempty"`
	DeploymentID             string                           `json:"DeploymentId,omitempty"`
	TaskID                   string                           `json:"TaskId,omitempty"`
	TenantID                 string                           `json:"TenantId,omitempty"`
	ChannelID                string                           `json:"ChannelId,omitempty"`
	ReleaseVersion           string                           `json:"ReleaseVersion,omitempty"`
	Created                  *time.Time                       `json:"Created,omitempty"`
	QueueTime                *time.Time                       `json:"QueueTime,omitempty"`
	StartTime                *time.Time                       `json:"StartTime,omitempty"`
	CompletedTime            *time.Time                       `json:"CompletedTime,omitempty"`
	State                    string                           `json:"State,omitempty"` // Enum: [Canceled Cancelling Executing Failed Queued Success TimedOut]
	HasPendingInterruptions  bool                             `json:"HasPendingInterruptions"`
	HasWarningsOrErrors      bool                             `json:"HasWarningsOrErrors"`
	ErrorMessage             string                           `json:"ErrorMessage,omitempty"`
	Duration                 string                           `json:"Duration,omitempty"`
	IsCurrent                bool                             `json:"IsCurrent"`
	IsPrevious               bool                             `json:"IsPrevious"`
	IsCompleted              bool                             `json:"IsCompleted"`
	PendingInterruptionTypes []interruptions.InterruptionType `json:"PendingInterruptionTypes,omitempty"`

	resources.Resource
}
