package filters

import (
	"time"
)

type DateOfMonthScheduledTriggerFilter struct {
	DateOfMonth string `json:"DateOfMonth,omitempty"`

	monthlyScheduledTriggerFilter
}

func NewDateOfMonthScheduledTriggerFilter(dateOfMonth string, start time.Time) *DateOfMonthScheduledTriggerFilter {
	return &DateOfMonthScheduledTriggerFilter{
		DateOfMonth:                   dateOfMonth,
		monthlyScheduledTriggerFilter: *newMonthlyScheduledTriggerFilter(DateOfMonth, start),
	}
}
