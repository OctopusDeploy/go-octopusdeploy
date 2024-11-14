package filters

type FilterType int

const (
	ContinuousDailySchedule FilterType = iota
	CronExpressionSchedule
	DailySchedule
	DaysPerMonthSchedule
	DaysPerWeekSchedule
	MachineFilter
	OnceDailySchedule
	FeedFilter
	GitFilter
)
