package filters

import (
	"encoding/json"
	"time"
)

type DateOfMonthScheduledTriggerFilter struct {
	DateOfMonth string `json:"DateOfMonth,omitempty"`

	DaysPerMonthScheduledTriggerFilter
}

func NewDateOfMonthScheduledTriggerFilter(dateOfMonth string, start time.Time) *DateOfMonthScheduledTriggerFilter {
	return &DateOfMonthScheduledTriggerFilter{
		DateOfMonth:                        dateOfMonth,
		DaysPerMonthScheduledTriggerFilter: *NewDaysPerMonthScheduledTriggerFilter(DateOfMonth, start),
	}
}

// UnmarshalJSON sets this DateOfMonthScheduledTriggerFilter struct to its representation in JSON.
func (a *DateOfMonthScheduledTriggerFilter) UnmarshalJSON(b []byte) error {
	// we don't do anything special here, but because our embedded DaysPerMonthScheduledTriggerFilter has a custom UnmarshalJSON, we must do it too
	err := json.Unmarshal(b, &a.DaysPerMonthScheduledTriggerFilter)
	if err != nil {
		return err
	}

	locals := struct {
		DateOfMonth string `json:"DateOfMonth,omitempty"`
	}{}
	err = json.Unmarshal(b, &locals)
	if err != nil {
		return err
	}
	a.DateOfMonth = locals.DateOfMonth
	return nil
}
