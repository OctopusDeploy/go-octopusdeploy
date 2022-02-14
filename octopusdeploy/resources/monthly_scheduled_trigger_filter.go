package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"time"
)

type monthlyScheduledTriggerFilter struct {
	MonthlySchedule MonthlySchedule `json:"MonthlyScheduleType"`

	octopusdeploy.scheduledTriggerFilter
}

func newMonthlyScheduledTriggerFilter(monthlySchedule MonthlySchedule, start time.Time) *monthlyScheduledTriggerFilter {
	return &monthlyScheduledTriggerFilter{
		MonthlySchedule:        monthlySchedule,
		scheduledTriggerFilter: *octopusdeploy.newScheduledTriggerFilter(DaysPerMonthSchedule, start),
	}
}
