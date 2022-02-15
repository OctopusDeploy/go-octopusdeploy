package service

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createEnvironmentService(t *testing.T) *environmentService {
	service := newEnvironmentService(nil, TestURIEnvironments, TestURIEnvironmentSortOrder, TestURIEnvironmentsSummary)
	testNewService(t, service, TestURIEnvironments, ServiceEnvironmentService)
	return service
}

func TestEnvironmentServiceAdd(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterEnvironment))
	require.Nil(t, resource)

	resource, err = service.Add(&Environment{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestEnvironmentServiceDelete(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)

	err = service.DeleteByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)
}

func TestEnvironmentServiceGetByID(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}

func TestEnvironmentServiceGetBy(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	environments, err := service.GetByName(emptyString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByName, ParameterName))
	assert.NotNil(t, environments)

	environments, err = service.GetByName(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByName, ParameterName))
	assert.NotNil(t, environments)

	environments, err = service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByPartialName, ParameterName))
	require.NotNil(t, environments)
	require.Len(t, environments, 0)

	environments, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByPartialName, ParameterName))
	require.NotNil(t, environments)
	require.Len(t, environments, 0)
}

func TestEnvironmentServiceNew(t *testing.T) {
	ServiceFunction := newEnvironmentService
	client := &sling.Sling{}
	uriTemplate := emptyString
	sortOrderPath := emptyString
	summaryPath := emptyString
	ServiceName := ServiceEnvironmentService

	testCases := []struct {
		name          string
		f             func(*sling.Sling, string, string, string) *environmentService
		client        *sling.Sling
		uriTemplate   string
		sortOrderPath string
		summaryPath   string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, sortOrderPath, summaryPath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, sortOrderPath, summaryPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, sortOrderPath, summaryPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.sortOrderPath, tc.summaryPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
