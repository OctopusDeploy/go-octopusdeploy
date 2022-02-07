package octopusdeploy

type triggerFilter struct {
	Type FilterType `json:"FilterType"`

	Resource
}

func newTriggerFilter(filterType FilterType) *triggerFilter {
	return &triggerFilter{
		Type:     filterType,
		Resource: *newResource(),
	}
}
