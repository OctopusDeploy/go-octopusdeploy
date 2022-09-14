package filters

import (
	"encoding/json"
	"time"
)

type scheduledTriggerFilter struct {
	Start time.Time `json:"StartTime,omitempty"`

	scheduleTriggerFilter
}

const RFC3339NanoNoZone = "2006-01-02T15:04:05.999999999"

// UnmarshalJSON sets this scheduledTriggerFilter struct to its representation in JSON.
func (a *scheduledTriggerFilter) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &a.scheduleTriggerFilter)
	if err != nil {
		return err
	}
	// Octopus Server 2022.3 or newer give us a string like this "2022-09-13T09:00:00.000"
	// The standard golang json unmarshaler does not accept that, because it doesn't handle the milliseconds field.
	type startTimeHolder struct {
		Start string `json:"StartTime,omitempty"`
	}
	var holder startTimeHolder
	err = json.Unmarshal(b, &holder)
	if err != nil {
		return err
	}

	if holder.Start != "" {
		t, err := time.Parse(RFC3339NanoNoZone, holder.Start)
		if err != nil { // fallback
			t, err = time.Parse(time.RFC3339Nano, holder.Start)
		}
		if err != nil {
			return err
		}
		a.Start = t
	}
	return nil
}

func newScheduledTriggerFilter(filterType FilterType, start time.Time) *scheduledTriggerFilter {
	return &scheduledTriggerFilter{
		Start:                 start,
		scheduleTriggerFilter: *newScheduleTriggerFilter(filterType, start.Location().String()),
	}
}

func (t *scheduledTriggerFilter) GetFilterType() FilterType {
	return t.Type
}

func (t *scheduledTriggerFilter) SetFilterType(filterType FilterType) {
	t.Type = filterType
}

var _ ITriggerFilter = &scheduledTriggerFilter{}
