package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/access_management"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestTeam(t *testing.T, client *octopusdeploy.client) *access_management.Team {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	team := access_management.NewTeam(name)
	require.NoError(t, team.Validate())

	createdTeam, err := client.Teams.Add(team)
	require.NoError(t, err)
	require.NotNil(t, createdTeam)
	require.NotEmpty(t, createdTeam.GetID())
	require.Equal(t, name, createdTeam.Name)

	return createdTeam
}

func DeleteTestTeam(t *testing.T, client *octopusdeploy.client, team *access_management.Team) {
	require.NotNil(t, team)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Teams.Delete(team)
	assert.NoError(t, err)

	// verify the delete operation was successful
	teams, err := client.Teams.GetByID(team.GetID())
	assert.Error(t, err)
	assert.Nil(t, teams)
}

func IsEqualTeams(t *testing.T, expected *access_management.Team, actual *access_management.Team) {
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

func UpdateTeam(t *testing.T, client *octopusdeploy.client, team *access_management.Team) *access_management.Team {
	require.NotNil(t, team)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedTeam, err := client.Teams.Update(team)
	require.NoError(t, err)
	require.NotNil(t, updatedTeam)

	return updatedTeam
}

func TestTeamSetAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	team := CreateTestTeam(t, client)
	require.NotNil(t, team)
	defer DeleteTestTeam(t, client, team)

	teamToCompare, err := client.Teams.GetByID(team.GetID())
	require.NoError(t, err)
	require.NotNil(t, teamToCompare)
	IsEqualTeams(t, team, teamToCompare)
}

func TestTeamServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	teams, err := client.Teams.GetAll()
	require.NoError(t, err)
	require.NotNil(t, teams)

	for _, team := range teams {
		if team.CanBeDeleted {
			defer DeleteTestTeam(t, client, team)
		}
	}
}

func TestTeamServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 30 test teams (to be deleted)
	for i := 0; i < 30; i++ {
		team := CreateTestTeam(t, client)
		require.NotNil(t, team)
		defer DeleteTestTeam(t, client, team)
	}

	allTeams, err := client.Teams.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allTeams)
	require.True(t, len(allTeams) >= 30)
}

func TestTeamServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	team, err := client.Teams.GetByID(id)
	require.Error(t, err)
	require.Nil(t, team)

	teams, err := client.Teams.GetAll()
	require.NoError(t, err)
	require.NotNil(t, teams)

	for _, team := range teams {
		teamToCompare, err := client.Teams.GetByID(team.GetID())
		require.NoError(t, err)
		IsEqualTeams(t, team, teamToCompare)

		scopedUserRoles, err := client.Teams.GetScopedUserRoles(*team, services.SkipTakeQuery{Take: 1})
		require.NoError(t, err)
		require.NotNil(t, scopedUserRoles)
	}
}

func TestTeamServiceUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	createdTeam := CreateTestTeam(t, client)
	updatedTeam := UpdateTeam(t, client, createdTeam)
	IsEqualTeams(t, createdTeam, updatedTeam)
	defer DeleteTestTeam(t, client, updatedTeam)
}
