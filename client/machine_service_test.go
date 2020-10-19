package client

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMachineService(t *testing.T) *machineService {
	service := newMachineService(nil, TestURIMachines, TestURIDiscoverMachine, TestURIMachineOperatingSystems, TestURIMachineShells)
	testNewService(t, service, TestURIMachines, serviceMachineService)
	return service
}

func CreateTestDeploymentTarget(t *testing.T, service *machineService, name string, environment *model.Environment) *model.DeploymentTarget {
	if service == nil {
		service = createMachineService(t)
	}
	require.NotNil(t, service)

	// thumbprints must be unique, therefore accept a testName string so we can
	// pass through a fixed ID with the name machine that will be consistent
	// through the same test, but different for different tests
	h := md5.New()

	_, err := io.WriteString(h, name)
	require.NoError(t, err)

	_, err = io.WriteString(h, environment.ID)
	require.NoError(t, err)

	thumbprint := fmt.Sprintf("%x", h.Sum(nil))
	environmentIDs := []string{environment.ID}
	roles := []string{"Prod"}

	endpoint := model.NewOfflineDropEndpoint()
	require.NotNil(t, endpoint)

	endpoint.ApplicationsDirectory = "C:\\Applications"
	endpoint.OctopusWorkingDirectory = "C:\\Octopus"

	deploymentTarget := model.NewDeploymentTarget(name, endpoint, environmentIDs, roles)
	deploymentTarget.IsDisabled = true
	deploymentTarget.MachinePolicyID = "MachinePolicies-1"
	deploymentTarget.Status = "Disabled"
	deploymentTarget.Thumbprint = strings.ToUpper(thumbprint[:16])
	deploymentTarget.URI = "https://example.com/"

	resource, err := service.Add(deploymentTarget)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func DeleteTestDeploymentTarget(t *testing.T, service *machineService, deploymentTarget *model.DeploymentTarget) error {
	require.NotNil(t, deploymentTarget)

	if service == nil {
		service = createMachineService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(deploymentTarget.GetID())
}

func IsEqualDeploymentTargets(t *testing.T, expected *model.DeploymentTarget, actual *model.DeploymentTarget) {
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

func TestMachineServiceAdd(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(t, resource)
}

func TestMachineServiceAddGetDelete(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	environmentService := createEnvironmentService(t)
	require.NotNil(t, environmentService)

	environment := CreateTestEnvironment(t, nil)
	expected := CreateTestDeploymentTarget(t, service, getRandomName(), environment)

	actual, err := service.GetByID(expected.ID)
	require.NoError(t, err)
	IsEqualDeploymentTargets(t, expected, actual)

	err = DeleteTestDeploymentTarget(t, service, expected)
	require.NoError(t, err)

	err = DeleteTestEnvironment(t, environmentService, environment)
	require.NoError(t, err)
}

func TestMachineServiceDelete(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	assert.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	assert.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	assert.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
}

func TestMachineServiceDeleteAll(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		defer DeleteTestDeploymentTarget(t, service, resource)
	}
}

func TestMachineServiceGetAll(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	environmentService := createEnvironmentService(t)
	require.NotNil(t, environmentService)

	environment := CreateTestEnvironment(t, nil)
	defer DeleteTestEnvironment(t, environmentService, environment)

	const count int = 32
	expected := map[string]*model.DeploymentTarget{}
	for i := 0; i < count; i++ {
		deploymentTarget := CreateTestDeploymentTarget(t, service, getRandomName(), environment)
		defer DeleteTestDeploymentTarget(t, service, deploymentTarget)
		expected[deploymentTarget.ID] = deploymentTarget
	}

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)
	require.GreaterOrEqual(t, len(resources), count)

	for _, actual := range resources {
		_, ok := expected[actual.ID]
		if ok {
			IsEqualDeploymentTargets(t, expected[actual.ID], actual)
		}
	}
}

func TestMachineServiceGetByID(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	require.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	require.Nil(t, resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	require.Nil(t, resource)

	environmentService := createEnvironmentService(t)
	require.NotNil(t, environmentService)

	environment := CreateTestEnvironment(t, nil)
	defer DeleteTestEnvironment(t, environmentService, environment)

	expected := CreateTestDeploymentTarget(t, service, getRandomName(), environment)
	defer DeleteTestDeploymentTarget(t, service, expected)

	machines, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, machines)

	for _, machine := range machines {
		machineToCompare, err := service.GetByID(machine.ID)
		assert.NoError(t, err)
		IsEqualDeploymentTargets(t, machine, machineToCompare)
	}
}

func TestMachineServiceNew(t *testing.T) {
	serviceFunction := newMachineService
	client := &sling.Sling{}
	uriTemplate := emptyString
	discoverMachinePath := emptyString
	operatingSystemsPath := emptyString
	shellsPath := emptyString
	serviceName := serviceMachineService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *machineService
		client               *sling.Sling
		uriTemplate          string
		discoverMachinePath  string
		operatingSystemsPath string
		shellsPath           string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, discoverMachinePath, operatingSystemsPath, shellsPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, discoverMachinePath, operatingSystemsPath, shellsPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, discoverMachinePath, operatingSystemsPath, shellsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.discoverMachinePath, tc.operatingSystemsPath, tc.shellsPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestMachineServiceUpdate(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	environmentService := createEnvironmentService(t)
	require.NotNil(t, environmentService)

	environment := CreateTestEnvironment(t, nil)
	defer DeleteTestEnvironment(t, environmentService, environment)

	expected := CreateTestDeploymentTarget(t, service, getRandomName(), environment)
	defer DeleteTestDeploymentTarget(t, service, expected)

	expected.Name = getRandomName()

	endpoint, ok := expected.Endpoint.(model.OfflineDropEndpoint)
	require.True(t, ok)

	endpoint.ApplicationsDirectory = getRandomName()
	endpoint.OctopusWorkingDirectory = getRandomName()
	expected.Endpoint = endpoint

	actual, err := service.Update(*expected)
	require.NoError(t, err)
	require.NotNil(t, actual)

	IsEqualDeploymentTargets(t, expected, actual)
}
