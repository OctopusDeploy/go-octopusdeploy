package resources

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/stretchr/testify/assert"
)

func TestVariableScope(t *testing.T) {
	variableScope := variables.VariableScope{}
	assert.Nil(t, variableScope.Actions)
	assert.Len(t, variableScope.Actions, 0)
}

func TestVariableScopeIsEmpty(t *testing.T) {
	variableScope := variables.VariableScope{}
	assert.True(t, variableScope.IsEmpty())

	variableScope.Actions = nil
	assert.True(t, variableScope.IsEmpty())

	variableScope.Actions = []string{}
	assert.True(t, variableScope.IsEmpty())

	variableScope.Actions = []string{"foo"}
	assert.False(t, variableScope.IsEmpty())

	variableScope.Actions = []string{}
	assert.True(t, variableScope.IsEmpty())
}
