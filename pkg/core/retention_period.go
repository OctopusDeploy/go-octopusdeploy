package core

type RetentionPeriod struct {
	QuantityToKeep    int32  `json:"QuantityToKeep"`
	ShouldKeepForever bool   `json:"ShouldKeepForever"`
	Unit              string `json:"Unit"`
	Strategy          string `json:"Strategy"`
}

func NewRetentionPeriod(strategy string, quantityToKeep int32, unit string, shouldKeepForever bool) *RetentionPeriod {
	return &RetentionPeriod{
		Strategy:          strategy,
		QuantityToKeep:    0,
		ShouldKeepForever: false,
		Unit:              "item",
	}
}
