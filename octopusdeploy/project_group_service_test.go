package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createProjectGroupService(t *testing.T) *projectGroupService {
	service := newProjectGroupService(nil, TestURIProjectGroups)
	services.testNewService(t, service, TestURIProjectGroups, ServiceProjectGroupService)
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
	ServiceFunction := newProjectGroupService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServiceProjectGroupService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *projectGroupService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestProjectGroupServiceGetWithEmptyID(t *testing.T) {
	service := newProjectGroupService(&sling.Sling{}, services.emptyString)

	resource, err := service.GetByID(services.emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}
