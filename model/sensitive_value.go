package model

// NewSensitiveValue initializes a SensitiveValue
func NewSensitiveValue(newValue string) SensitiveValue {
	return SensitiveValue{
		HasValue: len(newValue) > 0,
		NewValue: &newValue,
	}
}

type SensitiveValue struct {
	HasValue bool    `json:"HasValue"`
	NewValue *string `json:"NewValue"`
}
