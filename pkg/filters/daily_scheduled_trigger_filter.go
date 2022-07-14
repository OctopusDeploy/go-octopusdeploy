package filters

import "time"

type DailyScheduledTriggerFilter struct {
	HourInterval   int16                         `json:"HourInterval,omitempty"`
	Interval       DailyScheduledInterval        `json:"Interval"`
	MinuteInterval int16                         `json:"MinuteInterval,omitempty"`
	RunType        ScheduledTriggerFilterRunType `json:"RunType"`

	scheduledTriggerFilter
}

func NewDailyScheduledTriggerFilter(start time.Time) *DailyScheduledTriggerFilter {
	return &DailyScheduledTriggerFilter{
		scheduledTriggerFilter: *newScheduledTriggerFilter(DailySchedule, start),
	}
}
