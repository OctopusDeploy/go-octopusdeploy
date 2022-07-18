package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestDeploymentProcess(t *testing.T, client *client.Client, project *projects.Project) *deployments.DeploymentProcess {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)
	require.NotNil(t, project)

	name := "Hello world (using PowerShell)"

	action := deployments.NewDeploymentAction(name, "Octopus.Script")
	action.Properties["Octopus.Action.RunOnServer"] = core.NewPropertyValue("true", false)
	action.Properties["Octopus.Action.Script.ScriptBody"] = core.NewPropertyValue("Console.WriteLine('a');", false)

	step := deployments.NewDeploymentStep(name)
	step.Actions = append(step.Actions, action)
	step.TargetRoles = append(step.TargetRoles, "role-1")
	step.Properties["Octopus.Action.TargetRoles"] = core.NewPropertyValue("role-1", false)

	deploymentProcess := deployments.NewDeploymentProcess(project.GetID())
	deploymentProcess.Steps = append(deploymentProcess.Steps, step)

	return deploymentProcess
}

func GetTestDeploymentProcess(t *testing.T, client *client.Client, project *projects.Project) *deployments.DeploymentProcess {
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

func GetByGitRefTestDeploymentProcess(t *testing.T, client *client.Client, project *projects.Project, gitRef string) *deployments.DeploymentProcess {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	deploymentProcess, err := client.DeploymentProcesses.Get(project, gitRef)
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

	newDeploymentProcess := CreateTestDeploymentProcess(t, client, project)
	require.NotNil(t, newDeploymentProcess)

	step := deployments.NewDeploymentStep("foo")
	step.Actions = append(step.Actions, deployments.NewDeploymentAction("foo", "Octopus.Script"))
	step.Actions[0].Properties["Octopus.Action.Script.ScriptBody"] = core.NewPropertyValue("Console.WriteLine('a');", false)
	step.Properties["Octopus.Action.TargetRoles"] = core.NewPropertyValue("role 1", false)

	deploymentProcess.Steps = append(deploymentProcess.Steps, step)
	updatedDeploymentProcess, err := client.DeploymentProcesses.Update(deploymentProcess)
	require.NotNil(t, updatedDeploymentProcess)
	require.NoError(t, err)

	channels, err := client.Projects.GetChannels(project)
	require.NotNil(t, channels)
	require.NoError(t, err)

	template, err := client.DeploymentProcesses.GetTemplate(deploymentProcess, channels[0].GetID(), "")
	require.NoError(t, err)
	require.NotNil(t, template)
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

	deploymentActionWindowsService := deployments.NewDeploymentAction("Install Windows Service", "Octopus.WindowsService")
	deploymentActionWindowsService.Properties["Octopus.Action.EnabledFeatures"] = core.NewPropertyValue("Octopus.Features.WindowsService,Octopus.Features.ConfigurationVariables,Octopus.Features.ConfigurationTransforms,Octopus.Features.SubstituteInFiles", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.AutomaticallyRunConfigurationTransformationFiles"] = core.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.AutomaticallyUpdateAppSettingsAndConnectionStrings"] = core.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.FeedId"] = core.NewPropertyValue("feeds-nugetfeed", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.DownloadOnTentacle"] = core.NewPropertyValue("False", false)
	deploymentActionWindowsService.Properties["Octopus.Action.Package.PackageId"] = core.NewPropertyValue("Newtonsoft.Json", false)
	deploymentActionWindowsService.Properties["Octopus.Action.SubstituteInFiles.Enabled"] = core.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.CreateOrUpdateService"] = core.NewPropertyValue("True", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.DisplayName"] = core.NewPropertyValue("my display name", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.Description"] = core.NewPropertyValue("my desc", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.ExecutablePath"] = core.NewPropertyValue("bin\\Myservice.exe", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.ServiceAccount"] = core.NewPropertyValue("LocalSystem", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.ServiceName"] = core.NewPropertyValue("My service name", false)
	deploymentActionWindowsService.Properties["Octopus.Action.WindowsService.StartMode"] = core.NewPropertyValue("auto", false)
	deploymentActionWindowsService.Properties["Octopus.Action.SubstituteInFiles.TargetFiles"] = core.NewPropertyValue("*.sh", false)
	deploymentActionWindowsService.Properties["test"] = core.NewPropertyValue("foo", true)

	deploymentStep := deployments.NewDeploymentStep(internal.GetRandomName())
	deploymentStep.Actions = append(deploymentStep.Actions, deploymentActionWindowsService)
	deploymentStep.Properties["Octopus.Action.TargetRoles"] = core.NewPropertyValue("octopus-server", false)

	deploymentProcess.Steps = append(deploymentProcess.Steps, deploymentStep)

	updatedDeploymentProcess, err := client.DeploymentProcesses.Update(deploymentProcess)
	require.NoError(t, err)
	require.Equal(t, updatedDeploymentProcess.Steps[0].Properties, deploymentProcess.Steps[0].Properties)
	require.Equal(t, updatedDeploymentProcess.Steps[0].Actions[0].ActionType, deploymentProcess.Steps[0].Actions[0].ActionType)
}
