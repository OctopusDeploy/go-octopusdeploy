package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
)

const testProjectName = "Test Project GoLang"

func init() {
	client = initTest()
}

func TestProjectAdd(t *testing.T) {
	p := &octopusdeploy.Project{}
	p.LifecycleID = "Lifecycles-1"
	p.Name = testProjectName
	p.ProjectGroupID = "ProjectGroups-1"

	createdProject, err := client.Projects.Add(p)

	assert.Nil(t, err)
	assert.Equal(t, testProjectName, createdProject.Name)
}

func TestProjectGetByName(t *testing.T) {
	foundProject, err := client.Projects.GetByName(testProjectName)
	assert.Nil(t, err)
	assert.Equal(t, testProjectName, foundProject.Name)
}

func TestProjectGetByNameAndDelete(t *testing.T) {
	foundProject, err := client.Projects.GetByName(testProjectName)
	assert.Nil(t, err, "error when looking for project")

	errDelete := client.Projects.Delete(foundProject.ID)
	assert.Nil(t, errDelete, "error when deleting project")
}
