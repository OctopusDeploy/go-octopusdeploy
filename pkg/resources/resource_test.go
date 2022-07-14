package resources

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
	assert.Empty(t, example.GetModifiedBy())
	assert.Empty(t, example.GetModifiedOn())
	assert.Empty(t, example.GetLinks())

	example.ID = "id-value"

	assert.Equal(t, "id-value", example.GetID())
}
