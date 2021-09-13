package octopusdeploy

type triggerFilter struct {
	Type FilterType `json:"FilterType"`

	resource
}

func newTriggerFilter(filterType FilterType) *triggerFilter {
	return &triggerFilter{
		Type:     filterType,
		resource: *newResource(),
	}
}
