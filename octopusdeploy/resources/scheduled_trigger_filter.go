package resources

import (
	"time"
)

type scheduledTriggerFilter struct {
	Start time.Time `json:"StartTime,omitempty"`

	scheduleTriggerFilter
}

func newScheduledTriggerFilter(filterType FilterType, start time.Time) *scheduledTriggerFilter {
	return &scheduledTriggerFilter{
		Start:                 start,
		scheduleTriggerFilter: *newScheduleTriggerFilter(filterType, start.Location()),
	}
}

func (t *scheduledTriggerFilter) GetFilterType() FilterType {
	return t.Type
}

func (t *scheduledTriggerFilter) SetFilterType(filterType FilterType) {
	t.Type = filterType
}

var _ ITriggerFilter = &scheduledTriggerFilter{}
