package deployments_test

import (
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createDeploymentProcessService(t *testing.T) *deployments.DeploymentProcessService {
	service := deployments.NewDeploymentProcessService(nil, constants.TestURIDeploymentProcesses)
	services.NewServiceTests(t, service, constants.TestURIDeploymentProcesses, constants.ServiceDeploymentProcessesService)
	return service
}

func TestNewDeploymentProcessService(t *testing.T) {
	ServiceFunction := deployments.NewDeploymentProcessService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceDeploymentProcessesService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *deployments.DeploymentProcessService
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

func TestDeploymentProcessServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", ""},
		{"Whitespace", " "},
		{"InvalidID", internal.GetRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createDeploymentProcessService(t)
			require.NotNil(t, service)

			if internal.IsEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
				require.Nil(t, resource)
			} else {
				resource, err := service.GetByID(tc.parameter)
				require.Error(t, err)
				require.Nil(t, resource)
			}
		})
	}
}

func TestDeploymentProcessServiceGetWithEmptyID(t *testing.T) {
	service := deployments.NewDeploymentProcessService(&sling.Sling{}, "")

	resource, err := service.GetByID("")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	require.Nil(t, resource)
}

func TestDeploymentProcessServiceGetWithMissingLinks(t *testing.T) {
	service := createDeploymentProcessService(t)

	repositoryURL, err := url.Parse("https://github.com/OctopusDeploy/manifests.git")
	require.NoError(t, err)

	project := projects.NewProject(internal.GetRandomName(), "Lifecycles-1", "ProjectGroups-1")
	project.PersistenceSettings = projects.NewGitPersistenceSettings("", nil, "main", nil, repositoryURL)
	project.Links = map[string]string{}

	resource, err := service.Get(project, "main")
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = deployments.GetDeploymentProcessByGitRef(nil, "Spaces-1", project, "main")
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestDeploymentProcessServiceGetTemplateWithMissingLinks(t *testing.T) {
	service := createDeploymentProcessService(t)

	deploymentProcess := deployments.NewDeploymentProcess("Projects-1")
	deploymentProcess.Links = map[string]string{}

	resource, err := service.GetTemplate(deploymentProcess, "Channels-1", "")
	require.Error(t, err)
	require.Nil(t, resource)

	template, err := deployments.GetDeploymentProcessTemplate(nil, deploymentProcess, "Channels-1", "")
	require.Error(t, err)
	require.Nil(t, template)
}
