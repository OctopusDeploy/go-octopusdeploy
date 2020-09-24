package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func TestDeploymentProcessGet(t *testing.T) {
	octopusClient := getOctopusClient()

	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	deploymentProcess, err := octopusClient.DeploymentProcesses.GetByID(project.DeploymentProcessID)

	assert.Equal(t, project.DeploymentProcessID, deploymentProcess.ID)
	assert.NoError(t, err, "there should be error raised getting a projects deployment process")
}

func TestDeploymentProcessGetAll(t *testing.T) {
	octopusClient := getOctopusClient()

	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	allDeploymentProcess, err := octopusClient.DeploymentProcesses.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all deployment processes failed when it shouldn't: %s", err)
	}

	numberOfDeploymentProcesses := len(allDeploymentProcess)

	additionalProject := createTestProject(t, getRandomName())
	defer cleanProject(t, additionalProject.ID)

	allDeploymentProcessAfterCreatingAdditional, err := octopusClient.DeploymentProcesses.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all deployment processes failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for deployment processes when not expected")
	assert.Equal(t, len(allDeploymentProcessAfterCreatingAdditional), numberOfDeploymentProcesses+1, "created an additional project and expected number of deployment processes to increase by 1")
}

func TestDeploymentProcessUpdate(t *testing.T) {
	octopusClient := getOctopusClient()

	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	deploymentProcess, err := octopusClient.DeploymentProcesses.GetByID(project.DeploymentProcessID)

	if err != nil {
		t.Fatalf("Retrieving deployment processes failed when it shouldn't: %s", err)
	}

	deploymentActionWindowsService := &model.DeploymentAction{
		Name:       "Install Windows Service",
		ActionType: "Octopus.WindowsService",
		Properties: map[string]string{
			"Octopus.Action.WindowsService.CreateOrUpdateService":                       "True",
			"Octopus.Action.WindowsService.ServiceAccount":                              "LocalSystem",
			"Octopus.Action.WindowsService.StartMode":                                   "auto",
			"Octopus.Action.Package.AutomaticallyRunConfigurationTransformationFiles":   "True",
			"Octopus.Action.Package.AutomaticallyUpdateAppSettingsAndConnectionStrings": "True",
			"Octopus.Action.EnabledFeatures":                                            "Octopus.Features.WindowsService,Octopus.Features.ConfigurationVariables,Octopus.Features.ConfigurationTransforms,Octopus.Features.SubstituteInFiles",
			"Octopus.Action.Package.FeedId":                                             "feeds-nugetfeed",
			"Octopus.Action.Package.DownloadOnTentacle":                                 "False",
			"Octopus.Action.Package.PackageId":                                          "Newtonsoft.Json",
			"Octopus.Action.WindowsService.ServiceName":                                 "My service name",
			"Octopus.Action.WindowsService.DisplayName":                                 "my display name",
			"Octopus.Action.WindowsService.Description":                                 "my desc",
			"Octopus.Action.WindowsService.ExecutablePath":                              "bin\\Myservice.exe",
			"Octopus.Action.SubstituteInFiles.Enabled":                                  "True",
			"Octopus.Action.SubstituteInFiles.TargetFiles":                              "*.sh",
		},
	}

	step1 := &model.DeploymentStep{
		Name: "My First Step",
		Properties: map[string]string{
			"Octopus.Action.TargetRoles": "octopus-server",
		},
	}

	step1.Actions = append(step1.Actions, *deploymentActionWindowsService)

	deploymentProcess.Steps = append(deploymentProcess.Steps, *step1)

	updated, err := octopusClient.DeploymentProcesses.Update(*deploymentProcess)

	assert.NoError(t, err, "error when updating deployment process")
	assert.Equal(t, updated.Steps[0].Properties, deploymentProcess.Steps[0].Properties)
	assert.Equal(t, updated.Steps[0].Actions[0].ActionType, deploymentProcess.Steps[0].Actions[0].ActionType)
}
