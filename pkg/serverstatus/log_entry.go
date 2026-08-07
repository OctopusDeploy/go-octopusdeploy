package serverstatus

import "time"

type LogEntry struct {
	Category    string     `json:"Category,omitempty"`
	Detail      string     `json:"Detail,omitempty"`
	MessageText string     `json:"MessageText,omitempty"`
	Number      int        `json:"Number"`
	OccurredAt  *time.Time `json:"OccurredAt,omitempty"`
}
