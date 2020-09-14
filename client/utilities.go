package client

import (
	"fmt"
	"strings"
)

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func createInvalidParameterError(methodName string, parameterName string) error {
	return fmt.Errorf("%s: invalid input parameter, %s", methodName, parameterName)
}

func createInvalidClientStateError(methodName string) error {
	return fmt.Errorf("%s: the state of the internal client is invalid", methodName)
}

func createClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal client", methodName)
}

func createValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("%s: validation failure: %v", methodName, err)
}
