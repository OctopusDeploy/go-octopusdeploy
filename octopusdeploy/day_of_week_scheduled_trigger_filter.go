package octopusdeploy

import (
	"time"
)

type DayOfWeekScheduledTriggerFilter struct {
	Day     Weekday `json:"DayOfWeek,omitempty"`
	Ordinal string  `json:"DayNumberOfMonth,omitempty"`

	monthlyScheduledTriggerFilter
}

func NewDayOfWeekScheduledTriggerFilter(ordinal string, day Weekday, start time.Time) *DayOfWeekScheduledTriggerFilter {
	return &DayOfWeekScheduledTriggerFilter{
		Day:                           day,
		Ordinal:                       ordinal,
		monthlyScheduledTriggerFilter: *newMonthlyScheduledTriggerFilter(DayOfMonth, start),
	}
}
