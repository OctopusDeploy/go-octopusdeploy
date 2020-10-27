package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectTriggerAddGetAndDelete(t *testing.T) {
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

	toCreateTrigger := getTestProjectTrigger(project.GetID())
	toCreateTrigger.Filter.Roles = []string{"MyRole1", "MyRole2"}
	toCreateTrigger.Filter.EventGroups = []string{"Machine"}
	toCreateTrigger.Action.ShouldRedeployWhenMachineHasBeenDeployedTo = true

	projectTrigger := createTestProjectTrigger(t, client, toCreateTrigger)
	defer cleanProjectTrigger(t, client, projectTrigger.GetID())

	getProjectTrigger, err := client.ProjectTriggers.GetByID(projectTrigger.GetID())

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
	require.Equal(t, createResourceNotFoundError(octopusdeploy.ServiceProjectTriggerService, "ID", id), err)
	require.Nil(t, resource)
}

func TestProjectTriggerGetAll(t *testing.T) {
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

	trigger := getTestProjectTrigger(project.GetID())
	createdTrigger := createTestProjectTrigger(t, client, trigger)
	defer cleanProjectTrigger(t, client, createdTrigger.GetID())

	allProjectsTriggers, err := client.ProjectTriggers.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allProjectsTriggers)

	numberOfProjectTriggers := len(allProjectsTriggers)

	additionalTrigger := getTestProjectTrigger(project.GetID())
	additionalTrigger.Name = getRandomName()
	createdAdditionalTrigger := createTestProjectTrigger(t, client, additionalTrigger)
	defer cleanProjectTrigger(t, client, createdAdditionalTrigger.GetID())

	allProjectTriggersAfterCreatingAdditional, err := client.ProjectTriggers.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projectstriggers failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(allProjectTriggersAfterCreatingAdditional), numberOfProjectTriggers+1, "created an additional projecttrigger and expected number of projects to increase by 1")
}

func TestProjectTriggerUpdate(t *testing.T) {
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

	trigger := getTestProjectTrigger(project.GetID())
	createdTrigger := createTestProjectTrigger(t, client, trigger)

	newProjectTriggerName := getRandomName()
	newProjectTriggerRole := []string{"Roley", "Roley 2"}
	const newIsDisabled = true

	createdTrigger.Name = newProjectTriggerName
	createdTrigger.Filter.Roles = newProjectTriggerRole
	createdTrigger.IsDisabled = newIsDisabled

	updatedProjectTrigger, err := client.ProjectTriggers.Update(*createdTrigger)
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
