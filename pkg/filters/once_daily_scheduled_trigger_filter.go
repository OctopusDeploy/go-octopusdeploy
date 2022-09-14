package filters

import (
	"encoding/json"
	"time"
)

type OnceDailyScheduledTriggerFilter struct {
	Days []Weekday `json:"DaysOfWeek,omitempty"`

	scheduledTriggerFilter
}

func NewOnceDailyScheduledTriggerFilter(days []Weekday, start time.Time) *OnceDailyScheduledTriggerFilter {
	filter := &OnceDailyScheduledTriggerFilter{
		Days:                   days,
		scheduledTriggerFilter: *newScheduledTriggerFilter(OnceDailySchedule, start),
	}

	return filter
}

// UnmarshalJSON sets this OnceDailyScheduledTriggerFilter struct to its representation in JSON.
func (a *OnceDailyScheduledTriggerFilter) UnmarshalJSON(b []byte) error {
	// we don't do anything special here, but because our embedded scheduledTriggerFilter has a custom UnmarshalJSON, we must do it too
	err := json.Unmarshal(b, &a.scheduledTriggerFilter)
	if err != nil {
		return err
	}

	locals := struct {
		Days []Weekday `json:"DaysOfWeek,omitempty"`
	}{}
	err = json.Unmarshal(b, &locals)
	if err != nil {
		return err
	}
	a.Days = locals.Days
	return nil
}
