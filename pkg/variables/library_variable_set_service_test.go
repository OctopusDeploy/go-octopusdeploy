package variables

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createLibraryVariableSetService(t *testing.T) *LibraryVariableSetService {
	service := NewLibraryVariableSetService(nil, constants.TestURILibraryVariables)
	services.NewServiceTests(t, service, constants.TestURILibraryVariables, constants.ServiceLibraryVariableSetService)
	return service
}

func CreateTestLibraryVariableSet(t *testing.T, service *LibraryVariableSetService) *LibraryVariableSet {
	if service == nil {
		service = createLibraryVariableSetService(t)
	}
	require.NotNil(t, service)

	name := internal.GetRandomName()

	libraryVariableSet := NewLibraryVariableSet(name)
	require.NoError(t, libraryVariableSet.Validate())

	createdLibraryVariableSet, err := service.Add(libraryVariableSet)
	require.NoError(t, err)
	require.NotNil(t, createdLibraryVariableSet)
	require.NotEmpty(t, createdLibraryVariableSet.GetID())
	require.Equal(t, name, createdLibraryVariableSet.Name)

	return createdLibraryVariableSet
}

func DeleteTestLibraryVariableSet(t *testing.T, service *LibraryVariableSetService, libraryVariableSet *LibraryVariableSet) error {
	require.NotNil(t, libraryVariableSet)

	if service == nil {
		service = createLibraryVariableSetService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(libraryVariableSet.GetID())
}

func IsEqualLibraryVariableSets(t *testing.T, expected *LibraryVariableSet, actual *LibraryVariableSet) {
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

	// library variable set
	assert.Equal(t, expected.ContentType, actual.ContentType)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.Templates, actual.Templates)
	assert.Equal(t, expected.VariableSetID, actual.VariableSetID)
}

func UpdateLibraryVariableSet(t *testing.T, service *LibraryVariableSetService, libraryVariableSet *LibraryVariableSet) *LibraryVariableSet {
	if service == nil {
		service = createLibraryVariableSetService(t)
	}
	require.NotNil(t, service)

	updatedLibraryVariableSet, err := service.Update(libraryVariableSet)
	assert.NoError(t, err)
	require.NotNil(t, updatedLibraryVariableSet)

	return updatedLibraryVariableSet
}

func TestLibraryVariableSetServiceSetAddGetDelete(t *testing.T) {
	service := createLibraryVariableSetService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, "libraryVariableSet"))
	require.Nil(t, resource)

	resource, err = service.Add(&LibraryVariableSet{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestLibraryVariableSet(t, service)
	require.NotNil(t, resource)

	resourceToCompare, err := service.GetByID(resource.GetID())
	require.NoError(t, err)
	require.NotNil(t, resourceToCompare)
	IsEqualLibraryVariableSets(t, resource, resourceToCompare)

	err = DeleteTestLibraryVariableSet(t, service, resource)
	require.NoError(t, err)
}

func TestLibraryVariableSetServiceAdd(t *testing.T) {
	service := createLibraryVariableSetService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, "libraryVariableSet"))
	require.Nil(t, resource)

	resource, err = service.Add(&LibraryVariableSet{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestLibraryVariableSet(t, service)
	err = DeleteTestLibraryVariableSet(t, service, resource)
	require.NoError(t, err)
}

func TestLibraryVariableSetServiceDeleteAll(t *testing.T) {
	service := createLibraryVariableSetService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		err = DeleteTestLibraryVariableSet(t, service, resource)
		assert.NoError(t, err)
	}
}

func TestLibraryVariableSetServiceGetAll(t *testing.T) {
	service := createLibraryVariableSetService(t)
	require.NotNil(t, service)

	libraryVariableSets := []LibraryVariableSet{}

	// create 30 test library variable sets (to be deleted)
	for i := 0; i < 30; i++ {
		libraryVariableSet := CreateTestLibraryVariableSet(t, service)
		require.NotNil(t, libraryVariableSet)
		libraryVariableSets = append(libraryVariableSets, *libraryVariableSet)
	}

	allLibraryVariableSets, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allLibraryVariableSets)
	require.True(t, len(allLibraryVariableSets) >= 30)

	for _, libraryVariableSet := range libraryVariableSets {
		require.NotNil(t, libraryVariableSet)
		require.NotEmpty(t, libraryVariableSet.GetID())
		err = DeleteTestLibraryVariableSet(t, service, &libraryVariableSet)
		require.NoError(t, err)
	}
}

func TestLibraryVariableSetServiceGetByID(t *testing.T) {
	service := createLibraryVariableSetService(t)
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
		IsEqualLibraryVariableSets(t, resource, resourceToCompare)
	}
}

func TestLibraryVariableSetServiceParameters(t *testing.T) {
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
			service := createLibraryVariableSetService(t)
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

func TestLibraryVariableSetServiceNew(t *testing.T) {
	ServiceFunction := NewLibraryVariableSetService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceLibraryVariableSetService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *LibraryVariableSetService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestLibraryVariableSetGetByPartialName(t *testing.T) {
	service := createLibraryVariableSetService(t)
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

func TestLibraryVariableSetServiceUpdate(t *testing.T) {
	service := createLibraryVariableSetService(t)
	require.NotNil(t, service)

	resource, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = service.Update(&LibraryVariableSet{})
	require.Error(t, err)
	require.Nil(t, resource)

	expected := CreateTestLibraryVariableSet(t, service)
	actual := UpdateLibraryVariableSet(t, service, expected)
	IsEqualLibraryVariableSets(t, expected, actual)
	defer DeleteTestLibraryVariableSet(t, service, expected)
}
