package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"time"
)

type ContinuousDailyScheduledTriggerFilter struct {
	Days           []Weekday `json:"DaysOfWeek,omitempty"`
	HourInterval   *int16    `json:"HourInterval,omitempty"`
	Interval       *DailyScheduledInterval `json:"Interval,omitempty"`
	MinuteInterval *int16                  `json:"MinuteInterval,omitempty"`
	RunAfter       *time.Time              `json:"RunAfter,omitempty"`
	RunUntil       *time.Time              `json:"RunUntil,omitempty"`

	octopusdeploy.scheduleTriggerFilter
}

func NewContinuousDailyScheduledTriggerFilter(days []Weekday, timeZone *time.Location) *ContinuousDailyScheduledTriggerFilter {
	return &ContinuousDailyScheduledTriggerFilter{
		Days:                  days,
		scheduleTriggerFilter: *octopusdeploy.newScheduleTriggerFilter(ContinuousDailySchedule, timeZone),
	}
}
