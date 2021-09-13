package octopusdeploy

import "time"

type scheduleTriggerFilter struct {
	TimeZone string `json:"Timezone,omitempty"`

	triggerFilter
}

func newScheduleTriggerFilter(filterType FilterType, timeZone *time.Location) *scheduleTriggerFilter {
	return &scheduleTriggerFilter{
		TimeZone:      timeZone.String(),
		triggerFilter: *newTriggerFilter(filterType),
	}
}

func (t *scheduleTriggerFilter) GetFilterType() FilterType {
	return t.Type
}

func (t *scheduleTriggerFilter) SetFilterType(filterType FilterType) {
	t.Type = filterType
}

var _ ITriggerFilter = &scheduleTriggerFilter{}
