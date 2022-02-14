package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func GetTestDeploymentProcess(t *testing.T, client *octopusdeploy.client, project *services.Project) *services.DeploymentProcess {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	require.NotNil(t, deploymentProcess)
	require.NoError(t, err)
	require.Equal(t, project.DeploymentProcessID, deploymentProcess.GetID())

	return deploymentProcess
}

func TestDeploymentProcessGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess := GetTestDeploymentProcess(t, client, project)
	require.NotNil(t, deploymentProcess)
}

func TestDeploymentProcessGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	deploymentProcesses, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	numberOfDeploymentProcesses := len(deploymentProcesses)

	space := GetDefaultSpace(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	totalDeploymentProcesses, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	assert.Equal(t, len(totalDeploymentProcesses), numberOfDeploymentProcesses+1, "created an additional project and expected number of deployment processes to increase by 1")
}

func TestDeploymentProcessUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess := GetTestDeploymentProcess(t, client, project)
	require.NotNil(t, deploymentProcess)

	deploymentActionWindowsService := services.NewDeploymentAction("Install Windows service", "Octopus.WindowsService")
	deploymentActionWindowsService.Properties["Octopus.Action.EnabledFeatures"] = services.NewPropertyValue("Octopus.Features.WindowsService,Octopus.Features.ConfigurationVariables,Octopus.Features.ConfigurationTransforms,Octopus.Features.SubstituteInFiles", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.AutomaticallyRunConfigurationTransformationFiles"] = services.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.AutomaticallyUpdateAppSettingsAndConnectionStrings"] = services.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.FeedId"] = services.NewPropertyValue("feeds-nugetfeed", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.DownloadOnTentacle"] = services.NewPropertyValue("False", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.PackageId"] = services.NewPropertyValue("Newtonsoft.Json", false)
	deploymentActionWindowsService.Properties["Octopus.Action.SubstituteInFiles.Enabled"] = services.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.CreateOrUpdateService"] = services.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.DisplayName"] = services.NewPropertyValue("my display name", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.Description"] = services.NewPropertyValue("my desc", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.ExecutablePath"] = services.NewPropertyValue("bin\\Myservice.exe", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.ServiceAccount"] = services.NewPropertyValue("LocalSystem", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.ServiceName"] = services.NewPropertyValue("My service name", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.StartMode"] = services.NewPropertyValue("auto", false)
	deploymentActionWindowsService.Properties["Octopus.Action.SubstituteInFiles.TargetFiles"] = services.NewPropertyValue("*.sh", false)
	deploymentActionWindowsService.Properties["test"] = services.NewPropertyValue("foo", true)

	deploymentStep := services.NewDeploymentStep(getRandomName())
	deploymentStep.Actions = append(deploymentStep.Actions, *deploymentActionWindowsService)
	deploymentStep.Properties["Octopus.Action.TargetRoles"] = services.NewPropertyValue("octopus-server", false)

	deploymentProcess.Steps = append(deploymentProcess.Steps, *deploymentStep)

	updatedDeploymentProcess, err := client.DeploymentProcesses.Update(deploymentProcess)
	require.NoError(t, err)
	require.Equal(t, updatedDeploymentProcess.Steps[0].Properties, deploymentProcess.Steps[0].Properties)
	require.Equal(t, updatedDeploymentProcess.Steps[0].Actions[0].ActionType, deploymentProcess.Steps[0].Actions[0].ActionType)
}
