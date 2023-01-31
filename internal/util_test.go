package internal

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateRequiredUUID(t *testing.T) {
	uuidToTest := uuid.Nil

	err := ValidateRequiredUUID("", &uuidToTest)
	assert.Error(t, err)

	err = ValidateRequiredUUID(" ", &uuidToTest)
	assert.Error(t, err)

	propertyName := GetRandomName()

	err = ValidateRequiredUUID(propertyName, &uuidToTest)
	assert.Error(t, err)

	uuidToTest = uuid.New()

	err = ValidateRequiredUUID(propertyName, &uuidToTest)
	assert.NoError(t, err)
}

func TestValidateSemanticVersion(t *testing.T) {
	semanticVersion := ""

	err := ValidateSemanticVersion("", semanticVersion)
	assert.Error(t, err)

	err = ValidateSemanticVersion(" ", semanticVersion)
	assert.Error(t, err)

	propertyName := GetRandomName()

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

func TestTrimTemplate(t *testing.T) {
	template := "/api/users{/id}{?skip,take,ids,filter}"
	result := TrimTemplate(template)
	assert.Equal(t, "/api/users", result)

	template = "/api/users/{userId}/apikeys"
	result = TrimTemplate(template)
	assert.Equal(t, "/api/users", result)

}
