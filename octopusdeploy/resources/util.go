package resources

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// ValidateStringInSlice checks if a string is in the given slice
func ValidateStringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}

	return false
}

// ValidatePropertyValues returns an error if the given string is not in a slice of strings
func ValidatePropertyValues(propertyName string, propertyValue string, validValues []string) error {
	if internal.IsEmpty(propertyName) {
		return internal.CreateInvalidParameterError("ValidatePropertyValues", "propertyName")
	}

	if ValidateStringInSlice(propertyValue, validValues) {
		return nil
	}

	return fmt.Errorf("%s must be one of \"%v\"", propertyName, strings.Join(validValues, ","))
}

// ValidateRequiredPropertyValue returns an error if the property value is empty
func ValidateRequiredPropertyValue(propertyName string, propertyValue string) error {
	if internal.IsEmpty(propertyName) {
		return internal.CreateInvalidParameterError("ValidateRequiredPropertyValue", "propertyName")
	}

	if len(propertyValue) > 0 {
		return nil
	}

	return fmt.Errorf("%s is a required property and cannot be empty", propertyName)
}

func ValidateRequiredUUID(propertyName string, id *uuid.UUID) error {
	if internal.IsEmpty(propertyName) {
		return internal.CreateInvalidParameterError("ValidateRequiredUUID", "propertyName")
	}

	if id == nil {
		return internal.CreateInvalidParameterError("ValidateRequiredUUID", "id")
	}

	if *id == uuid.Nil {
		return fmt.Errorf("%s is a required property; its value is an empty UUID", propertyName)
	}

	return nil
}

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

// ValidateMultipleProperties returns the first error in a list of property validations
func ValidateMultipleProperties(validatePropertyErrors []error) error {
	for _, check := range validatePropertyErrors {
		if check != nil {
			return check
		}
	}

	return nil
}

// ValidatePropertiesMatch checks two values against each other
func ValidatePropertiesMatch(firstProperty, firstPropertyName, secondProperty, secondPropertyName string) error {
	if firstProperty != secondProperty {
		return fmt.Errorf("%s and %s must match. They are currently %s and %s", firstPropertyName, secondPropertyName, firstProperty, secondProperty)
	}

	return nil
}

func ValidateSemanticVersion(propertyName string, version string) error {
	if internal.IsEmpty(propertyName) {
		return internal.CreateInvalidParameterError("ValidateSemanticVersion", "propertyName")
	}

	re := regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if re.MatchString(version) {
		return nil
	}

	return fmt.Errorf("%s is must be a semantic version string", propertyName)
}
