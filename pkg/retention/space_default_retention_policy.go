package retention

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type SpaceDefaultRetentionPolicy struct {
	SpaceId       string        `json:"SpaceId"`
	Id            string        `json:"Id"`
	Name          string        `json:"Name"`
	RetentionType RetentionType `json:"RetentionType"`
	resources.Resource
}

const (
	RetentionStrategyForever string = "Forever"
	RetentionStrategyCount   string = "Count"
)

const (
	RetentionUnitDays  string = "Days"
	RetentionUnitItems string = "Items"
)

// PolicyResource contains all possible fields whereas Policy just contains
func (policy *SpaceDefaultRetentionPolicy) GetSpaceID() string {
	return policy.SpaceId
}
