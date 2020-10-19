package model

type RunbookRetentionPeriod struct {
	QuantityToKeep    int32 `json:"QuantityToKeep"`
	ShouldKeepForever bool  `json:"ShouldKeepForever"`
}

func NewRunbookRetentionPeriod() *RunbookRetentionPeriod {
	return &RunbookRetentionPeriod{
		QuantityToKeep:    100,
		ShouldKeepForever: false,
	}
}
