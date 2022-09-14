package filters

import (
	"encoding/json"
	"time"
)

type DaysPerMonthScheduledTriggerFilter struct {
	MonthlySchedule  MonthlySchedule `json:"MonthlyScheduleType"`
	DateOfMonth      string          `json:"DateOfMonth,omitempty"`
	DayNumberOfMonth string          `json:"DayNumberOfMonth,omitempty"`
	Day              *Weekday        `json:"DayOfWeek,omitempty"` // must be a pointer because it's optional

	scheduledTriggerFilter
}

func NewDaysPerMonthScheduledTriggerFilter(monthlySchedule MonthlySchedule, start time.Time) *DaysPerMonthScheduledTriggerFilter {
	return &DaysPerMonthScheduledTriggerFilter{
		MonthlySchedule:        monthlySchedule,
		scheduledTriggerFilter: *newScheduledTriggerFilter(DaysPerMonthSchedule, start),
	}
}

// UnmarshalJSON sets this DaysPerMonthScheduledTriggerFilter struct to its representation in JSON.
func (a *DaysPerMonthScheduledTriggerFilter) UnmarshalJSON(b []byte) error {
	// we don't do anything special here, but because our embedded scheduledTriggerFilter has a custom UnmarshalJSON, we must do it too
	err := json.Unmarshal(b, &a.scheduledTriggerFilter)
	if err != nil {
		return err
	}

	locals := struct {
		MonthlySchedule  MonthlySchedule `json:"MonthlyScheduleType"`
		DateOfMonth      string          `json:"DateOfMonth,omitempty"`
		DayNumberOfMonth string          `json:"DayNumberOfMonth,omitempty"`
		Day              *Weekday        `json:"DayOfWeek,omitempty"`
	}{}
	err = json.Unmarshal(b, &locals)
	if err != nil {
		return err
	}
	a.MonthlySchedule = locals.MonthlySchedule
	a.DateOfMonth = locals.DateOfMonth
	a.DayNumberOfMonth = locals.DayNumberOfMonth
	a.Day = locals.Day
	return nil
}
