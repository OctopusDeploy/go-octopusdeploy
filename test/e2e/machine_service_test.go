package e2e

import (
	"crypto/md5"
	"io"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestDeploymentTarget(t *testing.T, client *client.Client, environment *environments.Environment) *machines.DeploymentTarget {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	// thumbprints must be unique, therefore accept a testName string so we can
	// pass through a fixed ID with the name machine that will be consistent
	// through the same test, but different for different tests
	h := md5.New()

	_, err := io.WriteString(h, name)
	require.NoError(t, err)

	_, err = io.WriteString(h, environment.GetID())
	require.NoError(t, err)

	thumbprint := internal.GetRandomThumbprint()
	environmentIDs := []string{environment.GetID()}
	roles := []string{"Prod"}

	endpoint := machines.NewOfflinePackageDropEndpoint()
	require.NotNil(t, endpoint)

	endpoint.ApplicationsDirectory = "C:\\Applications"
	endpoint.WorkingDirectory = "C:\\Octopus"

	deploymentTarget := machines.NewDeploymentTarget(name, endpoint, environmentIDs, roles)
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

func DeleteTestDeploymentTarget(t *testing.T, client *client.Client, deploymentTarget *machines.DeploymentTarget) {
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

func IsEqualDeploymentTargets(t *testing.T, expected *machines.DeploymentTarget, actual *machines.DeploymentTarget) {
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
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

	// machine fields
	assert.Equal(t, expected.Endpoint, actual.Endpoint)
	// asserting on actual.HasLatestCalamari is unreliable in an e2e test because it changes randomly depending on the octopus server docker container
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

	id := internal.GetRandomName()
	err := client.Machines.DeleteByID(id)
	require.Error(t, err)

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
	expected := map[string]*machines.DeploymentTarget{}
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

func TestMachineServiceGetQuery(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := machines.MachinesQuery{
		CommunicationStyles: []string{"Kubernetes"},
	}

	resources, err := client.Machines.Get(query)
	require.NoError(t, err)
	require.NotNil(t, resources)
}

func TestMachineServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	deploymentTarget, err := client.Machines.GetByID(id)
	require.Error(t, err)
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

	deploymentTarget.Name = internal.GetRandomName()

	endpoint, ok := deploymentTarget.Endpoint.(*machines.OfflinePackageDropEndpoint)
	require.True(t, ok)

	endpoint.ApplicationsDirectory = internal.GetRandomName()
	endpoint.WorkingDirectory = internal.GetRandomName()
	deploymentTarget.Endpoint = endpoint

	updatedDeploymentTarget, err := client.Machines.Update(deploymentTarget)
	require.NoError(t, err)
	require.NotNil(t, updatedDeploymentTarget)
	IsEqualDeploymentTargets(t, deploymentTarget, updatedDeploymentTarget)
}

// === NEW ===

func TestMachineServiceAddGetDelete_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	err := machines.DeleteByID(client, client.GetSpaceID(), id)
	require.Error(t, err)

	environment := CreateTestEnvironment_NewClient(t, client)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment_NewClient(t, client, environment)

	createdDeploymentTarget := CreateTestDeploymentTarget_NewClient(t, client, environment)
	require.NotNil(t, createdDeploymentTarget)
	defer DeleteTestDeploymentTarget_NewClient(t, client, createdDeploymentTarget)

	deploymentTarget, err := machines.GetByID(client, createdDeploymentTarget.SpaceID, createdDeploymentTarget.GetID())
	require.NoError(t, err)
	require.NotNil(t, deploymentTarget)
	IsEqualDeploymentTargets(t, createdDeploymentTarget, deploymentTarget)
}

func CreateTestDeploymentTarget_NewClient(t *testing.T, client *client.Client, environment *environments.Environment) *machines.DeploymentTarget {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	// thumbprints must be unique, therefore accept a testName string so we can
	// pass through a fixed ID with the name machine that will be consistent
	// through the same test, but different for different tests
	h := md5.New()

	_, err := io.WriteString(h, name)
	require.NoError(t, err)

	_, err = io.WriteString(h, environment.GetID())
	require.NoError(t, err)

	thumbprint := internal.GetRandomThumbprint()
	environmentIDs := []string{environment.GetID()}
	roles := []string{"Prod"}

	endpoint := machines.NewOfflinePackageDropEndpoint()
	require.NotNil(t, endpoint)

	endpoint.ApplicationsDirectory = "C:\\Applications"
	endpoint.WorkingDirectory = "C:\\Octopus"

	deploymentTarget := machines.NewDeploymentTarget(name, endpoint, environmentIDs, roles)
	deploymentTarget.IsDisabled = true
	deploymentTarget.MachinePolicyID = "MachinePolicies-1"
	deploymentTarget.Status = "Disabled"
	deploymentTarget.Thumbprint = strings.ToUpper(thumbprint[:16])
	deploymentTarget.URI = "https://example.com/"

	require.NoError(t, deploymentTarget.Validate())

	createdDeploymentTarget, err := machines.Add(client, deploymentTarget)
	require.NoError(t, err)
	require.NotNil(t, createdDeploymentTarget)
	require.NotEmpty(t, createdDeploymentTarget.GetID())

	return createdDeploymentTarget
}

func DeleteTestDeploymentTarget_NewClient(t *testing.T, client *client.Client, deploymentTarget *machines.DeploymentTarget) {
	require.NotNil(t, deploymentTarget)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := machines.DeleteByID(client, deploymentTarget.SpaceID, deploymentTarget.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedDeploymentTarget, err := machines.GetByID(client, deploymentTarget.SpaceID, deploymentTarget.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedDeploymentTarget)
}
