package model

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
