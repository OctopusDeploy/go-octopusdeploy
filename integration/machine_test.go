package integration

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	octopusClient = initTest()
}

func TestMachineAddAndDelete(t *testing.T) {
	testName := "TestMachineAddAndDelete"
	testEnvironment := createTestEnvironment(t, testName)
	defer cleanEnvironment(t, testEnvironment.ID)

	destination := "AWS"
	machineName := strings.Split(getRandomName(), " ")[1]
	expected := getTestMachine(testEnvironment.ID, machineName, destination)
	actual := createTestMachine(t, testEnvironment.ID, machineName, destination)
	defer cleanMachine(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, expected.Destination, "machine name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "machine doesn't contain an ID from the octopus server")
}

func TestMachineAddGetAndDelete(t *testing.T) {
	testName := "TestMachineAddGetAndDelete"
	testEnvironment := createTestEnvironment(t, testName)
	defer cleanEnvironment(t, testEnvironment.ID)

	destination := "AWS"
	machineName := strings.Split(getRandomName(), " ")[1]
	machine := createTestMachine(t, testEnvironment.ID, machineName, destination)
	defer cleanMachine(t, machine.ID)

	getMachine, err := octopusClient.Machines.Get(machine.ID)
	assert.Nil(t, err, "there was an error raised getting machine when there should not be")
	assert.Equal(t, machine.Name, getMachine.Name)
	assert.Equal(t, machine.Thumbprint, getMachine.Thumbprint)
	assert.Equal(t, machine.URI, getMachine.Endpoint.URI)
}

func TestMachineGetThatDoesNotExist(t *testing.T) {
	machineID := "there-is-no-way-this-machine-id-exists-i-hope"
	expected := client.ErrItemNotFound
	machine, err := octopusClient.Machines.Get(machineID)

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
		destination := "AWS"
		machineName := strings.Split(getRandomName(), " ")[1]
		machine := createTestMachine(t, testEnvironment.ID, machineName, destination)
		defer cleanMachine(t, machine.ID)
		sum += i
	}

	allMachines, err := octopusClient.Machines.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all machines failed when it shouldn't: %s", err)
	}

	numberOfMachines := len(*allMachines)

	// check there are greater than or equal to the amount of machines requested to be created, otherwise pagination isn't working
	if numberOfMachines < machinesToCreate {
		t.Fatalf("There should be at least %d machines created but there was only %d. Pagination is likely not working.", machinesToCreate, numberOfMachines)
	}

	destination := "AWS"
	machineName := strings.Split(getRandomName(), " ")[1]
	additionalMachine := createTestMachine(t, testEnvironment.ID, machineName, destination)
	defer cleanMachine(t, additionalMachine.ID)

	allMachinesAfterCreatingAdditional, err := octopusClient.Machines.GetAll()
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

	destination := "AWS"
	machineName := strings.Split(getRandomName(), " ")[1]
	machine := createTestMachine(t, testEnvironment.ID, machineName, destination)
	defer cleanMachine(t, machine.ID)

	newApplicationsDirectory := "C:\\New-Applications-Directory"
	newWorkingDirectory := "C:\\New-WorkingDirectory"

	newMachineName := strings.Split(getRandomName(), " ")[1]
	machine.Name = newMachineName
	machine.Endpoint.ApplicationsDirectory = newApplicationsDirectory
	machine.Endpoint.WorkingDirectory = newWorkingDirectory

	updatedMachine, err := octopusClient.Machines.Update(&machine)
	assert.Nil(t, err, "error when updating machine")
	assert.Equal(t, newMachineName, updatedMachine.Name, "machine name was not updated")
	assert.Equal(t, newApplicationsDirectory, updatedMachine.Endpoint.ApplicationsDirectory, "machine endpoint's applications Directory was not updated")
	assert.Equal(t, newWorkingDirectory, updatedMachine.Endpoint.WorkingDirectory, "machine endpoint's working Directory was not updated")
}

func getTestMachine(environmentID string, machineName string, destination string) model.Machine {
	// Thumbprints have to be unique, so accept a testName string so we can pass through a fixed ID
	// with the name machine that will be consistent through the same test, but different for different
	// tests
	h := md5.New()

	_, err := io.WriteString(h, machineName)
	if err != nil {
		panic(err)
	}

	_, err = io.WriteString(h, environmentID)
	if err != nil {
		panic(err)
	}

	thumbprint := fmt.Sprintf("%x", h.Sum(nil))

	endpoint, err := model.NewMachineEndpoint()
	endpoint.ApplicationsDirectory = "C:\\Applications"
	endpoint.CommunicationStyle = "OfflineDrop"
	endpoint.WorkingDirectory = "C:\\Octopus"

	e := model.Machine{
		DeploymentMode:  "Untenanted",
		EnvironmentIDs:  []string{environmentID},
		Endpoint:        endpoint,
		IsDisabled:      true,
		MachinePolicyID: "MachinePolicies-1",
		Name:            machineName,
		Roles:           []string{"Prod"},
		Status:          "Disabled",
		TenantIDs:       []string{},
		TenantTags:      []string{},
		Thumbprint:      strings.ToUpper(thumbprint[:16]),
		URI:             "https://localhost/",
		Destination:     destination,
	}

	return e
}

func createTestMachine(t *testing.T, environmentID string, machineName string, destination string) model.Machine {
	e := getTestMachine(environmentID, machineName, destination)
	createdMachine, err := octopusClient.Machines.Add(&e)

	if err != nil {
		t.Fatalf("creating machine %s failed when it shouldn't: %s", machineName, err)
	}

	return *createdMachine
}

func cleanMachine(t *testing.T, machineID string) {
	err := octopusClient.Machines.Delete(machineID)
	if err == nil {
		return
	}

	if err == client.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting machine failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
