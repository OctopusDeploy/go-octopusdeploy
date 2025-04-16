package core

// NewSensitiveValue creates and initializes a sensitive value.
func NewSensitiveValue(newValue string) *SensitiveValue {
	if len(newValue) > 0 {
		return &SensitiveValue{
			HasValue: true,
			NewValue: &newValue,
		}
	}

	return &SensitiveValue{
		HasValue: false,
	}
}

func (sensitiveValue SensitiveValue) String() string {
	if sensitiveValue.HasValue {
		return *sensitiveValue.NewValue
	}
	return ""
}

type SensitiveValue struct {
	HasValue bool
	Hint     *string
	NewValue *string
}
