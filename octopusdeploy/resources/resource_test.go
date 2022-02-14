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

	example.ID = "id-value"

	assert.Equal(t, "id-value", example.GetID())
}
