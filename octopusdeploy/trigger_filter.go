package octopusdeploy

type TriggerFilter struct {
	CronExpression      string   `json:"CronExpression,omitempty"`
	DateOfMonth         string   `json:"DateOfMonth,omitempty"`
	DayNumberOfMonth    string   `json:"DayNumberOfMonth,omitempty"`
	DayOfWeek           string   `json:"DayOfWeek,omitempty" validate:"oneof=Sunday Monday Tuesday Wednesday Thursday Friday Saturday"`
	DaysOfWeek          []string `json:"DaysOfWeek,omitempty"`
	EnvironmentIDs      []string `json:"EnvironmentIds,omitempty"`
	EventCategories     []string `json:"EventCategories,omitempty"`
	EventGroups         []string `json:"EventGroups,omitempty"`
	FilterType          string   `json:"FilterType" validate:"required,oneof=MachineFilter OnceDailySchedule ContinuousDailySchedule DaysPerMonthSchedule CronExpressionSchedule"`
	HourInterval        int      `json:"HourInterval,omitempty"`
	Interval            string   `json:"Interval" validate:"oneof=OnceHourly OnceEveryMinute"`
	MinuteInterval      int      `json:"MinuteInterval,omitempty"`
	MonthlyScheduleType string   `json:"MonthlyScheduleType,omitempty" validate:"oneof=DateOfMonth DayOfMonth"`
	Roles               []string `json:"Roles,omitempty"`
	RunAfter            string   `json:"RunAfter,omitempty"`
	RunUntil            string   `json:"RunUntil,omitempty"`
	StartTime           string   `json:"StartTime,omitempty"`
	Timezone            string   `json:"Timezone,omitempty"`

	resource
}

func NewTriggerFilter() *TriggerFilter {
	return &TriggerFilter{
		resource: *newResource(),
	}
}
