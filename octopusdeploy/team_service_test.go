package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTeamService(t *testing.T) *teamService {
	service := newTeamService(nil, TestURITeams)
	testNewService(t, service, TestURITeams, serviceTeamService)
	return service
}

func CreateTestTeam(t *testing.T, service *teamService) *Team {
	if service == nil {
		service = createTeamService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()

	team := NewTeam(name)
	require.NoError(t, team.Validate())

	createdTeam, err := service.Add(team)
	require.NoError(t, err)
	require.NotNil(t, createdTeam)
	require.NotEmpty(t, createdTeam.GetID())
	require.Equal(t, name, createdTeam.Name)

	return createdTeam
}

func DeleteTestTeam(t *testing.T, service *teamService, team *Team) error {
	require.NotNil(t, team)

	if service == nil {
		service = createTeamService(t)
	}
	require.NotNil(t, service)

	err := service.DeleteByID(team.GetID())
	assert.NoError(t, err)

	return err
}

func IsEqualTeams(t *testing.T, expected *Team, actual *Team) {
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

	// team
	assert.Equal(t, expected.CanBeDeleted, actual.CanBeDeleted)
	assert.Equal(t, expected.CanBeRenamed, actual.CanBeRenamed)
	assert.Equal(t, expected.CanChangeMembers, actual.CanChangeMembers)
	assert.Equal(t, expected.CanChangeRoles, actual.CanChangeRoles)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.ExternalSecurityGroups, actual.ExternalSecurityGroups)
	assert.Equal(t, expected.MemberUserIDs, actual.MemberUserIDs)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func UpdateTeam(t *testing.T, service *teamService, team *Team) *Team {
	if service == nil {
		service = createTeamService(t)
	}
	require.NotNil(t, service)

	updatedTeam, err := service.Update(team)
	require.NoError(t, err)
	require.NotNil(t, updatedTeam)

	return updatedTeam
}

func TestTeamSetAddGetDelete(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Team{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestTeam(t, service)
	require.NotNil(t, resource)
	defer DeleteTestTeam(t, service, resource)

	resourceToCompare, err := service.GetByID(resource.GetID())
	require.NoError(t, err)
	require.NotNil(t, resourceToCompare)
	IsEqualTeams(t, resource, resourceToCompare)
}

func TestTeamServiceAdd(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Team{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestTeam(t, service)
	err = DeleteTestTeam(t, service, resource)
	require.NoError(t, err)
}

func TestTeamServiceDeleteAll(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		err = DeleteTestTeam(t, service, resource)
		require.NoError(t, err)
	}
}

func TestTeamServiceGetAll(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	teams := []Team{}

	// create 30 test teams (to be deleted)
	for i := 0; i < 30; i++ {
		team := CreateTestTeam(t, service)
		require.NotNil(t, team)
		teams = append(teams, *team)
	}

	allTeams, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allTeams)
	require.True(t, len(allTeams) >= 30)

	for _, team := range teams {
		require.NotNil(t, team)
		require.NotEmpty(t, team.GetID())
		err = DeleteTestTeam(t, service, &team)
		require.NoError(t, err)
	}
}

func TestTeamServiceGetByID(t *testing.T) {
	service := createTeamService(t)
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
		IsEqualTeams(t, resource, resourceToCompare)
	}
}

func TestTeamServiceParameters(t *testing.T) {
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
			service := createTeamService(t)
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
				require.Equal(t, err, createResourceNotFoundError(serviceTeamService, "ID", tc.parameter))
				require.Nil(t, resource)

				err = service.DeleteByID(tc.parameter)
				require.Error(t, err)
				require.Equal(t, err, createResourceNotFoundError(serviceTeamService, "ID", tc.parameter))
			}
		})
	}
}

func TestTeamServiceNew(t *testing.T) {
	serviceFunction := newTeamService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceTeamService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *teamService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestTeamGetByPartialName(t *testing.T) {
	service := createTeamService(t)
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

func TestTeamServiceUpdate(t *testing.T) {
	service := createTeamService(t)
	require.NotNil(t, service)

	resource, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = service.Update(&Team{})
	require.Error(t, err)
	require.Nil(t, resource)

	expected := CreateTestTeam(t, service)
	actual := UpdateTeam(t, service, expected)
	IsEqualTeams(t, expected, actual)
	defer DeleteTestTeam(t, service, expected)
}
