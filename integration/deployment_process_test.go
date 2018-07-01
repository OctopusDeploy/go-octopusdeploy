package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	client = initTest()
}

func TestDeploymentProcessGet(t *testing.T) {
	project := createTestProject(t, getRandomProjectName())
	defer cleanProject(t, project.ID)

	deploymentProcess, err := client.DeploymentProcess.Get(project.DeploymentProcessID)

	assert.Equal(t, project.DeploymentProcessID, deploymentProcess.ID)
	assert.NoError(t, err, "there should be error raised getting a projects deployment process")
}

func TestDeploymentProcessGetAll(t *testing.T) {
	project := createTestProject(t, getRandomProjectName())
	defer cleanProject(t, project.ID)

	allDeploymentProcess, err := client.DeploymentProcess.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all deployment processes failed when it shouldn't: %s", err)
	}

	numberOfDeploymentProcesses := len(*allDeploymentProcess)

	additionalProject := createTestProject(t, getRandomProjectName())
	defer cleanProject(t, additionalProject.ID)

	allDeploymentProcessAfterCreatingAdditional, err := client.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all deployment processes failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for deployment processes when not expected")
	assert.Equal(t, len(*allDeploymentProcessAfterCreatingAdditional), numberOfDeploymentProcesses+1, "created an additional project and expected number of deployment processes to increase by 1")
}
