package services

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createScriptModuleService(t *testing.T) *scriptModuleService {
	service := newScriptModuleService(nil, TestURILibraryVariables)
	testNewService(t, service, TestURILibraryVariables, ServiceLibraryVariableSetService)
	return service
}

func CreateTestScriptModule(t *testing.T, service *scriptModuleService) *ScriptModule {
	if service == nil {
		service = createScriptModuleService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()
	description := getRandomName()
	scriptBody := "function Say-Hello()\r\n{\r\n    Write-Output \"Hello, Octopus!\"\r\n}\r\n"
	syntax := "PowerShell"

	scriptModule := NewScriptModule(name)
	scriptModule.Description = description
	scriptModule.ScriptBody = scriptBody
	scriptModule.Syntax = syntax
	require.NoError(t, scriptModule.Validate())

	createdScriptModule, err := service.Add(scriptModule)
	require.NoError(t, err)
	require.NotNil(t, createdScriptModule)
	require.NotEmpty(t, createdScriptModule.GetID())
	require.Equal(t, description, createdScriptModule.Description)
	require.Equal(t, name, createdScriptModule.Name)
	require.Equal(t, scriptBody, createdScriptModule.ScriptBody)
	require.Equal(t, syntax, createdScriptModule.Syntax)

	return createdScriptModule
}

func DeleteTestScriptModule(t *testing.T, service *scriptModuleService, libraryVariableSet *ScriptModule) error {
	require.NotNil(t, libraryVariableSet)

	if service == nil {
		service = createScriptModuleService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(libraryVariableSet.GetID())
}

func IsEqualScriptModules(t *testing.T, expected *ScriptModule, actual *ScriptModule) {
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

	// script module
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ScriptBody, actual.ScriptBody)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.Syntax, actual.Syntax)
	assert.Equal(t, expected.VariableSetID, actual.VariableSetID)
}

func UpdateScriptModule(t *testing.T, service *scriptModuleService, libraryVariableSet *ScriptModule) *ScriptModule {
	if service == nil {
		service = createScriptModuleService(t)
	}
	require.NotNil(t, service)

	updatedScriptModule, err := service.Update(libraryVariableSet)
	assert.NoError(t, err)
	require.NotNil(t, updatedScriptModule)

	return updatedScriptModule
}

func TestScriptModuleServiceSetAddGetDelete(t *testing.T) {
	service := createScriptModuleService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&ScriptModule{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestScriptModule(t, service)
	require.NotNil(t, resource)

	resourceToCompare, err := service.GetByID(resource.GetID())
	require.NoError(t, err)
	require.NotNil(t, resourceToCompare)
	IsEqualScriptModules(t, resource, resourceToCompare)

	err = DeleteTestScriptModule(t, service, resource)
	require.NoError(t, err)
}

func TestScriptModuleServiceAdd(t *testing.T) {
	service := createScriptModuleService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&ScriptModule{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestScriptModule(t, service)
	err = DeleteTestScriptModule(t, service, resource)
	require.NoError(t, err)
}

func TestScriptModuleServiceDeleteAll(t *testing.T) {
	service := createScriptModuleService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources.Items {
		err = DeleteTestScriptModule(t, service, resource)
		assert.NoError(t, err)
	}
}

func TestScriptModuleServiceGetAll(t *testing.T) {
	service := createScriptModuleService(t)
	require.NotNil(t, service)

	libraryVariableSets := []ScriptModule{}

	// create 30 test library variable sets (to be deleted)
	for i := 0; i < 30; i++ {
		libraryVariableSet := CreateTestScriptModule(t, service)
		require.NotNil(t, libraryVariableSet)
		libraryVariableSets = append(libraryVariableSets, *libraryVariableSet)
	}

	allScriptModules, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allScriptModules)
	require.True(t, len(allScriptModules.Items) >= 30)

	for _, libraryVariableSet := range libraryVariableSets {
		require.NotNil(t, libraryVariableSet)
		require.NotEmpty(t, libraryVariableSet.GetID())
		err = DeleteTestScriptModule(t, service, &libraryVariableSet)
		require.NoError(t, err)
	}
}

func TestScriptModuleServiceGetByID(t *testing.T) {
	service := createScriptModuleService(t)
	require.NotNil(t, service)

	id := getRandomName()
	resource, err := service.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources.Items {
		resourceToCompare, err := service.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualScriptModules(t, resource, resourceToCompare)
	}
}

func TestScriptModuleServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
		{"InvalidID", getRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createScriptModuleService(t)
			require.NotNil(t, service)

			if IsEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
				require.Nil(t, resource)

				resourceList, err := service.GetByPartialName(tc.parameter)
				require.Equal(t, createInvalidParameterError(OperationGetByPartialName, ParameterName), err)
				require.NotNil(t, resourceList)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, createInvalidParameterError(OperationDeleteByID, ParameterID))
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

func TestScriptModuleServiceNew(t *testing.T) {
	ServiceFunction := newScriptModuleService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceLibraryVariableSetService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *scriptModuleService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestScriptModuleGetByPartialName(t *testing.T) {
	service := createScriptModuleService(t)
	require.NotNil(t, service)

	resources, err := service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByPartialName, ParameterName))
	require.NotNil(t, resources)
	require.Len(t, resources, 0)

	resources, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByPartialName, ParameterName))
	require.NotNil(t, resources)
	require.Len(t, resources, 0)

	allScriptModules, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allScriptModules)

	for _, resource := range allScriptModules.Items {
		namedResources, err := service.GetByPartialName(resource.Name)
		require.NoError(t, err)
		require.NotNil(t, namedResources)
	}
}

func TestScriptModuleServiceUpdate(t *testing.T) {
	service := createScriptModuleService(t)
	require.NotNil(t, service)

	resource, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = service.Update(&ScriptModule{})
	require.Error(t, err)
	require.Nil(t, resource)

	expected := CreateTestScriptModule(t, service)
	defer DeleteTestScriptModule(t, service, expected)

	expected.ScriptBody = "function Say-Hello()\r\n{\r\n    Write-Output \"Hello, World!\"\r\n}\r\n"
	actual := UpdateScriptModule(t, service, expected)
	IsEqualScriptModules(t, expected, actual)

	expected.Syntax = "Bash"
	actual = UpdateScriptModule(t, service, expected)
	IsEqualScriptModules(t, expected, actual)
}
