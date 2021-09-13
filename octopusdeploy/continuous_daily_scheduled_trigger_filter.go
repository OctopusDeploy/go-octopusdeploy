package octopusdeploy

import (
	"time"
)

type ContinuousDailyScheduledTriggerFilter struct {
	Days           []Weekday               `json:"DaysOfWeek,omitempty"`
	HourInterval   *int16                  `json:"HourInterval,omitempty"`
	Interval       *DailyScheduledInterval `json:"Interval,omitempty"`
	MinuteInterval *int16                  `json:"MinuteInterval,omitempty"`
	RunAfter       *time.Time              `json:"RunAfter,omitempty"`
	RunUntil       *time.Time              `json:"RunUntil,omitempty"`

	scheduleTriggerFilter
}

func NewContinuousDailyScheduledTriggerFilter(days []Weekday, timeZone *time.Location) *ContinuousDailyScheduledTriggerFilter {
	return &ContinuousDailyScheduledTriggerFilter{
		Days:                  days,
		scheduleTriggerFilter: *newScheduleTriggerFilter(ContinuousDailySchedule, timeZone),
	}
}
