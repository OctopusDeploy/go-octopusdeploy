package filters

import (
	"encoding/json"
	"time"
)

type DailyScheduledTriggerFilter struct {
	HourInterval   int16                         `json:"HourInterval,omitempty"`
	Interval       DailyScheduledInterval        `json:"Interval"`
	MinuteInterval int16                         `json:"MinuteInterval,omitempty"`
	RunType        ScheduledTriggerFilterRunType `json:"RunType"`

	scheduledTriggerFilter
}

func NewDailyScheduledTriggerFilter(start time.Time) *DailyScheduledTriggerFilter {
	return &DailyScheduledTriggerFilter{
		scheduledTriggerFilter: *newScheduledTriggerFilter(DailySchedule, start),
	}
}

// UnmarshalJSON sets this DailyScheduledTriggerFilter struct to its representation in JSON.
func (a *DailyScheduledTriggerFilter) UnmarshalJSON(b []byte) error {
	// we don't do anything special here, but because our embedded scheduledTriggerFilter has a custom UnmarshalJSON, we must do it too
	err := json.Unmarshal(b, &a.scheduledTriggerFilter)
	if err != nil {
		return err
	}

	locals := struct {
		HourInterval   int16                         `json:"HourInterval,omitempty"`
		Interval       DailyScheduledInterval        `json:"Interval"`
		MinuteInterval int16                         `json:"MinuteInterval,omitempty"`
		RunType        ScheduledTriggerFilterRunType `json:"RunType"`
	}{}
	err = json.Unmarshal(b, &locals)
	if err != nil {
		return err
	}
	a.HourInterval = locals.HourInterval
	a.Interval = locals.Interval
	a.MinuteInterval = locals.MinuteInterval
	a.RunType = locals.RunType
	return nil
}
