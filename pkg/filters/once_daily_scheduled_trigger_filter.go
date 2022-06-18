package filters

import (
	"time"
)

type OnceDailyScheduledTriggerFilter struct {
	Days []Weekday `json:"DaysOfWeek,omitempty"`

	scheduledTriggerFilter
}

func NewOnceDailyScheduledTriggerFilter(days []Weekday, start time.Time) *OnceDailyScheduledTriggerFilter {
	filter := &OnceDailyScheduledTriggerFilter{
		Days:                   days,
		scheduledTriggerFilter: *newScheduledTriggerFilter(OnceDailySchedule, start),
	}

	return filter
}
