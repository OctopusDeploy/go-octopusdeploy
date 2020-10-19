package model

import (
	"time"
)

// ActivityLogEntry represents an activity log entry.
type ActivityLogEntry struct {
	Category    string    `json:"Category,omitempty" validate:"required,oneof=Abandoned Alert Error Fatal Finished Gap Highlight Info Planned Trace Updated Verbose Wait Warning"`
	Detail      string    `json:"Detail,omitempty"`
	MessageText string    `json:"MessageText,omitempty"`
	OccurredAt  time.Time `json:"OccurredAt,omitempty"`
}
