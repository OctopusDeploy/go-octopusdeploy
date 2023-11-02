package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/scopeduserroles"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/userroles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScopedUserRoleService_AddGetUpdateDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// Create user role with at least 1 space permission
	name := internal.GetRandomName()
	userrole := userroles.NewUserRole(name)
	userrole.GrantedSpacePermissions = []string{"ProjectView"}
	require.NoError(t, userrole.Validate())

	createdUserRole, err := userroles.Add(client, userrole)
	require.NotNil(t, createdUserRole)
	require.NoError(t, err)

	team := CreateTestTeam_NewClient(t, client)
	require.NotNil(t, team)

	// Add
	scopedUserRole := scopeduserroles.NewScopedUserRole(createdUserRole.ID, team.ID)
	scopedUserRole.SpaceID = client.GetSpaceID()
	require.NoError(t, scopedUserRole.Validate())

	createdRole, err := scopeduserroles.Add(client, scopedUserRole)
	require.NotNil(t, createdRole)
	require.NoError(t, err)

	// Get
	userRoleToCompare, err := scopeduserroles.GetByID(client, createdRole.SpaceID, createdRole.GetID())
	require.NotNil(t, userRoleToCompare)
	require.NoError(t, err)

	AssertEqualScopedUserRoles(t, createdRole, userRoleToCompare)

	// Update
	updatedRole, err := scopeduserroles.Update(client, createdRole)
	require.Nil(t, err)
	assert.EqualValues(t, createdRole.ID, updatedRole.ID)
	assert.EqualValues(t, createdRole.TeamID, updatedRole.TeamID)

	t.Cleanup(func() {
		DeleteScopedUserRole(t, client, createdRole)
		DeleteTestUserRole_NewClient(t, client, createdUserRole)
		DeleteTestTeam_NewClient(t, client, team)
	})
}

func DeleteScopedUserRole(t *testing.T, client *client.Client, scopedUserRole *scopeduserroles.ScopedUserRole) {
	require.NotNil(t, scopedUserRole)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := scopeduserroles.DeleteByID(client, scopedUserRole.SpaceID, scopedUserRole.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedUserRole, err := scopeduserroles.GetByID(client, scopedUserRole.SpaceID, scopedUserRole.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedUserRole)
}

func AssertEqualScopedUserRoles(t *testing.T, expected *scopeduserroles.ScopedUserRole, actual *scopeduserroles.ScopedUserRole) {
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.Equal(t, expected.UserRoleID, actual.UserRoleID)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.TeamID, actual.TeamID)
	assert.True(t, internal.IsLinksEqual(expected.Links, actual.Links))
}
