package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectTriggerAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	// need a project to add a trigger to
	project := createTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, project.ID)

	toCreateTrigger := getTestProjectTrigger(project.ID)
	toCreateTrigger.Filter.Roles = []string{"MyRole1", "MyRole2"}
	toCreateTrigger.Filter.EventGroups = []string{"Machine"}
	toCreateTrigger.Action.ShouldRedeployWhenMachineHasBeenDeployedTo = true

	projectTrigger := createTestProjectTrigger(t, octopusClient, toCreateTrigger)
	defer cleanProjectTrigger(t, octopusClient, projectTrigger.ID)

	getProjectTrigger, err := octopusClient.ProjectTriggers.GetByID(projectTrigger.ID)

	assert.NoError(t, err, "there was an error raised getting projecttrigger when there should not be")
	assert.Equal(t, getProjectTrigger.Name, getProjectTrigger.Name)
	assert.ElementsMatch(t, getProjectTrigger.Filter.Roles, toCreateTrigger.Filter.Roles)
	assert.ElementsMatch(t, getProjectTrigger.Filter.EventGroups, toCreateTrigger.Filter.EventGroups)
	assert.Equal(t, getProjectTrigger.Action.ShouldRedeployWhenMachineHasBeenDeployedTo, toCreateTrigger.Action.ShouldRedeployWhenMachineHasBeenDeployedTo)
}

func TestProjectTriggerGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	resource, err := octopusClient.ProjectTriggers.GetByID(id)
	require.Equal(t, createResourceNotFoundError(serviceProjectTriggerService, "ID", id), err)
	require.Nil(t, resource)
}

func TestProjectTriggerGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	project := createTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, project.ID)

	trigger := getTestProjectTrigger(project.ID)
	createdTrigger := createTestProjectTrigger(t, octopusClient, trigger)
	defer cleanProjectTrigger(t, octopusClient, createdTrigger.ID)

	allProjectsTriggers, err := octopusClient.ProjectTriggers.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projectstriggers failed when it shouldn't: %s", err)
	}

	numberOfProjectTriggers := len(allProjectsTriggers)

	additionalTrigger := getTestProjectTrigger(project.ID)
	additionalTrigger.Name = getRandomName()
	createdAdditionalTrigger := createTestProjectTrigger(t, octopusClient, additionalTrigger)
	defer cleanProjectTrigger(t, octopusClient, createdAdditionalTrigger.ID)

	allProjectTriggersAfterCreatingAdditional, err := octopusClient.ProjectTriggers.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projectstriggers failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(allProjectTriggersAfterCreatingAdditional), numberOfProjectTriggers+1, "created an additional projecttrigger and expected number of projects to increase by 1")
}

func TestProjectTriggerUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	project := createTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, project.ID)

	trigger := getTestProjectTrigger(project.ID)
	createdTrigger := createTestProjectTrigger(t, octopusClient, trigger)

	newProjectTriggerName := getRandomName()
	newProjectTriggerRole := []string{"Roley", "Roley 2"}
	const newIsDisabled = true

	createdTrigger.Name = newProjectTriggerName
	createdTrigger.Filter.Roles = newProjectTriggerRole
	createdTrigger.IsDisabled = newIsDisabled

	updatedProjectTrigger, err := octopusClient.ProjectTriggers.Update(*createdTrigger)
	assert.NoError(t, err, "error when updating projecttrigger")
	assert.Equal(t, newProjectTriggerName, updatedProjectTrigger.Name, "projecttrigger name was not updated")
	assert.Equal(t, newProjectTriggerRole, updatedProjectTrigger.Filter.Roles, "projecttrigger roles was not updated")
	assert.Equal(t, newIsDisabled, updatedProjectTrigger.IsDisabled, "projecttrigger isdisabled setting not updated")
}

func createTestProjectTrigger(t *testing.T, octopusClient *octopusdeploy.Client, trigger *octopusdeploy.ProjectTrigger) *octopusdeploy.ProjectTrigger {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	createdProjectTrigger, err := octopusClient.ProjectTriggers.Add(trigger)

	if err != nil {
		t.Fatalf("creating projecttrigger %s failed when it shouldn't: %s", trigger.Name, err)
	}

	return createdProjectTrigger
}

func getTestProjectTrigger(projectID string) *octopusdeploy.ProjectTrigger {
	return octopusdeploy.NewProjectDeploymentTargetTrigger(getRandomName(), projectID, false, []string{"Role1", "Role2"}, []string{"Machine"}, []string{"MachineCleanupFailed"})
}

func cleanProjectTrigger(t *testing.T, octopusClient *octopusdeploy.Client, projectTriggerID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.ProjectTriggers.DeleteByID(projectTriggerID)
	assert.NoError(t, err)
}
