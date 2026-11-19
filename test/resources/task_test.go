package resources

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tasks"
)

func TestNewTask(t *testing.T) {
	task := tasks.NewTask()
	require.NotNil(t, task)
	require.NotNil(t, task.Arguments)
}
