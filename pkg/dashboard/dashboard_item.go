package dashboard

import "time"

type DashboardItem struct {
	ChannelID               string     `json:"ChannelId,omitempty"`
	CompletedTime           *time.Time `json:"CompletedTime,omitempty"`
	Created                 *time.Time `json:"Created,omitempty"`
	DeploymentID            string     `json:"DeploymentId,omitempty"`
	Duration                string     `json:"Duration,omitempty"`
	EnvironmentID           string     `json:"EnvironmentId,omitempty"`
	ErrorMessage            string     `json:"ErrorMessage,omitempty"`
	HasPendingInterruptions bool       `json:"HasPendingInterruptions"`
	HasWarningsOrErrors     bool       `json:"HasWarningsOrErrors"`
	IsCompleted             bool       `json:"IsCompleted"`
	IsCurrent               bool       `json:"IsCurrent"`
	IsPrevious              bool       `json:"IsPrevious"`
	ProjectID               string     `json:"ProjectId,omitempty"`
	QueueTime               *time.Time `json:"QueueTime,omitempty"`
	ReleaseID               string     `json:"ReleaseId,omitempty"`
	ReleaseVersion          string     `json:"ReleaseVersion,omitempty"`
	StartTime               *time.Time `json:"StartTime,omitempty"`

	// Enum: [Canceled Cancelling Executing Failed Queued Success TimedOut]
	State    string `json:"State,omitempty"`
	TaskID   string `json:"TaskId,omitempty"`
	TenantID string `json:"TenantId,omitempty"`
}
