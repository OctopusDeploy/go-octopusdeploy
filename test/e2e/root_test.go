package e2e

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
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

func TestGetSpecificSpaceRoot(t *testing.T) {
	octopusClient := getOctopusClient()
	spaceID := octopusClient.GetSpaceID()
	resource, err := client.GetSpaceRoot(octopusClient, &spaceID)

	assert.NoError(t, err)
	if assert.NotNil(t, resource, "resource should not be nil") {
		assert.NotEmpty(t, resource.Links)
		assert.Empty(t, resource.Version)
		assert.Empty(t, resource.APIVersion)
	}
}

func TestGetServerRoot(t *testing.T) {
	octopusClient := getOctopusClient()
	resource, err := client.GetServerRoot(octopusClient)

	assert.NoError(t, err)
	if assert.NotNil(t, resource, "resource should not be nil") {
		assert.NotEmpty(t, resource.Links)
		assert.NotEmpty(t, resource.Version)
		assert.NotEmpty(t, resource.APIVersion)
	}
}
