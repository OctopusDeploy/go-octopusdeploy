package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createEnvironmentService(t *testing.T) *environmentService {
	service := newEnvironmentService(nil, TestURIEnvironments, TestURIEnvironmentSortOrder, TestURIEnvironmentsSummary)
	testNewService(t, service, TestURIEnvironments, serviceEnvironmentService)
	return service
}

func CreateTestEnvironment(t *testing.T, service *environmentService) *Environment {
	if service == nil {
		service = createEnvironmentService(t)
	}
	require.NotNil(t, service)

	allowDynamicInfrastructure := createRandomBoolean()
	name := getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	useGuidedFailure := createRandomBoolean()

	environment := NewEnvironment(name)
	environment.AllowDynamicInfrastructure = allowDynamicInfrastructure
	environment.Description = description
	environment.UseGuidedFailure = useGuidedFailure

	require.NoError(t, environment.Validate())

	createdEnvironment, err := service.Add(environment)
	require.NoError(t, err)
	require.NotNil(t, createdEnvironment)
	require.NotEmpty(t, createdEnvironment.GetID())

	return createdEnvironment
}

func DeleteTestEnvironment(t *testing.T, service *environmentService, artifact *Environment) error {
	if service == nil {
		service = createEnvironmentService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(artifact.GetID())
}

func IsEqualEnvironments(t *testing.T, expected *Environment, actual *Environment) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.Equal(t, expected.Links, actual.Links)
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))
	assert.Equal(t, expected.Resource, actual.Resource)

	// Environment
	assert.Equal(t, expected.AllowDynamicInfrastructure, actual.AllowDynamicInfrastructure)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.SortOrder, actual.SortOrder)
	assert.Equal(t, expected.UseGuidedFailure, actual.UseGuidedFailure)
}

func UpdateEnvironment(t *testing.T, service *environmentService, environment *Environment) *Environment {
	if service == nil {
		service = createEnvironmentService(t)
	}
	require.NotNil(t, service)

	updatedEnvironment, err := service.Update(environment)
	assert.NoError(t, err)
	require.NotNil(t, updatedEnvironment)

	return updatedEnvironment
}

func TestEnvironmentServiceAdd(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterEnvironment))
	require.Nil(t, resource)

	// TODO: test this call; it should NOT send anything via HTTP
	resource, err = service.Add(&Environment{})
	require.Error(t, err)
	require.Nil(t, resource)

	environment := CreateTestEnvironment(t, service)
	err = DeleteTestEnvironment(t, service, environment)
	require.NoError(t, err)
}

func TestEnvironmentServiceAddGetDelete(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterEnvironment))
	require.Nil(t, resource)

	// TODO: test this call; it should NOT send anything via HTTP
	resource, err = service.Add(&Environment{})
	require.Error(t, err)
	require.Nil(t, resource)

	environment := CreateTestEnvironment(t, service)

	environment, err = service.GetByID(environment.ID)
	require.NoError(t, err)
	require.NotNil(t, environment)

	err = DeleteTestEnvironment(t, service, environment)
	require.NoError(t, err)
}

func TestEnvironmentServiceDelete(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	require.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
}

func TestEnvironmentServiceDeleteAll(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	environments, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, environments)

	for _, environment := range environments {
		err = DeleteTestEnvironment(t, service, environment)
		require.NoError(t, err)
	}
}

func TestEnvironmentServiceGetAll(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	// create 30 test environments (to be deleted)
	for i := 0; i < 30; i++ {
		environment := CreateTestEnvironment(t, service)
		require.NotNil(t, environment)
	}

	environments, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, environments)

	for _, environment := range environments {
		err = DeleteTestEnvironment(t, service, environment)
		require.NoError(t, err)
	}
}

func TestEnvironmentServiceGetByID(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	require.Nil(t, resource)

	environments, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, environments)

	for _, environment := range environments {
		environmentToCompare, err := service.GetByID(environment.GetID())
		require.NoError(t, err)
		IsEqualEnvironments(t, environment, environmentToCompare)
	}
}

func TestEnvironmentServiceGetByName(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	environments, err := service.GetByName(emptyString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByName, parameterName))
	assert.NotNil(t, environments)

	environments, err = service.GetByName(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByName, parameterName))
	assert.NotNil(t, environments)

	environments, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, environments)

	for _, environment := range environments {
		namedEnvironments, err := service.GetByName(environment.Name)
		require.NoError(t, err)
		require.NotNil(t, namedEnvironments)
	}
}

func TestEnvironmentServiceGetByPartialName(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	environments, err := service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, environments)
	require.Len(t, environments, 0)

	environments, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, environments)
	require.Len(t, environments, 0)

	environments, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, environments)

	for _, environment := range environments {
		namedEnvironments, err := service.GetByPartialName(environment.Name)
		require.NoError(t, err)
		require.NotNil(t, namedEnvironments)
	}
}

func TestEnvironmentServiceNew(t *testing.T) {
	serviceFunction := newEnvironmentService
	client := &sling.Sling{}
	uriTemplate := emptyString
	sortOrderPath := emptyString
	summaryPath := emptyString
	serviceName := serviceEnvironmentService

	testCases := []struct {
		name          string
		f             func(*sling.Sling, string, string, string) *environmentService
		client        *sling.Sling
		uriTemplate   string
		sortOrderPath string
		summaryPath   string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, sortOrderPath, summaryPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, sortOrderPath, summaryPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, sortOrderPath, summaryPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.sortOrderPath, tc.summaryPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestEnvironmentServiceUpdate(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	environment := CreateTestEnvironment(t, service)
	err := DeleteTestEnvironment(t, service, environment)
	require.NoError(t, err)

	newAllowDynamicInfrastructure := createRandomBoolean()
	newDescription := getRandomName()
	newName := getRandomName()
	newSortOrder := environment.SortOrder + 1
	newUseGuidedFailure := createRandomBoolean()

	environment.AllowDynamicInfrastructure = newAllowDynamicInfrastructure
	environment.Description = newDescription
	environment.Name = newName
	environment.SortOrder = newSortOrder
	environment.UseGuidedFailure = newUseGuidedFailure

	updatedEnvironment := UpdateEnvironment(t, service, environment)
	require.NotNil(t, updatedEnvironment)

	require.NotEmpty(t, updatedEnvironment.GetID())
	require.Equal(t, updatedEnvironment.ID, updatedEnvironment.GetID())
	require.Equal(t, newAllowDynamicInfrastructure, updatedEnvironment.AllowDynamicInfrastructure)
	require.Equal(t, newDescription, updatedEnvironment.Description)
	require.Equal(t, newName, updatedEnvironment.Name)
	require.Equal(t, newSortOrder, updatedEnvironment.SortOrder)
	require.Equal(t, newUseGuidedFailure, updatedEnvironment.UseGuidedFailure)
}
