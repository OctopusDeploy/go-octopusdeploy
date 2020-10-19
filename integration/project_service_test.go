package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectAddAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	projectName := getRandomName()
	expected := getTestProject(projectName)
	actual := createTestProject(t, octopusClient, projectName)

	defer cleanProject(t, octopusClient, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "project name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "project doesn't contain an ID from the octopus server")
}

func TestProjectAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	project := createTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, project.ID)

	getProject, err := octopusClient.Projects.GetByID(project.ID)
	assert.NoError(t, err, "there was an error raised getting project when there should not be")
	assert.Equal(t, project.Name, getProject.Name)
}

func TestProjectGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	resource, err := octopusClient.Projects.GetByID(id)
	require.Equal(t, createResourceNotFoundError("ProjectService", "ID", id), err)
	require.Nil(t, resource)
}

func TestProjectGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	// create many projects to test pagination
	projectsToCreate := 32
	sum := 0
	for i := 0; i < projectsToCreate; i++ {
		project := createTestProject(t, octopusClient, getRandomName())
		defer cleanProject(t, octopusClient, project.ID)
		sum += i
	}

	allProjects, err := octopusClient.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjects := len(allProjects)

	// check there are greater than or equal to the amount of projects requested to be created, otherwise pagination isn't working
	if numberOfProjects < projectsToCreate {
		t.Fatalf("There should be at least %d projects created but there was only %d. Pagination is likely not working.", projectsToCreate, numberOfProjects)
	}

	additionalProject := createTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, additionalProject.ID)

	allProjectsAfterCreatingAdditional, err := octopusClient.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(allProjectsAfterCreatingAdditional), numberOfProjects+1, "created an additional project and expected number of projects to increase by 1")
}

func TestProjectUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	project := createTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, project.ID)

	newProjectName := getRandomName()
	const newDescription = "this should be updated"
	const newSkipMachineBehavior = "SkipUnavailableMachines"

	project.Name = newProjectName
	project.Description = newDescription
	project.ProjectConnectivityPolicy.SkipMachineBehavior = newSkipMachineBehavior

	updatedProject, err := octopusClient.Projects.Update(project)
	require.NoError(t, err)
	require.Equal(t, newProjectName, updatedProject.Name, "project name was not updated")
	require.Equal(t, newDescription, updatedProject.Description, "project description was not updated")
	require.Equal(t, newSkipMachineBehavior, project.ProjectConnectivityPolicy.SkipMachineBehavior, "project connectivity policy name was not updated")
}

func TestProjectGetByName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	project := createTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, project.ID)

	foundProject, err := octopusClient.Projects.GetByName(project.Name)
	require.NoError(t, err, "error when looking for project when not expected")
	require.Equal(t, project.Name, foundProject.Name, "project not found when searching by its name")
}

func createTestProject(t *testing.T, octopusClient *client.Client, projectName string) model.Project {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	p := getTestProject(projectName)
	createdProject, err := octopusClient.Projects.Add(&p)

	if err != nil {
		t.Fatalf("creating project %s failed when it shouldn't: %s", projectName, err)
	}

	return *createdProject
}

func getTestProject(projectName string) model.Project {
	p := model.NewProject(projectName, "Lifecycles-1", "ProjectGroups-1")
	return *p
}

func cleanProject(t *testing.T, octopusClient *client.Client, projectID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.Projects.DeleteByID(projectID)
	assert.NoError(t, err)
}
