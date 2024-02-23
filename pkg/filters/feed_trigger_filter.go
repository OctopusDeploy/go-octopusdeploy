package filters

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"

type FeedTriggerFilter struct {
	FeedCategory string                             `json:"FeedCategory,omitempty"`
	Packages     []packages.DeploymentActionPackage `json:"Packages,omitempty"`

	triggerFilter
}

func NewFeedTriggerFilter(feedCategory string, packages []packages.DeploymentActionPackage) *FeedTriggerFilter {
	return &FeedTriggerFilter{
		FeedCategory:  feedCategory,
		Packages:      packages,
		triggerFilter: *newTriggerFilter(FeedFilter),
	}
}

func (t *FeedTriggerFilter) GetFilterType() FilterType {
	return t.Type
}

func (t *FeedTriggerFilter) SetFilterType(filterType FilterType) {
	t.Type = filterType
}

var _ ITriggerFilter = &FeedTriggerFilter{}
