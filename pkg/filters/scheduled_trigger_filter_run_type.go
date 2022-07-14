package filters

type ScheduledTriggerFilterRunType int

const (
	ScheduledTime ScheduledTriggerFilterRunType = iota
	Continuously
)
