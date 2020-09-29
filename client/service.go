package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

const (
	emptyString      string = ""
	whitespaceString string = " "
)

// ServiceInterface defines the contract for all services that communicate with
// the Octopus API.
type ServiceInterface interface {
	getClient() *sling.Sling
	getName() string
	getURITemplate() *uritemplates.UriTemplate
}

func getAddPath(s ServiceInterface, r model.ResourceInterface) (string, error) {
	if isNil(r) {
		return emptyString, createInvalidParameterError(operationAdd, parameterResource)
	}

	err := r.Validate()
	if err != nil {
		return emptyString, createValidationFailureError(operationAdd, err)
	}

	err = validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	return s.getURITemplate().Expand(values)
}

func getPath(s ServiceInterface) (string, error) {
	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	return s.getURITemplate().Expand(values)
}

func getAllPath(s ServiceInterface) (string, error) {
	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	path, err := s.getURITemplate().Expand(values)
	if err != nil {
		return emptyString, err
	}

	return path + "/all", nil
}

func getByIDPath(s ServiceInterface, id string) (string, error) {
	if isEmpty(id) {
		return emptyString, createInvalidParameterError(operationGetByID, parameterID)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterID] = id

	return s.getURITemplate().Expand(values)
}

func getByIDsPath(s ServiceInterface, ids []string) (string, error) {
	if len(ids) == 0 {
		return s.getURITemplate().Expand(make(map[string]interface{}))
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	idValues := emptyString

	for i := 0; i < len(ids); i++ {
		idValues += ids[i]
		if i < len(ids)-1 {
			idValues += ","
		}
	}

	values := make(map[string]interface{})
	values[parameterIDs] = idValues

	return s.getURITemplate().Expand(values)
}

func getByNamePath(s ServiceInterface, name string) (string, error) {
	if isEmpty(name) {
		return emptyString, createInvalidParameterError(operationGetByName, parameterName)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterName] = name

	return s.getURITemplate().Expand(values)
}

func getByPartialNamePath(s ServiceInterface, name string) (string, error) {
	if isEmpty(name) {
		return emptyString, createInvalidParameterError(operationGetByPartialName, parameterName)
	}

	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterPartialName] = name

	return s.getURITemplate().Expand(values)
}

func getByAccountTypePath(s ServiceInterface, accountType enum.AccountType) (string, error) {
	err := validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterAccountType] = accountType

	return s.getURITemplate().Expand(values)
}

func deleteByID(s ServiceInterface, id string) error {
	if isEmpty(id) {
		return createInvalidParameterError(operationDeleteByID, parameterID)
	}

	err := validateInternalState(s)
	if err != nil {
		return err
	}

	values := make(map[string]interface{})
	values[parameterID] = id

	path, err := s.getURITemplate().Expand(values)
	if err != nil {
		return err
	}

	return apiDelete(s.getClient(), path)
}

func getUpdatePath(s ServiceInterface, r model.ResourceInterface) (string, error) {
	if isNil(r) {
		return emptyString, createInvalidParameterError(operationUpdate, parameterResource)
	}

	err := r.Validate()
	if err != nil {
		return emptyString, createValidationFailureError(operationUpdate, err)
	}

	err = validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterID] = r.GetID()
	return s.getURITemplate().Expand(values)
}

func validateInternalState(s ServiceInterface) error {
	if s.getClient() == nil {
		return createInvalidClientStateError(s.getName())
	}

	values := make(map[string]interface{})
	path, err := s.getURITemplate().Expand(values)

	if isEmpty(path) {
		return createInvalidPathError(s.getName())
	}

	return err
}
