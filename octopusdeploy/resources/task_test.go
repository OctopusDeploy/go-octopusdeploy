package resources

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTask(t *testing.T) {
	task := NewTask()
	require.NotNil(t, task)
	require.NotNil(t, task.Arguments)
}
