package integration

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanMachine(t *testing.T, octopusClient *client.Client, machineID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.Machines.DeleteByID(machineID)
	assert.NoError(t, err)
}

func createTestMachine(t *testing.T, octopusClient *client.Client, environmentID string, name string) model.Machine {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	e := createMachine(t, environmentID, name)
	resource, err := octopusClient.Machines.Add(&e)
	require.NoError(t, err)

	return *resource
}

func createMachine(t *testing.T, environmentID string, machineName string) model.Machine {
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

	endpoint := &model.MachineEndpoint{}

	assert.NoError(t, err)
	assert.NotNil(t, endpoint)

	endpoint.ApplicationsDirectory = "C:\\Applications"

	communicationStyle, _ := enum.ParseCommunicationStyle("OfflineDrop")
	endpoint.CommunicationStyle = communicationStyle

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
	}

	return e
}

func isEqualMachines(t *testing.T, expected model.Machine, actual model.Machine) {
	assert := assert.New(t)

	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	assert.Equal(expected.DeploymentMode, actual.DeploymentMode)
	assert.Equal(expected.Endpoint, actual.Endpoint)
	assert.Equal(expected.EnvironmentIDs, actual.EnvironmentIDs)
	assert.Equal(expected.HasLatestCalamari, actual.HasLatestCalamari)
	assert.Equal(expected.HealthStatus, actual.HealthStatus)
	assert.Equal(expected.ID, actual.ID)
	assert.Equal(expected.IsDisabled, actual.IsDisabled)
	assert.Equal(expected.IsInProcess, actual.IsInProcess)
	assert.Equal(expected.Links, actual.Links)
	assert.Equal(expected.MachinePolicyID, actual.MachinePolicyID)
	assert.Equal(expected.Name, actual.Name)
	assert.Equal(expected.OperatingSystem, actual.OperatingSystem)
	assert.Equal(expected.Roles, actual.Roles)
	assert.Equal(expected.ShellName, actual.ShellName)
	assert.Equal(expected.ShellVersion, actual.ShellVersion)
	assert.Equal(expected.Status, actual.Status)
	assert.Equal(expected.StatusSummary, actual.StatusSummary)
	assert.Equal(expected.TenantIDs, actual.TenantIDs)
	assert.Equal(expected.TenantTags, actual.TenantTags)
	assert.Equal(expected.Thumbprint, actual.Thumbprint)
	assert.Equal(expected.TenantTags, actual.TenantTags)
}

func TestMachines(t *testing.T) {
	t.Run("AddGetDelete", TestMachineAddGetDelete)
	t.Run("GetAll", TestMachineGetAll)
	t.Run("Update", TestMachineUpdate)
}

func TestMachineAddGetDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	testEnvironment := createTestEnvironment(t, getRandomName())
	defer cleanEnvironment(t, testEnvironment.ID)

	expected := createTestMachine(t, octopusClient, testEnvironment.ID, getRandomName())
	defer cleanMachine(t, octopusClient, expected.ID)

	actual, err := octopusClient.Machines.GetByID(expected.ID)
	require.NoError(t, err)
	isEqualMachines(t, expected, *actual)
}

func TestMachineGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	testEnvironment := createTestEnvironment(t, getRandomName())
	defer cleanEnvironment(t, testEnvironment.ID)

	const count int = 32
	expected := map[string]model.Machine{}
	for i := 0; i < count; i++ {
		resource := createTestMachine(t, octopusClient, testEnvironment.ID, getRandomName())
		defer cleanMachine(t, octopusClient, resource.ID)
		expected[resource.ID] = resource
	}

	resources, err := octopusClient.Machines.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resources)
	assert.GreaterOrEqual(t, len(resources), count)

	for _, actual := range resources {
		_, ok := expected[actual.ID]
		if ok {
			isEqualMachines(t, expected[actual.ID], actual)
		}
	}
}

func TestMachineUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	testEnvironment := createTestEnvironment(t, getRandomName())
	defer cleanEnvironment(t, testEnvironment.ID)

	expected := createTestMachine(t, octopusClient, testEnvironment.ID, getRandomName())
	defer cleanMachine(t, octopusClient, expected.ID)

	expected.Name = getRandomName()
	expected.Endpoint.ApplicationsDirectory = getRandomName()
	expected.Endpoint.WorkingDirectory = getRandomName()

	actual, err := octopusClient.Machines.Update(expected)
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	isEqualMachines(t, expected, *actual)
}
