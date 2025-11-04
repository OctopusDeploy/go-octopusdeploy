package retention

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type SpaceDefaultRetentionPolicy struct {
	SpaceId       string        `json:"SpaceId"`
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

type SpaceDefaultRetentionPolicyResource struct {
	SpaceDefaultRetentionPolicy
	QuantityToKeep int    `json:"QuantityToKeep"`
	Strategy       string `json:"Strategy"`
	Unit           string `json:"Unit"`
}

type ISpaceDefaultRetentionPolicy interface {
	GetSpaceID() string
	resources.IResource
}

func (policy *SpaceDefaultRetentionPolicy) GetSpaceID() string {
	return policy.SpaceId
}
