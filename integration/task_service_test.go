package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualTasks(t *testing.T, expected *services.Task, actual services.Task) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// TODO: compare remaining values
}

func TestTaskServiceGetQuery(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := services.TasksQuery{}

	tasks, err := client.Tasks.Get(query)
	require.NoError(t, err)
	require.NotNil(t, tasks)

	for _, task := range tasks.Items {
		query = services.TasksQuery{
			IDs: []string{task.GetID()},
		}

		taskToCompare, err := client.Tasks.Get(query)
		require.NoError(t, err)
		AssertEqualTasks(t, task, *taskToCompare.Items[0])
	}
}
