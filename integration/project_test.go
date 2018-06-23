package integration

import (
	"math/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
)

func init() {
	client = initTest()
}

func TestProjectAddAndDelete(t *testing.T) {
	expected := getTestProject()
	actual := createTestProject(t)

	defer cleanProject(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "project name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "project doesn't contain an ID from the octopus server")
}

func TestProjectGetByName(t *testing.T) {
	project := createTestProjectWithRandomName(t)
	defer cleanProject(t, project.ID)

	foundProject, err := client.Projects.GetByName(project.Name)
	assert.Nil(t, err, "error when looking for project when not expected")
	assert.Equal(t, project.Name, foundProject.Name, "project not found when searching by its name")
}

func TestProjectGetAll(t *testing.T) {
	project := createTestProjectWithRandomName(t)
	defer cleanProject(t, project.ID)

	allProjects, err := client.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjects := len(allProjects)

	additionalProject := createTestProjectWithRandomName(t)
	defer cleanProject(t, additionalProject.ID)

	allProjectsAfterCreatingAdditional, err := client.Projects.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for project when not expected")
	assert.Len(t, allProjectsAfterCreatingAdditional, numberOfProjects + 1, "created an additional project and expected number of projects to increase by 1")
}

func createTestProject(t *testing.T) octopusdeploy.Project {
	p := getTestProject()
	createdProject, err := client.Projects.Add(p)

	if err != nil {
		t.Fatalf("Creating a project failed when it shouldn't: %s", err)
	}

	return createdProject
}

func createTestProjectWithRandomName(t *testing.T) octopusdeploy.Project {
	p := getTestProject()
	p.Name = fmt.Sprintf("go-octopusdeploy rest client testing %f", rand.Float64())
	createdProject, err := client.Projects.Add(p)

	if err != nil {
		t.Fatalf("Creating a project failed when it shouldn't: %s", err)
	}

	return createdProject
}

func getTestProject() *octopusdeploy.Project {
	p := &octopusdeploy.Project{}
	p.LifecycleID = "Lifecycles-1"
	p.Name = "go-octopusdeploy rest client testing"
	p.ProjectGroupID = "ProjectGroups-1"

	return p
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
