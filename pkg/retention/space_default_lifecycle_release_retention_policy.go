package retention


type LifecycleReleaseRetentionPolicy struct {
	QuantityToKeep int  `json:"QuantityToKeep"`
	Strategy       string `json:"Strategy"`
	Unit 		 string `json:"Unit"`
	SpaceDefaultRetentionPolicy
}

func CountBasedLifecycleReleaseRetentionPolicy(quantityToKeep int, unit string) *LifecycleReleaseRetentionPolicy {
	return &LifecycleReleaseRetentionPolicy{
		QuantityToKeep: quantityToKeep,
		Strategy:       RetentionStrategyCount,
		Unit:          unit,
	}
}

func KeepForeverLifecycleReleaseRetentionPolicy() *LifecycleReleaseRetentionPolicy {
	return &LifecycleReleaseRetentionPolicy{
		QuantityToKeep: 0,
		Strategy:       RetentionStrategyForever,
		Unit:          RetentionUnitItems,
	}
}
