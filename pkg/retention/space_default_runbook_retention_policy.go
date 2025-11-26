package retention

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type RunbookRetentionPolicy struct {
	QuantityToKeep int    `json:"QuantityToKeep"`
	Strategy       string `json:"Strategy"`
	Unit           string `json:"Unit"`
	SpaceDefaultRetentionPolicy
}

func NewCountBasedRunbookRetentionPolicy(quantityToKeep int, unit string, spaceId string, policyId string) *RunbookRetentionPolicy {
	return &RunbookRetentionPolicy{
		QuantityToKeep: quantityToKeep,
		Strategy:       RetentionStrategyCount,
		Unit:           unit,
		SpaceDefaultRetentionPolicy: SpaceDefaultRetentionPolicy{
			SpaceId:       spaceId,
			RetentionType: RunbookRetentionType,
			Resource: resources.Resource{
				ID: policyId,
			},
		},
	}
}

func NewKeepForeverRunbookRetentionPolicy(spaceId string, policyId string) *RunbookRetentionPolicy {
	return &RunbookRetentionPolicy{
		QuantityToKeep: 0,
		Strategy:       RetentionStrategyForever,
		Unit:           RetentionUnitItems,
		SpaceDefaultRetentionPolicy: SpaceDefaultRetentionPolicy{
			SpaceId:       spaceId,
			RetentionType: RunbookRetentionType,
			Resource: resources.Resource{
				ID: policyId,
			},
		},
	}
}
