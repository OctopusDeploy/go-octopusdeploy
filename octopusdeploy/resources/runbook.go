package resources

import (
	"github.com/go-playground/validator/v10"
)

type Runbook struct {
	ConnectivityPolicy         *ConnectivityPolicy     `json:"ConnectivityPolicy,omitempty"`
	DefaultGuidedFailureMode   string                  `json:"DefaultGuidedFailureMode" validate:"required,oneof=EnvironmentDefault Off On"`
	Description                string                  `json:"Description,omitempty"`
	EnvironmentScope           string                  `json:"EnvironmentScope" validate:"required,oneof=All FromProjectLifecycles Specified"`
	Environments               []string                `json:"Environments,omitempty"`
	MultiTenancyMode           string                  `json:"MultiTenancyMode" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	Name                       string                  `json:"Name,omitempty"`
	ProjectID                  string                  `json:"ProjectId,omitempty"`
	PublishedRunbookSnapshotID string                  `json:"PublishedRunbookSnapshotId,omitempty"`
	RunRetentionPolicy         *RunbookRetentionPeriod `json:"RunRetentionPolicy,omitempty"`
	RunbookProcessID           string                  `json:"RunbookProcessId,omitempty"`
	SpaceID                    string                  `json:"SpaceId,omitempty"`

	Resource
}

// NewRunbook creates and initializes a runbook.
func NewRunbook(name string, projectID string) *Runbook {
	return &Runbook{
		DefaultGuidedFailureMode: "EnvironmentDefault",
		EnvironmentScope:         "All",
		MultiTenancyMode:         "Untenanted",
		Name:                     name,
		ProjectID:                projectID,
		RunRetentionPolicy:       NewRunbookRetentionPeriod(),
		Resource:                 *NewResource(),
	}
}

// Validate checks the state of the runbook and returns an error if invalid.
func (r *Runbook) Validate() error {
	return validator.New().Struct(r)
}
