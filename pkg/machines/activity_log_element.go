package machines

import "time"

// ActivityLogElement represents an activity log element.
type ActivityLogElement struct {
	Category    string    `json:"Category,omitempty"`
	Detail      string    `json:"Detail,omitempty"`
	MessageText string    `json:"MessageText,omitempty"`
	OccurredAt  time.Time `json:"OccurredAt,omitempty"`
}
