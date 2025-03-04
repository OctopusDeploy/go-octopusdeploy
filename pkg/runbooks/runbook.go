package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Runbook struct {
	ConnectivityPolicy         *core.ConnectivityPolicy    `json:"ConnectivityPolicy,omitempty"`
	DefaultGuidedFailureMode   string                      `json:"DefaultGuidedFailureMode" validate:"required,oneof=EnvironmentDefault Off On"`
	Description                string                      `json:"Description,omitempty"`
	EnvironmentScope           string                      `json:"EnvironmentScope" validate:"required,oneof=All FromProjectLifecycles Specified"`
	Environments               []string                    `json:"Environments,omitempty"`
	MultiTenancyMode           core.TenantedDeploymentMode `json:"MultiTenancyMode" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	Name                       string                      `json:"Name,omitempty"`
	ProjectID                  string                      `json:"ProjectId,omitempty"`
	PublishedRunbookSnapshotID string                      `json:"PublishedRunbookSnapshotId,omitempty"`
	RunRetentionPolicy         *RunbookRetentionPeriod     `json:"RunRetentionPolicy,omitempty"`
	RunbookProcessID           string                      `json:"RunbookProcessId,omitempty"`
	SpaceID                    string                      `json:"SpaceId,omitempty"`
	ForcePackageDownload       bool                        `json:"ForcePackageDownload"`

	resources.Resource
}

type CreatedRunbook struct {
	ID        string `json:"Id,omitempty"`
	ProjectID string `json:"ProjectId,omitempty"`
	Name      string `json:"Name,omitempty"`
	Slug      string `json:"Slug,omitempty"`
	GitRef    string `json:"GitRef,omitempty"`
}

// NewRunbook creates and initializes a runbook.
func NewRunbook(name string, projectID string) *Runbook {
	return &Runbook{
		DefaultGuidedFailureMode: "EnvironmentDefault",
		EnvironmentScope:         "All",
		MultiTenancyMode:         core.TenantedDeploymentModeUntenanted,
		Name:                     name,
		ProjectID:                projectID,
		RunRetentionPolicy:       NewRunbookRetentionPeriod(),
		Resource:                 *resources.NewResource(),
	}
}

// Validate checks the state of the runbook and returns an error if invalid.
func (r *Runbook) Validate() error {
	return validator.New().Struct(r)
}
