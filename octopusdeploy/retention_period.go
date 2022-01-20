package octopusdeploy

type RetentionPeriod struct {
	QuantityToKeep    int32  `json:"QuantityToKeep"`
	ShouldKeepForever bool   `json:"ShouldKeepForever"`
	Unit              string `json:"Unit"`
}

func NewRetentionPeriod(quantityToKeep int32, unit string, shouldKeepForever bool) RetentionPeriod {
	return RetentionPeriod{
		QuantityToKeep:    quantityToKeep,
		ShouldKeepForever: shouldKeepForever,
		Unit:              unit,
	}
}
