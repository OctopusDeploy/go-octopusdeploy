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

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)

	assert.Equal(t, project.DeploymentProcessID, deploymentProcess.GetID())
	assert.NoError(t, err, "there should be error raised getting a projects deployment process")
}

func TestDeploymentProcessGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	allDeploymentProcess, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	numberOfDeploymentProcesses := len(allDeploymentProcess)

	additionalProject := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, additionalProject)
	defer DeleteTestProject(t, client, additionalProject)

	allDeploymentProcessAfterCreatingAdditional, err := client.DeploymentProcesses.GetAll()
	require.NoError(t, err)

	assert.Equal(t, len(allDeploymentProcessAfterCreatingAdditional), numberOfDeploymentProcesses+1, "created an additional project and expected number of deployment processes to increase by 1")
}

func TestDeploymentProcessUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	require.NoError(t, err)
	require.NotNil(t, deploymentProcess)

	deploymentActionWindowService := octopusdeploy.NewDeploymentAction("Install Windows Service", "Octopus.WindowService")
	deploymentActionWindowService.Properties["Octopus.Action.EnabledFeatures"] = octopusdeploy.NewPropertyValue("Octopus.Features.WindowService,Octopus.Features.ConfigurationVariables,Octopus.Features.ConfigurationTransforms,Octopus.Features.SubstituteInFiles", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.AutomaticallyRunConfigurationTransformationFiles"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.AutomaticallyUpdateAppSettingsAndConnectionStrings"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.FeedId"] = octopusdeploy.NewPropertyValue("feeds-nugetfeed", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.DownloadOnTentacle"] = octopusdeploy.NewPropertyValue("False", false)
	deploymentActionWindowService.Properties["Octopus.Action.Package.PackageId"] = octopusdeploy.NewPropertyValue("Newtonsoft.Json", false)
	deploymentActionWindowService.Properties["Octopus.Action.SubstituteInFiles.Enabled"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.CreateOrUpdateService"] = octopusdeploy.NewPropertyValue("True", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.DisplayName"] = octopusdeploy.NewPropertyValue("my display name", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.Description"] = octopusdeploy.NewPropertyValue("my desc", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.ExecutablePath"] = octopusdeploy.NewPropertyValue("bin\\Myservice.exe", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.ServiceAccount"] = octopusdeploy.NewPropertyValue("LocalSystem", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.ServiceName"] = octopusdeploy.NewPropertyValue("My service name", false)
	deploymentActionWindowService.Properties["Octopus.Action.WindowService.StartMode"] = octopusdeploy.NewPropertyValue("auto", false)
	deploymentActionWindowService.Properties["Octopus.Action.SubstituteInFiles.TargetFiles"] = octopusdeploy.NewPropertyValue("*.sh", false)

	step1 := &octopusdeploy.DeploymentStep{
		Name: "My First Step",
		Properties: map[string]string{
			"Octopus.Action.TargetRoles": "octopus-server",
		},
	}

	step1.Actions = append(step1.Actions, *deploymentActionWindowService)

	deploymentProcess.Steps = append(deploymentProcess.Steps, *step1)

	updated, err := client.DeploymentProcesses.Update(deploymentProcess)

	assert.NoError(t, err, "error when updating deployment process")
	assert.Equal(t, updated.Steps[0].Properties, deploymentProcess.Steps[0].Properties)
	assert.Equal(t, updated.Steps[0].Actions[0].ActionType, deploymentProcess.Steps[0].Actions[0].ActionType)
}
