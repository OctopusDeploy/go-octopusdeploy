package filters

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

// ITriggerFilter defines the interface for trigger filters.
type ITriggerFilter interface {
	GetFilterType() FilterType
	SetFilterType(filterType FilterType)
}

type triggerFilter struct {
	Type FilterType `json:"FilterType"`

	resources.Resource
}

func newTriggerFilter(filterType FilterType) *triggerFilter {
	return &triggerFilter{
		Type:     filterType,
		Resource: *resources.NewResource(),
	}
}
