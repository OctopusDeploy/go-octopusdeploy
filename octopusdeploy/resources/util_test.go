package resources

import (
	"fmt"
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	emptyString      string = ""
	whitespaceString string = " "
)

var propertyName = "fake-property-name"

func getRandomName() string {
	fullName := fmt.Sprintf("test-id %s", uuid.New())
	fullName = fullName[0:44] //Some names in Octopus have a max limit of 50 characters (such as Environment Name)
	return fullName
}

func TestValidateRequiredUUID(t *testing.T) {
	uuidToTest := uuid.Nil

	err := ValidateRequiredUUID(emptyString, &uuidToTest)
	assert.Error(t, err)

	err = ValidateRequiredUUID(whitespaceString, &uuidToTest)
	assert.Error(t, err)

	err = ValidateRequiredUUID(propertyName, &uuidToTest)
	assert.Error(t, err)

	uuidToTest = uuid.New()

	err = ValidateRequiredUUID(propertyName, &uuidToTest)
	assert.NoError(t, err)
}

func TestValidateSemanticVersion(t *testing.T) {
	semanticVersion := emptyString

	err := ValidateSemanticVersion(emptyString, semanticVersion)
	assert.Error(t, err)

	err = ValidateSemanticVersion(whitespaceString, semanticVersion)
	assert.Error(t, err)

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.Error(t, err)

	semanticVersion = "foo"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.Error(t, err)

	semanticVersion = "-1.0.0"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.Error(t, err)

	semanticVersion = "0.0.0"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.NoError(t, err)

	semanticVersion = "1.0.0"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.NoError(t, err)

	semanticVersion = "10.10.10"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.NoError(t, err)

	semanticVersion = "10.-10.10"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.Error(t, err)

	semanticVersion = "v1.0.0"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.Error(t, err)

	semanticVersion = "1.0.0-hello"

	err = ValidateSemanticVersion(propertyName, semanticVersion)
	assert.NoError(t, err)
}

func TestValidateRequiredSensitiveValue(t *testing.T) {
	newValue := getRandomName()
	sensitiveValue := NewSensitiveValue(newValue)

	err := ValidateRequiredSensitiveValue(emptyString, sensitiveValue)
	assert.Error(t, err)

	err = ValidateRequiredSensitiveValue(whitespaceString, sensitiveValue)
	assert.Error(t, err)

	err = ValidateRequiredSensitiveValue(propertyName, sensitiveValue)
	assert.NoError(t, err)
}
