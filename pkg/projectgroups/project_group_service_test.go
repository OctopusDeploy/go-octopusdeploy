package projectgroups

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createProjectGroupService(t *testing.T) *ProjectGroupService {
	service := NewProjectGroupService(nil, constants.TestURIProjectGroups)
	services.NewServiceTests(t, service, constants.TestURIProjectGroups, constants.ServiceProjectGroupService)
	return service
}

func CreateTestProjectGroup(t *testing.T, service *ProjectGroupService) *ProjectGroup {
	if service == nil {
		service = createProjectGroupService(t)
	}
	require.NotNil(t, service)

	name := internal.GetRandomName()

	projectGroup := NewProjectGroup(name)
	require.NotNil(t, projectGroup)

	createdProjectGroup, err := service.Add(projectGroup)
	require.NoError(t, err)
	require.NotNil(t, createdProjectGroup)
	require.NotEmpty(t, createdProjectGroup.GetID())

	return createdProjectGroup
}

func DeleteTestProjectGroup(t *testing.T, service *ProjectGroupService, projectGroup *ProjectGroup) error {
	require.NotNil(t, projectGroup)

	if service == nil {
		service = createProjectGroupService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(projectGroup.GetID())

}

func TestNewProjectGroupService(t *testing.T) {
	ServiceFunction := NewProjectGroupService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceProjectGroupService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *ProjectGroupService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestProjectGroupServiceGetWithEmptyID(t *testing.T) {
	service := NewProjectGroupService(&sling.Sling{}, "")

	resource, err := service.GetByID("")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	assert.Nil(t, resource)
}
