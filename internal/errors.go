package internal

import "fmt"

func MissingSpaceIDError() error {
	return fmt.Errorf("SpaceID is not found")
}

func CreateBuiltInTeamsCannotDeleteError() error {
	return fmt.Errorf("the built-in teams cannot be deleted")
}

func CreateInvalidParameterError(methodName string, name string) error {
	return fmt.Errorf("%s: the input parameter (%s) is invalid", methodName, name)
}

func CreateInvalidClientStateError(ServiceName string) error {
	return fmt.Errorf("%s: the state of the internal client is invalid", ServiceName)
}

func CreateInvalidPathError(ServiceName string) error {
	return fmt.Errorf("%s: the internal path is not set", ServiceName)
}

func CreateItemNotFoundError(ServiceName string, methodName string, name string) error {
	return fmt.Errorf("%s: the item (%s) via %s was not found", ServiceName, name, methodName)
}

func CreateClientInitializationError(methodName string) error {
	return fmt.Errorf("%s: unable to initialize internal client", methodName)
}

func CreateRequiredParameterIsEmptyOrNilError(parameter string) error {
	return fmt.Errorf("the required parameter, %s is nil or empty", parameter)
}

func CreateResourceNotFoundError(name string, identifier string, value string) error {
	return fmt.Errorf("the service, %s could not find the %s (%s)", name, identifier, value)
}

func CreateValidationFailureError(methodName string, err error) error {
	return fmt.Errorf("validation failure in %s; %v", methodName, err)
}
