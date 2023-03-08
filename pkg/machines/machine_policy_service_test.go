package machines

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMachinePolicyService(t *testing.T) *MachinePolicyService {
	service := NewMachinePolicyService(nil, constants.TestURIMachinePolicies, constants.TestURIMachinePolicyTemplate)
	services.NewServiceTests(t, service, constants.TestURIMachinePolicies, constants.ServiceMachinePolicyService)
	return service
}

func CreateTestMachinePolicy(t *testing.T, service *MachinePolicyService) *MachinePolicy {
	if service == nil {
		service = createMachinePolicyService(t)
	}
	require.NotNil(t, service)

	name := internal.GetRandomName()
	connectionConnectTimeout := getRandomDuration(10)
	connectionRetrySleepInterval := getRandomDuration(1)
	connectionRetryTimeLimit := getRandomDuration(10)
	pollingRequestMaximumMessageProcessingTimeout := getRandomDuration(20)
	pollingRequestQueueTimeout := getRandomDuration(10)

	machineCleanupPolicy := NewMachineCleanupPolicy()
	machineCleanupPolicy.DeleteMachinesElapsedTimeSpan = getRandomDuration(1)

	machineHealthCheckPolicy := NewMachineHealthCheckPolicy()
	machineHealthCheckPolicy.HealthCheckInterval = getRandomDuration(1)

	machinePolicy := NewMachinePolicy(name)
	machinePolicy.ConnectionConnectTimeout = connectionConnectTimeout
	machinePolicy.ConnectionRetrySleepInterval = connectionRetrySleepInterval
	machinePolicy.ConnectionRetryTimeLimit = connectionRetryTimeLimit
	machinePolicy.MachineCleanupPolicy = machineCleanupPolicy
	machinePolicy.MachineHealthCheckPolicy = machineHealthCheckPolicy
	machinePolicy.PollingRequestMaximumMessageProcessingTimeout = pollingRequestMaximumMessageProcessingTimeout
	machinePolicy.PollingRequestQueueTimeout = pollingRequestQueueTimeout
	require.NoError(t, machinePolicy.Validate())

	createdMachinePolicy, err := service.Add(machinePolicy)
	require.NoError(t, err)
	require.NotNil(t, createdMachinePolicy)
	require.NotEmpty(t, createdMachinePolicy.GetID())
	require.Equal(t, name, createdMachinePolicy.Name)

	return createdMachinePolicy
}

func DeleteTestMachinePolicy(t *testing.T, service *MachinePolicyService, machinePolicy *MachinePolicy) {
	require.NotNil(t, machinePolicy)

	if service == nil {
		service = createMachinePolicyService(t)
	}
	require.NotNil(t, service)

	err := service.DeleteByID(machinePolicy.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedMachinePolicy, err := service.GetByID(machinePolicy.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedMachinePolicy)
}

func IsEqualMachinePolicies(t *testing.T, expected *MachinePolicy, actual *MachinePolicy) {
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

	// machine policy
	assert.Equal(t, expected.ConnectionConnectTimeout, actual.ConnectionConnectTimeout)
	assert.Equal(t, expected.ConnectionRetryCountLimit, actual.ConnectionRetryCountLimit)
	assert.Equal(t, expected.ConnectionRetrySleepInterval, actual.ConnectionRetrySleepInterval)
	assert.Equal(t, expected.ConnectionRetryTimeLimit, actual.ConnectionRetryTimeLimit)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.IsDefault, actual.IsDefault)
	assert.Equal(t, expected.MachineCleanupPolicy, actual.MachineCleanupPolicy)
	assert.Equal(t, expected.MachineConnectivityPolicy, actual.MachineConnectivityPolicy)
	assert.Equal(t, expected.MachineHealthCheckPolicy, actual.MachineHealthCheckPolicy)
	assert.Equal(t, expected.MachineUpdatePolicy, actual.MachineUpdatePolicy)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.PollingRequestMaximumMessageProcessingTimeout, actual.PollingRequestMaximumMessageProcessingTimeout)
	assert.Equal(t, expected.PollingRequestQueueTimeout, actual.PollingRequestQueueTimeout)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func UpdateMachinePolicy(t *testing.T, service *MachinePolicyService, machinePolicy *MachinePolicy) *MachinePolicy {
	if service == nil {
		service = createMachinePolicyService(t)
	}
	require.NotNil(t, service)

	updatedMachinePolicy, err := service.Update(machinePolicy)
	require.NoError(t, err)
	require.NotNil(t, updatedMachinePolicy)

	return updatedMachinePolicy
}

