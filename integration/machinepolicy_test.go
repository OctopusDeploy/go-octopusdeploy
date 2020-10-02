package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMachinePolicyGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	resource, err := octopusClient.MachinePolicies.GetByID(id)
	require.Equal(t, createResourceNotFoundError("machine policy", "ID", id), err)
	require.Nil(t, resource)
}

func TestMachinePolicyGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	allMachinePolicies, err := octopusClient.MachinePolicies.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all machinepolicies failed when it shouldn't: %s", err)
	}
	numberOfMachinePolicies := len(allMachinePolicies)

	assert.NoError(t, err, "error when looking for machine policies when not expected")
	assert.NotNil(t, allMachinePolicies, "machine policy object returned as nil")
	assert.Equal(t, numberOfMachinePolicies > 0, true, "expecting at least one machine policy to be found")
}
