package model

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
