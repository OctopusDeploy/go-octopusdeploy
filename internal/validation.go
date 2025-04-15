package internal

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/google/uuid"
	"regexp"
	"strings"
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
	if IsEmpty(propertyName) {
		return CreateInvalidParameterError("ValidatePropertyValues", "propertyName")
	}

	if ValidateStringInSlice(propertyValue, validValues) {
		return nil
	}

	return fmt.Errorf("%s must be one of \"%v\"", propertyName, strings.Join(validValues, ","))
}

// ValidateRequiredPropertyValue returns an error if the property value is empty
func ValidateRequiredPropertyValue(propertyName string, propertyValue string) error {
	if IsEmpty(propertyName) {
		return CreateInvalidParameterError("ValidateRequiredPropertyValue", "propertyName")
	}

	if len(propertyValue) > 0 {
		return nil
	}

	return fmt.Errorf("%s is a required property and cannot be empty", propertyName)
}

func ValidateRequiredUUID(propertyName string, id *uuid.UUID) error {
	if IsEmpty(propertyName) {
		return CreateInvalidParameterError("ValidateRequiredUUID", "propertyName")
	}

	if id == nil {
		return CreateInvalidParameterError("ValidateRequiredUUID", "id")
	}

	if *id == uuid.Nil {
		return fmt.Errorf("%s is a required property; its value is an empty UUID", propertyName)
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
	if IsEmpty(propertyName) {
		return CreateInvalidParameterError("ValidateSemanticVersion", "propertyName")
	}

	re := regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if re.MatchString(version) {
		return nil
	}

	return fmt.Errorf("%s is must be a semantic version string", propertyName)
}

func ValidateUsernamePasswordProperties(username string, password *core.SensitiveValue) error {
	if IsEmpty(username) && password != nil {
		return CreateRequiredParameterIsEmptyOrNilError("username")
	}

	if !IsEmpty(username) && password == nil {
		return CreateRequiredParameterIsEmptyOrNilError("password")
	}

	return nil
}
