package model

import (
	"errors"
	"strings"
)

// NewSensitiveValue initializes a SensitiveValue
func NewSensitiveValue(newValue string) (*SensitiveValue, error) {
	if len(strings.Trim(newValue, " ")) == 0 {
		return nil, errors.New("NewSensitiveValue: invalid newValue")
	}

	return &SensitiveValue{
		HasValue: false,
		NewValue: &newValue,
	}, nil
}

type SensitiveValue struct {
	HasValue bool    `json:"HasValue"`
	NewValue *string `json:"NewValue"`
}
