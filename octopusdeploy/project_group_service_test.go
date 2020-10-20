package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createProjectGroupService(t *testing.T) *projectGroupService {
	service := newProjectGroupService(nil, TestURIProjectGroups)
	testNewService(t, service, TestURIProjectGroups, serviceProjectGroupService)
	return service
}

func CreateTestProjectGroup(t *testing.T, service *projectGroupService) *ProjectGroup {
	if service == nil {
		service = createProjectGroupService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()

	projectGroup := NewProjectGroup(name)
	require.NotNil(t, projectGroup)

	createdProjectGroup, err := service.Add(projectGroup)
	require.NoError(t, err)
	require.NotNil(t, createdProjectGroup)
	require.NotEmpty(t, createdProjectGroup.GetID())

	return createdProjectGroup
}

func DeleteTestProjectGroup(t *testing.T, service *projectGroupService, projectGroup *ProjectGroup) error {
	require.NotNil(t, projectGroup)

	if service == nil {
		service = createProjectGroupService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(projectGroup.GetID())

}

func TestNewProjectGroupService(t *testing.T) {
	serviceFunction := newProjectGroupService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceProjectGroupService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *projectGroupService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestProjectGroupServiceGetWithEmptyID(t *testing.T) {
	service := newProjectGroupService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
