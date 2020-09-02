package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestProjectAddAndDelete(t *testing.T) {
	projectName := getRandomName()
	expected := getTestProject(projectName)
	actual := createTestProject(t, projectName)

	defer cleanProject(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "project name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "project doesn't contain an ID from the octopus server")
}

func TestProjectAddGetAndDelete(t *testing.T) {
	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	getProject, err := octopusClient.Projects.Get(project.ID)
	assert.Nil(t, err, "there was an error raised getting project when there should not be")
	assert.Equal(t, project.Name, getProject.Name)
}

func TestProjectGetThatDoesNotExist(t *testing.T) {
	projectID := "there-is-no-way-this-project-id-exists-i-hope"
	expected := client.ErrItemNotFound
	project, err := octopusClient.Projects.Get(projectID)

	assert.Error(t, err, "there should have been an error raised as this project should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, project, "no project should have been returned")
}

func TestProjectGetAll(t *testing.T) {
	// create many projects to test pagination
	projectsToCreate := 32
	sum := 0
	for i := 0; i < projectsToCreate; i++ {
		project := createTestProject(t, getRandomName())
		defer cleanProject(t, project.ID)
		sum += i
	}

	allProjects, err := octopusClient.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjects := len(*allProjects)

	// check there are greater than or equal to the amount of projects requested to be created, otherwise pagination isn't working
	if numberOfProjects < projectsToCreate {
		t.Fatalf("There should be at least %d projects created but there was only %d. Pagination is likely not working.", projectsToCreate, numberOfProjects)
	}

	additionalProject := createTestProject(t, getRandomName())
	defer cleanProject(t, additionalProject.ID)

	allProjectsAfterCreatingAdditional, err := octopusClient.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(*allProjectsAfterCreatingAdditional), numberOfProjects+1, "created an additional project and expected number of projects to increase by 1")
}

func TestProjectUpdate(t *testing.T) {
	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	newProjectName := getRandomName()
	const newDescription = "this should be updated"
	const newSkipMachineBehavior = "SkipUnavailableMachines"

	project.Name = newProjectName
	project.Description = newDescription
	project.ProjectConnectivityPolicy.SkipMachineBehavior = newSkipMachineBehavior

	updatedProject, err := octopusClient.Projects.Update(&project)
	assert.Nil(t, err, "error when updating project")
	assert.Equal(t, newProjectName, updatedProject.Name, "project name was not updated")
	assert.Equal(t, newDescription, updatedProject.Description, "project description was not updated")
	assert.Equal(t, newSkipMachineBehavior, project.ProjectConnectivityPolicy.SkipMachineBehavior, "project connectivity policy name was not updated")
}

func TestProjectGetByName(t *testing.T) {
	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	foundProject, err := octopusClient.Projects.GetByName(project.Name)
	assert.Nil(t, err, "error when looking for project when not expected")
	assert.Equal(t, project.Name, foundProject.Name, "project not found when searching by its name")
}

func createTestProject(t *testing.T, projectName string) model.Project {
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

func cleanProject(t *testing.T, projectID string) {
	err := octopusClient.Projects.Delete(projectID)

	if err == nil {
		return
	}

	if err == client.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting project failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
