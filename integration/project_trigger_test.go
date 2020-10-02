package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectTriggers(t *testing.T) {
	t.Run("AddGetAndDelete", TestProjectTriggerAddGetAndDelete)
	t.Run("GetThatDoesNotExist", TestProjectTriggerGetThatDoesNotExist)
	t.Run("GetAll", TestProjectTriggerGetAll)
	t.Run("Update", TestProjectTriggerUpdate)
}

func TestProjectTriggerAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	// need a project to add a trigger to
	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	toCreateTrigger := getTestProjectTrigger(project.ID)
	toCreateTrigger.Filter.Roles = []string{"MyRole1", "MyRole2"}
	toCreateTrigger.Filter.EventGroups = []string{"Machine"}
	toCreateTrigger.Action.ShouldRedeployWhenMachineHasBeenDeployedTo = true

	projectTrigger := createTestProjectTrigger(t, toCreateTrigger)
	defer cleanProjectTrigger(t, projectTrigger.ID)

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

	id := "there-is-no-way-this-projecttrigger-id-exists-i-hope"
	resource, err := octopusClient.ProjectTriggers.GetByID(id)
	assert.Equal(t, createResourceNotFoundError("project trigger", "ID", id), err)
	assert.Nil(t, resource)
}

func TestProjectTriggerGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	trigger := getTestProjectTrigger(project.ID)
	createdTrigger := createTestProjectTrigger(t, trigger)
	defer cleanProjectTrigger(t, createdTrigger.ID)

	allProjectsTriggers, err := octopusClient.ProjectTriggers.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projectstriggers failed when it shouldn't: %s", err)
	}

	numberOfProjectTriggers := len(allProjectsTriggers)

	additionalTrigger := getTestProjectTrigger(project.ID)
	additionalTrigger.Name = getRandomName()
	createdAdditionalTrigger := createTestProjectTrigger(t, additionalTrigger)
	defer cleanProjectTrigger(t, createdAdditionalTrigger.ID)

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

	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	trigger := getTestProjectTrigger(project.ID)
	createdTrigger := createTestProjectTrigger(t, trigger)

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

func createTestProjectTrigger(t *testing.T, trigger *model.ProjectTrigger) *model.ProjectTrigger {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	createdProjectTrigger, err := octopusClient.ProjectTriggers.Add(trigger)

	if err != nil {
		t.Fatalf("creating projecttrigger %s failed when it shouldn't: %s", trigger.Name, err)
	}

	return createdProjectTrigger
}

func getTestProjectTrigger(projectID string) *model.ProjectTrigger {
	return model.NewProjectDeploymentTargetTrigger(getRandomName(), projectID, false, []string{"Role1", "Role2"}, []string{"Machine"}, []string{"MachineCleanupFailed"})
}

func cleanProjectTrigger(t *testing.T, projectTriggerID string) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	err := octopusClient.ProjectTriggers.DeleteByID(projectTriggerID)

	if err == nil {
		return
	}

	if err == client.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting projecttrigger failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
