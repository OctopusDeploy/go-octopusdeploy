package dashboard

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type DashboardItem struct {
	ChannelID                string     `json:"ChannelId,omitempty"`
	CompletedTime            *time.Time `json:"CompletedTime,omitempty"`
	Created                  *time.Time `json:"Created,omitempty"`
	DeploymentID             string     `json:"DeploymentId,omitempty"`
	Duration                 string     `json:"Duration,omitempty"`
	EnvironmentID            string     `json:"EnvironmentId,omitempty"`
	ErrorMessage             string     `json:"ErrorMessage,omitempty"`
	HasPendingInterruptions  bool       `json:"HasPendingInterruptions"`
	HasPendingPreconditions  bool       `json:"HasPendingPreconditions"`
	HasWarningsOrErrors      bool       `json:"HasWarningsOrErrors"`
	IsCompleted              bool       `json:"IsCompleted"`
	IsCurrent                bool       `json:"IsCurrent"`
	IsPrevious               bool       `json:"IsPrevious"`
	PendingInterruptionTypes []string   `json:"PendingInterruptionTypes,omitempty"`
	PendingPreconditionTypes []string   `json:"PendingPreconditionTypes,omitempty"`
	ProjectID                string     `json:"ProjectId,omitempty"`
	QueueTime                *time.Time `json:"QueueTime,omitempty"`
	ReleaseID                string     `json:"ReleaseId,omitempty"`
	ReleaseVersion           string     `json:"ReleaseVersion,omitempty"`
	StartTime                *time.Time `json:"StartTime,omitempty"`

	// Enum: [Canceled Cancelling Executing Failed Queued Success TimedOut]
	State    string `json:"State,omitempty"`
	TaskID   string `json:"TaskId,omitempty"`
	TenantID string `json:"TenantId,omitempty"`

	// Resource carries the item's own Id and Links, which the dashboard returns
	// for each item and which callers need to reach the deployment, release and
	// task without reassembling paths from the IDs above.
	resources.Resource
}
