package model

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/enum"
)

type ActivityLogEntry struct {
	Category    enum.ActivityLogCategory `json:"Category,omitempty"`
	Detail      string                   `json:"Detail,omitempty"`
	MessageText string                   `json:"MessageText,omitempty"`
	OccurredAt  time.Time                `json:"OccurredAt,omitempty"`
}
