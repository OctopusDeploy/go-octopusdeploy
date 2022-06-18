package feeds

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

type BuiltInFeedStatistics struct {
	TotalPackages int32 `json:"TotalPackages,omitempty"`

	resources.Resource
}
