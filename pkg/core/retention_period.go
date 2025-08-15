package core

import (
	"encoding/json"
)

type RetentionPeriod struct {
	QuantityToKeep    int32  `json:"QuantityToKeep"`
	ShouldKeepForever bool   `json:"ShouldKeepForever"`
	Unit              string `json:"Unit"`
	Strategy          string `json:"Strategy,omitempty"`
}

const (
	RetentionStrategyDefault string = "Default"
	RetentionStrategyForever string = "Forever"
	RetentionStrategyCount   string = "Count"
)

const (
	RetentionUnitDays  string = "Days"
	RetentionUnitItems string = "Items"
)

func NewRetentionPeriod(quantityToKeep int32, unit string, shouldKeepForever bool) *RetentionPeriod {
	if shouldKeepForever {
		return KeepForeverRetentionPeriod()
	} else {
		return CountBasedRetentionPeriod(quantityToKeep, unit)
	}
}

func CountBasedRetentionPeriod(quantityToKeep int32, unit string) *RetentionPeriod {
	return &RetentionPeriod{
		QuantityToKeep:    quantityToKeep,
		Unit:              unit,
		ShouldKeepForever: false,
		Strategy:          RetentionStrategyCount,
	}
}

func KeepForeverRetentionPeriod() *RetentionPeriod {
	return &RetentionPeriod{
		QuantityToKeep:    0,
		Unit:              RetentionUnitItems,
		ShouldKeepForever: true,
		Strategy:          RetentionStrategyForever,
	}
}

func SpaceDefaultRetentionPeriod() *RetentionPeriod {
	return &RetentionPeriod{
		QuantityToKeep:    0,
		Unit:              RetentionUnitItems,
		ShouldKeepForever: true,
		Strategy:          RetentionStrategyDefault,
	}
}

// UnmarshalJSON sets a retention period to its representation in JSON.
func (r *RetentionPeriod) UnmarshalJSON(data []byte) error {
	var fields struct {
		QuantityToKeep    int32  `json:"QuantityToKeep"`
		ShouldKeepForever bool   `json:"ShouldKeepForever"`
		Unit              string `json:"Unit"`
		Strategy          string `json:"Strategy,omitempty"`
	}

	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	r.QuantityToKeep = fields.QuantityToKeep
	r.ShouldKeepForever = fields.ShouldKeepForever
	r.Unit = fields.Unit

	if fields.Strategy == "" {
		if r.ShouldKeepForever == true {
			r.Strategy = RetentionStrategyForever
		} else {
			r.Strategy = RetentionStrategyCount
		}
	} else {
		r.Strategy = fields.Strategy
	}

	return nil
}
