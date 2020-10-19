package integration

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	emptyString      string = ""
	whitespaceString string = " "
)

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func createInvalidParameterError(methodName string, parameterName string) error {
	return fmt.Errorf("%s: the input parameter (%s) is invalid", methodName, parameterName)
}

func createInvalidClientStateError(serviceName string) error {
	return fmt.Errorf("%s: the state of the internal client is invalid", serviceName)
}

func createInvalidPathError(serviceName string) error {
	return fmt.Errorf("%s: the internal path is not set", serviceName)
}

func createItemNotFoundError(serviceName string, methodName string, name string) error {
	return fmt.Errorf("%s: the item (%s) via %s was not found", serviceName, name, methodName)
}

func createClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal client", methodName)
}

func createResourceNotFoundError(name string, identifier string, value string) error {
	return fmt.Errorf("the service, %s could not find the %s (%s)", name, identifier, value)
}

func createValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("validation failure in %s; %v", methodName, err)
}

func IsEqualLinks(linksA map[string]string, linksB map[string]string) bool {
	return reflect.DeepEqual(linksA, linksB)
}
