package integration

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
)

func init() {
	client = initTest()
}

func TestMachineAddAndDelete(t *testing.T) {
	testName := "TestMachineAddAndDelete"
	testEnvironment := createTestEnvironment(t, testName)
	defer cleanEnvironment(t, testEnvironment.ID)

	machineName := strings.Split(getRandomName(), " ")[1]
	expected := getTestMachine(testEnvironment.ID, machineName)
	actual := createTestMachine(t, testEnvironment.ID, machineName)
	defer cleanMachine(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "machine name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "machine doesn't contain an ID from the octopus server")
}

func TestMachineAddGetAndDelete(t *testing.T) {
	testName := "TestMachineAddGetAndDelete"
	testEnvironment := createTestEnvironment(t, testName)
	defer cleanEnvironment(t, testEnvironment.ID)

	machineName := strings.Split(getRandomName(), " ")[1]
	machine := createTestMachine(t, testEnvironment.ID, machineName)
	defer cleanMachine(t, machine.ID)

	getMachine, err := client.Machine.Get(machine.ID)
	assert.Nil(t, err, "there was an error raised getting machine when there should not be")
	assert.Equal(t, machine.Name, getMachine.Name)
	assert.Equal(t, machine.Thumbprint, getMachine.Thumbprint)
	assert.Equal(t, machine.URI, getMachine.Endpoint.URI)
}

func TestMachineGetThatDoesNotExist(t *testing.T) {
	machineID := "there-is-no-way-this-machine-id-exists-i-hope"
	expected := octopusdeploy.ErrItemNotFound
	machine, err := client.Machine.Get(machineID)

	assert.Error(t, err, "there should have been an error raised as this machine should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, machine, "no machine should have been returned")
}

func TestMachineGetAll(t *testing.T) {
	testName := "TestMachineGetAll"
	testEnvironment := createTestEnvironment(t, testName)
	defer cleanEnvironment(t, testEnvironment.ID)

	// create many machines to test pagination
	machinesToCreate := 32
	sum := 0
	for i := 0; i < machinesToCreate; i++ {
		machineName := strings.Split(getRandomName(), " ")[1]
		machine := createTestMachine(t, testEnvironment.ID, machineName)
		defer cleanMachine(t, machine.ID)
		sum += i
	}

	allMachines, err := client.Machine.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all machines failed when it shouldn't: %s", err)
	}

	numberOfMachines := len(*allMachines)

	// check there are greater than or equal to the amount of machines requested to be created, otherwise pagination isn't working
	if numberOfMachines < machinesToCreate {
		t.Fatalf("There should be at least %d machines created but there was only %d. Pagination is likely not working.", machinesToCreate, numberOfMachines)
	}

	machineName := strings.Split(getRandomName(), " ")[1]
	additionalMachine := createTestMachine(t, testEnvironment.ID, machineName)
	defer cleanMachine(t, additionalMachine.ID)

	allMachinesAfterCreatingAdditional, err := client.Machine.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all machines failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for machine when not expected")
	assert.Equal(t, len(*allMachinesAfterCreatingAdditional), numberOfMachines+1, "created an additional machine and expected number of machines to increase by 1")
}

func TestMachineUpdate(t *testing.T) {
	testName := "TestMachineUpdate"
	testEnvironment := createTestEnvironment(t, testName)
	defer cleanEnvironment(t, testEnvironment.ID)

	machineName := strings.Split(getRandomName(), " ")[1]
	machine := createTestMachine(t, testEnvironment.ID, machineName)
	defer cleanMachine(t, machine.ID)

	const newURI = "https://127.0.0.1/"
	newMachineName := strings.Split(getRandomName(), " ")[1]
	machine.URI = newURI
	machine.Endpoint.URI = newURI
	machine.Name = newMachineName

	updatedMachine, err := client.Machine.Update(&machine)
	assert.Nil(t, err, "error when updating machine")
	assert.Equal(t, newMachineName, updatedMachine.Name, "machine name was not updated")
	assert.Equal(t, newURI, updatedMachine.URI, "machine uri was not updated")
}

func getTestMachine(environmentID, machineName string) octopusdeploy.Machine {
	// Thumbprints have to be unique, so accept a testName string so we can pass through a fixed ID
	// with the name machine that will be consistent through the same test, but different for different
	// tests
	h := md5.New()
	io.WriteString(h, machineName)
	io.WriteString(h, environmentID)
	thumbprint := fmt.Sprintf("%x", h.Sum(nil))

	e := octopusdeploy.Machine{
		EnvironmentIDs:                  []string{environmentID},
		IsDisabled:                      true,
		MachinePolicyID:                 "MachinePolicies-1",
		Name:                            machineName,
		Roles:                           []string{"Prod"},
		Status:                          "Disabled",
		TenantedDeploymentParticipation: octopusdeploy.Untenanted,
		TenantIDs:                       []string{},
		TenantTags:                      []string{},
		Thumbprint:                      strings.ToUpper(thumbprint[:16]),
		URI:                             "https://localhost/",
	}
	return e
}

func createTestMachine(t *testing.T, environmentID, machineName string) octopusdeploy.Machine {
	e := getTestMachine(environmentID, machineName)
	createdMachine, err := client.Machine.Add(&e)

	if err != nil {
		t.Fatalf("creating machine %s failed when it shouldn't: %s", machineName, err)
	}

	return *createdMachine
}

func cleanMachine(t *testing.T, machineID string) {
	err := client.Machine.Delete(machineID)
	if err == nil {
		return
	}

	if err == octopusdeploy.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting machine failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
