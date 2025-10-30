package retention


type LifecycleReleaseRetentionPolicy struct {
	QuantityToKeep int  `json:"QuantityToKeep"`
	Strategy       string `json:"Strategy"`
	Unit 		 string `json:"Unit"`
	SpaceDefaultRetentionPolicy
}
