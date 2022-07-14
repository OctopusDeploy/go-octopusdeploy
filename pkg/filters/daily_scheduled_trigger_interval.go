package filters

type DailyScheduledInterval int

const (
	OnceDaily DailyScheduledInterval = iota
	OnceHourly
	OnceEveryMinute
)
