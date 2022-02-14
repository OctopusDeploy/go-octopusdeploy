package resources

import (
	"github.com/go-playground/validator/v10"
)

type ProjectTriggerFilter struct {
	DateOfMonth         string   `json:"DateOfMonth"`
	DayNumberOfMonth    string   `json:"DayNumberOfMonth"`
	DayOfWeek           string   `json:"DayOfWeek"`
	EnvironmentIDs      []string `json:"EnvironmentIds,omitempty"`
	EventCategories     []string `json:"EventCategories,omitempty"`
	EventGroups         []string `json:"EventGroups,omitempty"`
	FilterType          string   `json:"FilterType"`
	MonthlyScheduleType string   `json:"MonthlyScheduleType"`
	Roles               []string `json:"Roles"`
	StartTime           string   `json:"StartTime"`
	Timezone            string   `json:"Timezone"`
}

func (f *ProjectTriggerFilter) GetFilterType() FilterType {
	filterType, _ := FilterTypeString(f.FilterType)
	return filterType
}

func (f *ProjectTriggerFilter) SetFilterType(filterType FilterType) {
	f.FilterType = filterType.String()
}

func (f *ProjectTriggerFilter) Validate() error {
	return validator.New().Struct(f)
}

var _ ITriggerFilter = &ProjectTriggerFilter{}
