package filters

import "time"

type CronScheduledTriggerFilter struct {
	CronExpression string `json:"CronExpression,omitempty"`

	scheduleTriggerFilter
}

func NewCronScheduledTriggerFilter(cronExpression string, timeZone *time.Location) *CronScheduledTriggerFilter {
	return &CronScheduledTriggerFilter{
		CronExpression:        cronExpression,
		scheduleTriggerFilter: *newScheduleTriggerFilter(CronExpressionSchedule, timeZone),
	}
}
