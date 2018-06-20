package integration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
)

func init() {
	client = initTest()
}

func TestProjectAdd(t *testing.T) {
	p := &octopusdeploy.Project{}
	p.LifecycleID = "Lifecycles-1"
	p.Name = "Test Project GoLang"
	p.ProjectGroupID = "ProjectGroups-1"

	createdProject, err := client.Projects.Add(p)

	if err != nil {
		fmt.Println(err.Error())
	}

	assert.Nil(t, err)
	assert.Equal(t, "Test Project GoLang", createdProject.Name)
}
