package octopusdeploy

// NewSensitiveValue creates and initializes a sensitive value.
func NewSensitiveValue(newValue string) *SensitiveValue {
	if len(newValue) == 0 {
		return &SensitiveValue{
			HasValue: false,
			NewValue: nil,
		}
	}

	return &SensitiveValue{
		HasValue: true,
		NewValue: &newValue,
	}
}

type SensitiveValue struct {
	HasValue bool    `json:"HasValue"`
	NewValue *string `json:"NewValue"`
}
