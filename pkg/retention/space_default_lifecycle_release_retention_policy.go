package retention

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type LifecycleReleaseRetentionPolicy struct {
	QuantityToKeep int    `json:"QuantityToKeep"`
	Strategy       string `json:"Strategy"`
	Unit           string `json:"Unit"`
	SpaceDefaultRetentionPolicy
}

func NewCountBasedLifecycleReleaseRetentionPolicy(quantityToKeep int, unit string, spaceId string, policyId string) *LifecycleReleaseRetentionPolicy {
	return &LifecycleReleaseRetentionPolicy{
		QuantityToKeep: quantityToKeep,
		Strategy:       RetentionStrategyCount,
		Unit:           unit,
		SpaceDefaultRetentionPolicy: SpaceDefaultRetentionPolicy{
			SpaceId:       spaceId,
			RetentionType: LifecycleReleaseRetentionType,
			Resource: resources.Resource{
				ID: policyId,
			},
		},
	}
}

func KeepForeverLifecycleReleaseRetentionPolicy(spaceId string, policyId string) *LifecycleReleaseRetentionPolicy {
	return &LifecycleReleaseRetentionPolicy{
		QuantityToKeep: 0,
		Strategy:       RetentionStrategyForever,
		Unit:           RetentionUnitItems,
		SpaceDefaultRetentionPolicy: SpaceDefaultRetentionPolicy{
			SpaceId:       spaceId,
			RetentionType: LifecycleReleaseRetentionType,
			Resource: resources.Resource{
				ID: policyId,
			},
		},
	}
}
