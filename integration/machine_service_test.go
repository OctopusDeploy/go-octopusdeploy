package integration

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestDeploymentTarget(t *testing.T, client *octopusdeploy.Client, environment *octopusdeploy.Environment) *octopusdeploy.DeploymentTarget {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	// thumbprints must be unique, therefore accept a testName string so we can
	// pass through a fixed ID with the name machine that will be consistent
	// through the same test, but different for different tests
	h := md5.New()

	_, err := io.WriteString(h, name)
	require.NoError(t, err)

	_, err = io.WriteString(h, environment.GetID())
	require.NoError(t, err)

	thumbprint := fmt.Sprintf("%x", h.Sum(nil))
	environmentIDs := []string{environment.GetID()}
	roles := []string{"Prod"}

	endpoint := octopusdeploy.NewOfflineDropEndpoint()
	require.NotNil(t, endpoint)

	endpoint.ApplicationsDirectory = "C:\\Applications"
	endpoint.OctopusWorkingDirectory = "C:\\Octopus"

	deploymentTarget := octopusdeploy.NewDeploymentTarget(name, endpoint, environmentIDs, roles)
	deploymentTarget.IsDisabled = true
	deploymentTarget.MachinePolicyID = "MachinePolicies-1"
	deploymentTarget.Status = "Disabled"
	deploymentTarget.Thumbprint = strings.ToUpper(thumbprint[:16])
	deploymentTarget.URI = "https://example.com/"

	require.NoError(t, deploymentTarget.Validate())

	createdDeploymentTarget, err := client.Machines.Add(deploymentTarget)
	require.NoError(t, err)
	require.NotNil(t, createdDeploymentTarget)
	require.NotEmpty(t, createdDeploymentTarget.GetID())

	return createdDeploymentTarget
}

func DeleteTestDeploymentTarget(t *testing.T, client *octopusdeploy.Client, deploymentTarget *octopusdeploy.DeploymentTarget) {
	require.NotNil(t, deploymentTarget)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Machines.DeleteByID(deploymentTarget.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedDeploymentTarget, err := client.Machines.GetByID(deploymentTarget.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedDeploymentTarget)
}

func IsEqualDeploymentTargets(t *testing.T, expected *octopusdeploy.DeploymentTarget, actual *octopusdeploy.DeploymentTarget) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// machine fields
	assert.Equal(t, expected.Endpoint, actual.Endpoint)
	assert.Equal(t, expected.HasLatestCalamari, actual.HasLatestCalamari)
	assert.Equal(t, expected.HealthStatus, actual.HealthStatus)
	assert.Equal(t, expected.IsDisabled, actual.IsDisabled)
	assert.Equal(t, expected.IsInProcess, actual.IsInProcess)
	assert.Equal(t, expected.MachinePolicyID, actual.MachinePolicyID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.OperatingSystem, actual.OperatingSystem)
	assert.Equal(t, expected.ShellName, actual.ShellName)
	assert.Equal(t, expected.ShellVersion, actual.ShellVersion)
	assert.Equal(t, expected.Status, actual.Status)
	assert.Equal(t, expected.StatusSummary, actual.StatusSummary)
	assert.Equal(t, expected.Thumbprint, actual.Thumbprint)
	assert.Equal(t, expected.URI, actual.URI)

	// deployment target fields
	assert.Equal(t, expected.TenantedDeploymentMode, actual.TenantedDeploymentMode)
	assert.Equal(t, expected.EnvironmentIDs, actual.EnvironmentIDs)
	assert.Equal(t, expected.Roles, actual.Roles)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.TenantIDs, actual.TenantIDs)
	assert.Equal(t, expected.TenantTags, actual.TenantTags)
}

func TestMachineServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	err := client.Machines.DeleteByID(id)
	assert.Equal(t, createResourceNotFoundError(octopusdeploy.ServiceMachineService, "ID", id), err)

	environment := CreateTestEnvironment(t, client)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, client, environment)

	createdDeploymentTarget := CreateTestDeploymentTarget(t, client, environment)
	require.NotNil(t, createdDeploymentTarget)
	defer DeleteTestDeploymentTarget(t, client, createdDeploymentTarget)

	deploymentTarget, err := client.Machines.GetByID(createdDeploymentTarget.GetID())
	require.NoError(t, err)
	require.NotNil(t, deploymentTarget)
	IsEqualDeploymentTargets(t, createdDeploymentTarget, deploymentTarget)
}

func TestMachineServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resources, err := client.Machines.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		defer DeleteTestDeploymentTarget(t, client, resource)
	}
}

func TestMachineServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	environment := CreateTestEnvironment(t, client)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, client, environment)

	const count int = 32
	expected := map[string]*octopusdeploy.DeploymentTarget{}
	for i := 0; i < count; i++ {
		deploymentTarget := CreateTestDeploymentTarget(t, client, environment)
		defer DeleteTestDeploymentTarget(t, client, deploymentTarget)
		expected[deploymentTarget.GetID()] = deploymentTarget
	}

	resources, err := client.Machines.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)
	require.GreaterOrEqual(t, len(resources), count)

	for _, actual := range resources {
		_, ok := expected[actual.GetID()]
		if ok {
			IsEqualDeploymentTargets(t, expected[actual.GetID()], actual)
		}
	}
}

func TestMachineServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	deploymentTarget, err := client.Machines.GetByID(id)
	require.Equal(t, createResourceNotFoundError(octopusdeploy.ServiceMachineService, "ID", id), err)
	require.Nil(t, deploymentTarget)

	environment := CreateTestEnvironment(t, client)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, client, environment)

	deploymentTarget = CreateTestDeploymentTarget(t, client, environment)
	require.NotNil(t, deploymentTarget)
	defer DeleteTestDeploymentTarget(t, client, deploymentTarget)

	allDeploymentTargets, err := client.Machines.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allDeploymentTargets)

	for _, deploymentTarget := range allDeploymentTargets {
		deploymentTargetToCompare, err := client.Machines.GetByID(deploymentTarget.GetID())
		assert.NoError(t, err)
		IsEqualDeploymentTargets(t, deploymentTarget, deploymentTargetToCompare)
	}
}

func TestMachineServiceUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	environment := CreateTestEnvironment(t, client)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, client, environment)

	deploymentTarget := CreateTestDeploymentTarget(t, client, environment)
	require.NotNil(t, deploymentTarget)
	defer DeleteTestDeploymentTarget(t, client, deploymentTarget)

	deploymentTarget.Name = getRandomName()

	endpoint, ok := deploymentTarget.Endpoint.(octopusdeploy.OfflineDropEndpoint)
	require.True(t, ok)

	endpoint.ApplicationsDirectory = getRandomName()
	endpoint.OctopusWorkingDirectory = getRandomName()
	deploymentTarget.Endpoint = endpoint

	updatedDeploymentTarget, err := client.Machines.Update(deploymentTarget)
	require.NoError(t, err)
	require.NotNil(t, updatedDeploymentTarget)
	IsEqualDeploymentTargets(t, deploymentTarget, updatedDeploymentTarget)
}
