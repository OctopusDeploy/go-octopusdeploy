package model

// NewSensitiveValue initializes a SensitiveValue
func NewSensitiveValue(newValue string) (*SensitiveValue, error) {
	sensitiveValue := &SensitiveValue{
		HasValue: len(newValue) > 0,
		NewValue: &newValue,
	}

	return sensitiveValue, nil
}

type SensitiveValue struct {
	HasValue bool    `json:"HasValue"`
	NewValue *string `json:"NewValue"`
}
