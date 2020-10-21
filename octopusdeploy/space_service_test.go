package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createSpaceService(t *testing.T) *spaceService {
	service := newSpaceService(nil, TestURISpaces, TestURISpaceHome)
	testNewService(t, service, TestURISpaces, serviceSpaceService)
	return service
}

func CreateTestSpace(t *testing.T, service *spaceService) *Space {
	if service == nil {
		service = createSpaceService(t)
	}
	require.NotNil(t, service)

	user := CreateTestUser(t, nil)

	name := getShortRandomName()

	space := NewSpace(name)
	require.NoError(t, space.Validate())

	space.SpaceManagersTeamMembers = []string{user.GetID()}

	createdSpace, err := service.Add(space)
	require.NoError(t, err)
	require.NotNil(t, createdSpace)
	require.NotEmpty(t, createdSpace.GetID())
	require.Equal(t, name, createdSpace.Name)

	return createdSpace
}

func DeleteTestSpace(t *testing.T, service *spaceService, space *Space) error {
	require.NotNil(t, space)

	if service == nil {
		service = createSpaceService(t)
	}
	require.NotNil(t, service)

	err := service.DeleteByID(space.GetID())
	assert.NoError(t, err)

	return err
}

func IsEqualSpaces(t *testing.T, expected *Space, actual *Space) {
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

	// TODO: complete space comparison
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
}

func UpdateTestSpace(t *testing.T, service *spaceService, space *Space) *Space {
	if service == nil {
		service = createSpaceService(t)
	}
	require.NotNil(t, service)

	updatedSpace, err := service.Update(space)
	require.NoError(t, err)
	require.NotNil(t, updatedSpace)

	return updatedSpace
}

func TestSpaceSetAddGetDelete(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Space{})
	require.Error(t, err)
	require.Nil(t, resource)

	space := CreateTestSpace(t, service)
	require.NotNil(t, space)

	spaceToCompare, err := service.GetByID(space.GetID())
	require.NoError(t, err)
	require.NotNil(t, spaceToCompare)
	IsEqualSpaces(t, space, spaceToCompare)

	space.TaskQueueStopped = true
	UpdateTestSpace(t, service, space)

	err = DeleteTestSpace(t, service, space)
	require.NoError(t, err)
}

func TestSpaceServiceAdd(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Space{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestSpace(t, service)
	defer DeleteTestSpace(t, service, resource)
}

func TestSpaceServiceDeleteAll(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		defer DeleteTestSpace(t, service, resource)
	}
}

func TestSpaceServiceGetAll(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	spaces := []Space{}

	// create 30 test spaces (to be deleted)
	for i := 0; i < 30; i++ {
		space := CreateTestSpace(t, service)
		require.NotNil(t, space)
		spaces = append(spaces, *space)
	}

	allSpaces, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allSpaces)
	require.True(t, len(allSpaces) >= 30)

	for _, space := range spaces {
		require.NotNil(t, space)
		require.NotEmpty(t, space.GetID())
		err = DeleteTestSpace(t, service, &space)
		require.NoError(t, err)
	}
}

func TestSpaceServiceGetByID(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	id := getRandomName()
	resource, err := service.GetByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	require.Nil(t, resource)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := service.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualSpaces(t, resource, resourceToCompare)
	}
}

func TestSpaceServiceParameters(t *testing.T) {
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
			service := createSpaceService(t)
			require.NotNil(t, service)

			if isEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
				require.Nil(t, resource)

				resourceList, err := service.GetByPartialName(tc.parameter)
				require.Equal(t, createInvalidParameterError(operationGetByPartialName, parameterName), err)
				require.NotNil(t, resourceList)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))
			} else {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createResourceNotFoundError(serviceSpaceService, "ID", tc.parameter))
				require.Nil(t, resource)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, createResourceNotFoundError(serviceSpaceService, "ID", tc.parameter))
			}
		})
	}
}

func TestSpaceServiceNew(t *testing.T) {
	serviceFunction := newSpaceService
	client := &sling.Sling{}
	uriTemplate := emptyString
	homePath := emptyString
	serviceName := serviceSpaceService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string, string) *spaceService
		client      *sling.Sling
		uriTemplate string
		homePath    string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, homePath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, homePath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, homePath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.homePath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestSpaceGetByPartialName(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resources, err := service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, resources)
	require.Len(t, resources, 0)

	resources, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
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

func TestSpaceServiceUpdate(t *testing.T) {
	service := createSpaceService(t)
	require.NotNil(t, service)

	resource, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = service.Update(&Space{})
	require.Error(t, err)
	require.Nil(t, resource)

	expected := CreateTestSpace(t, service)
	actual := UpdateTestSpace(t, service, expected)
	IsEqualSpaces(t, expected, actual)
	defer DeleteTestSpace(t, service, expected)
}
