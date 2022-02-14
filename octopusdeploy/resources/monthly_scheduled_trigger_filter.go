package resources

import (
	"time"
)

type monthlyScheduledTriggerFilter struct {
	MonthlySchedule MonthlySchedule `json:"MonthlyScheduleType"`

	scheduledTriggerFilter
}

func newMonthlyScheduledTriggerFilter(monthlySchedule MonthlySchedule, start time.Time) *monthlyScheduledTriggerFilter {
	return &monthlyScheduledTriggerFilter{
		MonthlySchedule:        monthlySchedule,
		scheduledTriggerFilter: *newScheduledTriggerFilter(DaysPerMonthSchedule, start),
	}
}
