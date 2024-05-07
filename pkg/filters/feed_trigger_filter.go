package filters

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"

type FeedTriggerFilter struct {
	Packages []packages.DeploymentActionSlugPackage `json:"Packages,omitempty"`

	triggerFilter
}

func NewFeedTriggerFilter(packages []packages.DeploymentActionSlugPackage) *FeedTriggerFilter {
	return &FeedTriggerFilter{
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
