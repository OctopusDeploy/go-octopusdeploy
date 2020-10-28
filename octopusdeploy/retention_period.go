package octopusdeploy

type RetentionPeriod struct {
	Unit              string `json:"Unit"`
	QuantityToKeep    int32  `json:"QuantityToKeep"`
	ShouldKeepForever bool   `json:"ShouldKeepForever"`
}
