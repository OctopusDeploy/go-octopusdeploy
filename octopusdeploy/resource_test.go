package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testExampleWithResource struct {
	Resource
}

func TestResourceEmbedding(t *testing.T) {
	example := &testExampleWithResource{}

	assert.Empty(t, example.GetID())
	assert.Empty(t, example.GetLastModifiedBy())
	assert.Empty(t, example.GetLastModifiedOn())
	assert.Empty(t, example.GetLinks())

	example.SetID("id-value")
	assert.Equal(t, "id-value", example.GetID())
}
