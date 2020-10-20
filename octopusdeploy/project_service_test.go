package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createProjectService(t *testing.T) *projectService {
	service := newProjectService(nil, TestURIProjects, TestURIProjectPulse, TestURIProjectsExperimentalSummaries)
	testNewService(t, service, TestURIProjects, serviceProjectService)
	return service
}

func CreateTestProject(t *testing.T, service *projectService, lifecycle *Lifecycle, projectGroup *ProjectGroup) *Project {
	if service == nil {
		service = createProjectService(t)
	}
	require.NotNil(t, service)

	if lifecycle == nil {
		lifecycle = CreateTestLifecycle(t, nil)
	}
	require.NotNil(t, lifecycle)

	if projectGroup == nil {
		projectGroup = CreateTestProjectGroup(t, nil)
	}
	require.NotNil(t, projectGroup)

	name := getRandomName()

	project := NewProject(name, lifecycle.GetID(), projectGroup.GetID())
	require.NotNil(t, project)

	createdProject, err := service.Add(project)
	require.NoError(t, err)
	require.NotNil(t, createdProject)
	require.NotEmpty(t, createdProject.GetID())

	return createdProject
}

func DeleteTestProject(t *testing.T, service *projectService, project *Project) error {
	if service == nil {
		service = createProjectService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(project.GetID())
}

func TestNewProjectService(t *testing.T) {
	serviceFunction := newProjectService
	client := &sling.Sling{}
	experimentalSummariesPath := emptyString
	pulsePath := emptyString
	serviceName := serviceProjectService
	uriTemplate := emptyString

	testCases := []struct {
		name                      string
		f                         func(*sling.Sling, string, string, string) *projectService
		client                    *sling.Sling
		uriTemplate               string
		experimentalSummariesPath string
		pulsePath                 string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, pulsePath, experimentalSummariesPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, pulsePath, experimentalSummariesPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, pulsePath, experimentalSummariesPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.pulsePath, tc.experimentalSummariesPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestProjectServiceGetWithEmptyID(t *testing.T) {
	service := newProjectService(nil, emptyString, TestURIProjectPulse, TestURIProjectsExperimentalSummaries)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
