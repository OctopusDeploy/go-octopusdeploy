package filters

type CronScheduledTriggerFilter struct {
	CronExpression string `json:"CronExpression,omitempty"`

	scheduleTriggerFilter
}

func NewCronScheduledTriggerFilter(cronExpression string, timeZone string) *CronScheduledTriggerFilter {
	return &CronScheduledTriggerFilter{
		CronExpression:        cronExpression,
		scheduleTriggerFilter: *newScheduleTriggerFilter(CronExpressionSchedule, timeZone),
	}
}
