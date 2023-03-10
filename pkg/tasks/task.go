package tasks

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type Task struct {
	Arguments                  map[string]interface{} `json:"Arguments,omitempty"`
	CanRerun                   bool                   `json:"CanRerun"`
	Completed                  string                 `json:"Completed,omitempty"`
	CompletedTime              *time.Time             `json:"CompletedTime,omitempty"`
	Description                string                 `json:"Description,omitempty"`
	Duration                   string                 `json:"Duration,omitempty"`
	ErrorMessage               string                 `json:"ErrorMessage,omitempty"`
	FinishedSuccessfully       *bool                  `json:"FinishedSuccessfully"`
	HasBeenPickedUpByProcessor bool                   `json:"HasBeenPickedUpByProcessor"`
	HasPendingInterruptions    bool                   `json:"HasPendingInterruptions"`
	HasWarningsOrErrors        bool                   `json:"HasWarningsOrErrors"`
	IsCompleted                *bool                  `json:"IsCompleted"`
	LastUpdatedTime            *time.Time             `json:"LastUpdatedTime,omitempty"`
	Name                       string                 `json:"Name,omitempty"`
	QueueTime                  *time.Time             `json:"QueueTime,omitempty"`
	QueueTimeExpiry            *time.Time             `json:"QueueTimeExpiry,omitempty"`
	ServerNode                 string                 `json:"ServerNode,omitempty"`
	SpaceID                    string                 `json:"SpaceId,omitempty"`
	StartTime                  *time.Time             `json:"StartTime,omitempty"`
	State                      string                 `json:"State,omitempty"`

	resources.Resource
}

// NewTask creates and initializes a task.
func NewTask() *Task {
	return &Task{
		Arguments: map[string]interface{}{},
		Resource:  *resources.NewResource(),
	}
}
