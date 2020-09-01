package model

type RetentionPeriod struct {
	Unit              RetentionUnit `json:"Unit"`
	QuantityToKeep    int32         `json:"QuantityToKeep"`
	ShouldKeepForever bool          `json:"ShouldKeepForever"`
}
