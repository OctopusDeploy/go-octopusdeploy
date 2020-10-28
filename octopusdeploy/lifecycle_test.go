package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLifecycleNew(t *testing.T) {
	name := "name"

	lifecycle := Lifecycle{}
	require.Error(t, lifecycle.Validate())

	lifecycle = Lifecycle{
		Name: name,
	}
	require.NoError(t, lifecycle.Validate())

	lifecycle = Lifecycle{
		resource: resource{},
	}
	require.Error(t, lifecycle.Validate())

	lifecycle = Lifecycle{
		Name:     name,
		resource: resource{},
	}
	require.NoError(t, lifecycle.Validate())

	lifecycleWithNew := NewLifecycle(emptyString)
	require.Error(t, lifecycleWithNew.Validate())

	lifecycleWithNew = NewLifecycle(whitespaceString)
	require.Error(t, lifecycleWithNew.Validate())

	lifecycleWithNew = NewLifecycle(name)
	require.NoError(t, lifecycleWithNew.Validate())
}
