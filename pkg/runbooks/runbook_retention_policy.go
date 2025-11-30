package runbooks

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
)

type RunbookRetentionPolicy struct {
	Strategy       string `json:"Strategy"`
	QuantityToKeep int32  `json:"QuantityToKeep"`
	Unit           string `json:"Unit,omitempty"`
}

const (
	RunbookRetentionStrategyDefault string = "Default"
	RunbookRetentionStrategyForever string = "Forever"
	RunbookRetentionStrategyCount   string = "Count"
)

const (
	RunbookRetentionUnitDays  string = "Days"
	RunbookRetentionUnitItems string = "Items"
)

func NewDefaultRunbookRetentionPolicy() *RunbookRetentionPolicy {
	return &RunbookRetentionPolicy{
		Strategy:       RunbookRetentionStrategyDefault,
		QuantityToKeep: 100,
		Unit:           RunbookRetentionUnitItems,
	}
}

func NewCountBasedRunbookRetentionPolicy(quantityToKeep int32, unit string) (*RunbookRetentionPolicy, error) {
	if quantityToKeep < 1 {
		return nil, internal.CreateInvalidParameterError("NewCountBasedRunbookRetentionPolicy", "quantityToKeep")
	}

	if unit != RunbookRetentionUnitDays && unit != RunbookRetentionUnitItems {
		return nil, internal.CreateInvalidParameterError("NewCountBasedRunbookRetentionPolicy", "unit")
	}

	return &RunbookRetentionPolicy{
		Strategy:       RunbookRetentionStrategyCount,
		QuantityToKeep: quantityToKeep,
		Unit:           unit,
	}, nil
}

func NewKeepForeverRunbookRetentionPolicy() *RunbookRetentionPolicy {
	return &RunbookRetentionPolicy{
		Strategy:       RunbookRetentionStrategyForever,
		QuantityToKeep: 0,
		Unit:           RunbookRetentionUnitItems,
	}
}

// MarshalJSON to handle backward compatibility with older server versions
func (r *RunbookRetentionPolicy) MarshalJSON() ([]byte, error) {
	var fields struct {
		QuantityToKeep    int32  `json:"QuantityToKeep"`
		ShouldKeepForever bool   `json:"ShouldKeepForever"`
		Unit              string `json:"Unit"`
		Strategy          string `json:"Strategy,omitempty"`
	}

	fields.QuantityToKeep = r.QuantityToKeep
	fields.Unit = r.Unit
	fields.Strategy = r.Strategy
	fields.ShouldKeepForever = r.Strategy == RunbookRetentionStrategyForever

	return json.Marshal(fields)
}

// MarshalJSON to handle backward compatibility with older server versions
func (r *RunbookRetentionPolicy) UnmarshalJSON(data []byte) error {
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
	r.Unit = fields.Unit

	// If the Strategy field is present, use it directly
	if fields.Strategy != "" {
		r.Strategy = fields.Strategy
		return nil
	}

	// Infer the Strategy based on other fields for backward compatibility
	if fields.QuantityToKeep == 0 || fields.ShouldKeepForever == true {
		r.Strategy = RunbookRetentionStrategyForever
		return nil
	}

	if fields.QuantityToKeep == 100 && r.Unit == RunbookRetentionUnitItems {
		r.Strategy = RunbookRetentionStrategyDefault
		return nil
	}

	r.Strategy = RunbookRetentionStrategyCount
	return nil
}
