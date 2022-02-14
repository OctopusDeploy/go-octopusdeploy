package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"time"
)

type DailyScheduledTriggerFilter struct {
	HourInterval   int16                                       `json:"HourInterval,omitempty"`
	Interval       DailyScheduledInterval                      `json:"Interval"`
	MinuteInterval int16                         `json:"MinuteInterval,omitempty"`
	RunType        ScheduledTriggerFilterRunType `json:"RunType"`

	octopusdeploy.scheduledTriggerFilter
}

func NewDailyScheduledTriggerFilter(start time.Time) *DailyScheduledTriggerFilter {
	return &DailyScheduledTriggerFilter{
		scheduledTriggerFilter: *octopusdeploy.newScheduledTriggerFilter(DailySchedule, start),
	}
}
