package integration

import (
	"testing"

	"github.com/MattHodge/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
)

func init() {
	client = initTest()
}

func TestMachinePolicyGetThatDoesNotExist(t *testing.T) {
	machineID := "there-is-no-way-this-environment-id-exists-i-hope"
	expected := octopusdeploy.ErrItemNotFound
	machine, err := client.Environment.Get(machineID)
	assert.Error(t, err, "there should have been an error raised as this environment should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, machine, "no environment should have been returned")
}

func TestMachinePolicyGetAll(t *testing.T) {
	allMachinePolicies, err := client.MachinePolicy.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all environments failed when it shouldn't: %s", err)
	}
	numberOfMachinePolicies := len(*allMachinePolicies)

	assert.Nil(t, err, "error when looking for machine policies when not expected")
	assert.NotNil(t, allMachinePolicies, "machine policy object returned as nil")
	assert.Equal(t, numberOfMachinePolicies > 0, true, "expecting at least one machine policy to be found")
}