func TestMachinePolicySetAddGetDelete(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterMachinePolicy))
	require.Nil(t, resource)

	resource, err = service.Add(&MachinePolicy{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestMachinePolicy(t, service)
	require.NotNil(t, resource)
	defer DeleteTestMachinePolicy(t, service, resource)

	resourceToCompare, err := service.GetByID(resource.GetID())
	require.NoError(t, err)
	require.NotNil(t, resourceToCompare)
	IsEqualMachinePolicies(t, resource, resourceToCompare)
}

func TestMachinePolicyServiceAdd(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterMachinePolicy))
	require.Nil(t, resource)

	resource, err = service.Add(&MachinePolicy{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestMachinePolicy(t, service)
	defer DeleteTestMachinePolicy(t, service, resource)
}

func TestMachinePolicyServiceDeleteAll(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	machinePolicies, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, machinePolicies)

	for _, machinePolicy := range machinePolicies {
		if !machinePolicy.IsDefault {
			defer DeleteTestMachinePolicy(t, service, machinePolicy)
		}
	}
}

func TestMachinePolicyServiceGetAll(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	// create 30 test machine policies (to be deleted)
	for i := 0; i < 30; i++ {
		machinePolicy := CreateTestMachinePolicy(t, service)
		require.NotNil(t, machinePolicy)
		defer DeleteTestMachinePolicy(t, service, machinePolicy)
	}

	allMachinePolicies, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allMachinePolicies)
	require.True(t, len(allMachinePolicies) >= 30)
}

func TestMachinePolicyServiceGetByID(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	id := internal.GetRandomName()
	resource, err := service.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := service.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualMachinePolicies(t, resource, resourceToCompare)
	}
}

func TestMachinePolicyServiceGetTemplate(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	machinePolicy, err := service.GetTemplate()
	require.NoError(t, err)
	require.NotNil(t, machinePolicy)
}

func TestMachinePolicyServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", ""},
		{"Whitespace", " "},
		{"InvalidID", internal.GetRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createMachinePolicyService(t)
			require.NotNil(t, service)

			if internal.IsEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
				require.Nil(t, resource)

				resourceList, err := service.GetByPartialName(tc.parameter)
				require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName), err)
				require.NotNil(t, resourceList)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID))
			} else {
				resource, err := service.GetByID(tc.parameter)
				require.Error(t, err)
				require.Nil(t, resource)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
			}
		})
	}
}

func TestMachinePolicyServiceNew(t *testing.T) {
	ServiceFunction := NewMachinePolicyService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceMachinePolicyService

	testCases := []struct {
		name         string
		f            func(*sling.Sling, string, string) *MachinePolicyService
		client       *sling.Sling
		uriTemplate  string
		templatePath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, constants.TestURIMachinePolicyTemplate},
		{"EmptyURITemplate", ServiceFunction, client, "", constants.TestURIMachinePolicyTemplate},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", constants.TestURIMachinePolicyTemplate},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.templatePath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestMachinePolicyGetByPartialName(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	resources, err := service.GetByPartialName("")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName))
	require.NotNil(t, resources)
	require.Len(t, resources, 0)

	resources, err = service.GetByPartialName(" ")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName))
	require.NotNil(t, resources)
	require.Len(t, resources, 0)

	resources, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		namedResources, err := service.GetByPartialName(resource.Name)
		require.NoError(t, err)
		require.NotNil(t, namedResources)
	}
}

func TestMachinePolicyServiceUpdate(t *testing.T) {
	service := createMachinePolicyService(t)
	require.NotNil(t, service)

	resource, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = service.Update(&MachinePolicy{})
	require.Error(t, err)
	require.Nil(t, resource)

	expected := CreateTestMachinePolicy(t, service)
	actual := UpdateMachinePolicy(t, service, expected)
	IsEqualMachinePolicies(t, expected, actual)
	defer DeleteTestMachinePolicy(t, service, expected)
}
