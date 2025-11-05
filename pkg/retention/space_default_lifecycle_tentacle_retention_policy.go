package retention

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type LifecycleTentacleRetentionPolicy struct {
	QuantityToKeep int    `json:"QuantityToKeep"`
	Strategy       string `json:"Strategy"`
	Unit           string `json:"Unit"`
	SpaceDefaultRetentionPolicy
}

func NewCountBasedLifecycleTentacleRetentionPolicy(quantityToKeep int, unit string, spaceId string, policyId string) *LifecycleTentacleRetentionPolicy {
	return &LifecycleTentacleRetentionPolicy{
		QuantityToKeep: quantityToKeep,
		Strategy:       RetentionStrategyCount,
		Unit:           unit,
		SpaceDefaultRetentionPolicy: SpaceDefaultRetentionPolicy{
			SpaceId:       spaceId,
			RetentionType: LifecycleTentacleRetentionType,
			Resource: resources.Resource{
				ID: policyId,
			},
		},
	}
}

func NewKeepForeverLifecycleTentacleRetentionPolicy(spaceId string, policyId string) *LifecycleTentacleRetentionPolicy {
	return &LifecycleTentacleRetentionPolicy{
		QuantityToKeep: 0,
		Strategy:       RetentionStrategyForever,
		Unit:           RetentionUnitItems,
		SpaceDefaultRetentionPolicy: SpaceDefaultRetentionPolicy{
			SpaceId:       spaceId,
			RetentionType: LifecycleTentacleRetentionType,
			Resource: resources.Resource{
				ID: policyId,
			},
		},
	}
}
