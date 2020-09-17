package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/stretchr/testify/assert"
)

func TestMachinePolicyGetThatDoesNotExist(t *testing.T) {
	machinePolicyID := "there-is-no-way-this-machinepolicy-id-exists-i-hope"
	expected := client.ErrItemNotFound
	machinePolicy, err := octopusClient.MachinePolicies.Get(machinePolicyID)
	assert.Error(t, err, "there should have been an error raised as this machinepolicy should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, machinePolicy, "no machinepolicy should have been returned")
}

func TestMachinePolicyGetAll(t *testing.T) {
	allMachinePolicies, err := octopusClient.MachinePolicies.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all machinepolicies failed when it shouldn't: %s", err)
	}
	numberOfMachinePolicies := len(allMachinePolicies)

	assert.Nil(t, err, "error when looking for machine policies when not expected")
	assert.NotNil(t, allMachinePolicies, "machine policy object returned as nil")
	assert.Equal(t, numberOfMachinePolicies > 0, true, "expecting at least one machine policy to be found")
}
