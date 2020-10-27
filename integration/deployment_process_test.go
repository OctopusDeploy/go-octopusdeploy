package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeploymentProcessGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle, err := CreateTestLifecycle(t, client)
	require.NoError(t, err)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup, err := CreateTestProjectGroup(t, client)
	require.NoError(t, err)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project, err := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NoError(t, err)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)

	assert.Equal(t, project.DeploymentProcessID, deploymentProcess.GetID())
	assert.NoError(t, err, "there should be error raised getting a projects deployment process")
}

func TestDeploymentProcessGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle, err := CreateTestLifecycle(t, client)
	require.NoError(t, err)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup, err := CreateTestProjectGroup(t, client)
	require.NoError(t, err)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project, err := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NoError(t, err)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	allDeploymentProcess, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	numberOfDeploymentProcesses := len(allDeploymentProcess)

	additionalProject, err := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NoError(t, err)
	require.NotNil(t, additionalProject)
	defer DeleteTestProject(t, client, additionalProject)

	allDeploymentProcessAfterCreatingAdditional, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	assert.Equal(t, len(allDeploymentProcessAfterCreatingAdditional), numberOfDeploymentProcesses+1, "created an additional project and expected number of deployment processes to increase by 1")
}

func TestDeploymentProcessUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle, err := CreateTestLifecycle(t, client)
	require.NoError(t, err)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup, err := CreateTestProjectGroup(t, client)
	require.NoError(t, err)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project, err := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NoError(t, err)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	require.NoError(t, err)
	require.NotNil(t, deploymentProcess)

	deploymentActionWindowService := &octopusdeploy.DeploymentAction{
		Name:       "Install Windows Service",
		ActionType: "Octopus.WindowService",
		Properties: map[string]string{
			"Octopus.Action.WindowService.CreateOrUpdateService":                        "True",
			"Octopus.Action.WindowService.ServiceAccount":                               "LocalSystem",
			"Octopus.Action.WindowService.StartMode":                                    "auto",
			"Octopus.Action.Package.AutomaticallyRunConfigurationTransformationFiles":   "True",
			"Octopus.Action.Package.AutomaticallyUpdateAppSettingsAndConnectionStrings": "True",
			"Octopus.Action.EnabledFeatures":                                            "Octopus.Features.WindowService,Octopus.Features.ConfigurationVariables,Octopus.Features.ConfigurationTransforms,Octopus.Features.SubstituteInFiles",
			"Octopus.Action.Package.FeedId":                                             "feeds-nugetfeed",
			"Octopus.Action.Package.DownloadOnTentacle":                                 "False",
			"Octopus.Action.Package.PackageId":                                          "Newtonsoft.Json",
			"Octopus.Action.WindowService.ServiceName":                                  "My service name",
			"Octopus.Action.WindowService.DisplayName":                                  "my display name",
			"Octopus.Action.WindowService.Description":                                  "my desc",
			"Octopus.Action.WindowService.ExecutablePath":                               "bin\\Myservice.exe",
			"Octopus.Action.SubstituteInFiles.Enabled":                                  "True",
			"Octopus.Action.SubstituteInFiles.TargetFiles":                              "*.sh",
		},
	}

	step1 := &octopusdeploy.DeploymentStep{
		Name: "My First Step",
		Properties: map[string]string{
			"Octopus.Action.TargetRoles": "octopus-server",
		},
	}

	step1.Actions = append(step1.Actions, *deploymentActionWindowService)

	deploymentProcess.Steps = append(deploymentProcess.Steps, *step1)

	updated, err := client.DeploymentProcesses.Update(*deploymentProcess)

	assert.NoError(t, err, "error when updating deployment process")
	assert.Equal(t, updated.Steps[0].Properties, deploymentProcess.Steps[0].Properties)
	assert.Equal(t, updated.Steps[0].Actions[0].ActionType, deploymentProcess.Steps[0].Actions[0].ActionType)
}
