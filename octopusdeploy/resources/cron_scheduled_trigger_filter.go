package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"time"
)

type CronScheduledTriggerFilter struct {
	CronExpression string `json:"CronExpression,omitempty"`

	octopusdeploy.scheduleTriggerFilter
}

func NewCronScheduledTriggerFilter(cronExpression string, timeZone *time.Location) *CronScheduledTriggerFilter {
	return &CronScheduledTriggerFilter{
		CronExpression:        cronExpression,
		scheduleTriggerFilter: *octopusdeploy.newScheduleTriggerFilter(CronExpressionSchedule, timeZone),
	}
}
