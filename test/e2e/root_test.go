package e2e

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	octopusClient := getOctopusClient()

	root, err := octopusClient.Root.Get()

	assert.NoError(t, err)
	assert.NotEmpty(t, root)

	if root == nil {
		return
	}

	assert.NoError(t, root.Validate())
	assert.NotEmpty(t, root.Application)
	assert.NotEmpty(t, root.Version)
	assert.NotEmpty(t, root.APIVersion)
	assert.NotEqual(t, root.InstallationID, uuid.Nil)
	assert.NotEmpty(t, root.Links)
}
