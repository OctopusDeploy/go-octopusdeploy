package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
)

func init() {
	client = initTest()
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

	getProject, err := client.Projects.Get(project.ID)
	assert.Nil(t, err, "there was an error raised getting project when there should not be")
	assert.Equal(t, project.Name, getProject.Name)
}

func TestProjectGetThatDoesNotExist(t *testing.T) {
	projectID := "there-is-no-way-this-project-id-exists-i-hope"
	expected := octopusdeploy.ErrItemNotFound
	project, err := client.Projects.Get(projectID)

	assert.Error(t, err, "there should have been an error raised as this project should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, project, "no project should have been returned")
}

func TestProjectGetAll(t *testing.T) {
	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	allProjects, err := client.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjects := len(*allProjects)

	additionalProject := createTestProject(t, getRandomName())
	defer cleanProject(t, additionalProject.ID)

	allProjectsAfterCreatingAdditional, err := client.Projects.GetAll()
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

	updatedProject, err := client.Projects.Update(&project)
	assert.Nil(t, err, "error when updating project")
	assert.Equal(t, newProjectName, updatedProject.Name, "project name was not updated")
	assert.Equal(t, newDescription, updatedProject.Description, "project description was not updated")
	assert.Equal(t, newSkipMachineBehavior, project.ProjectConnectivityPolicy.SkipMachineBehavior, "project connectivity policy name was not updated")
}

func TestProjectGetByName(t *testing.T) {
	project := createTestProject(t, getRandomName())
	defer cleanProject(t, project.ID)

	foundProject, err := client.Projects.GetByName(project.Name)
	assert.Nil(t, err, "error when looking for project when not expected")
	assert.Equal(t, project.Name, foundProject.Name, "project not found when searching by its name")
}

func createTestProject(t *testing.T, projectName string) octopusdeploy.Project {
	p := getTestProject(projectName)
	createdProject, err := client.Projects.Add(&p)

	if err != nil {
		t.Fatalf("creating project %s failed when it shouldn't: %s", projectName, err)
	}

	return *createdProject
}

func getTestProject(projectName string) octopusdeploy.Project {
	p := octopusdeploy.NewProject(projectName, "Lifecycles-1", "ProjectGroups-1")

	return *p
}

func cleanProject(t *testing.T, projectID string) {
	err := client.Projects.Delete(projectID)

	if err == nil {
		return
	}

	if err == octopusdeploy.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting project failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
