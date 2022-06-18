package core

import (
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/stretchr/testify/assert"
)

func ValidateRequiredSensitiveValue(propertyName string, sensitiveValue *SensitiveValue) error {
	if internal.IsEmpty(propertyName) {
		return internal.CreateInvalidParameterError("ValidateRequiredSensitiveValue", "propertyName")
	}

	if !sensitiveValue.HasValue {
		return fmt.Errorf("%s is a required property; its underlying value is not set", propertyName)
	}

	if len(*sensitiveValue.NewValue) == 0 {
		return fmt.Errorf("%s is a required property; its underlying value is not set", propertyName)
	}

	return nil
}

func TestValidateRequiredSensitiveValue(t *testing.T) {
	newValue := internal.GetRandomName()
	sensitiveValue := NewSensitiveValue(newValue)

	err := ValidateRequiredSensitiveValue("", sensitiveValue)
	assert.Error(t, err)

	err = ValidateRequiredSensitiveValue(" ", sensitiveValue)
	assert.Error(t, err)

	propertyName := internal.GetRandomName()

	err = ValidateRequiredSensitiveValue(propertyName, sensitiveValue)
	assert.NoError(t, err)
}
