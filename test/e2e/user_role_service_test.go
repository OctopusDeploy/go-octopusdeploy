package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/userroles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestUserRole(t *testing.T, client *client.Client) *userroles.UserRole {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	userRole := userroles.NewUserRole(name)
	require.NoError(t, userRole.Validate())

	createdUserRole, err := client.UserRoles.Add(userRole)
	require.NotNil(t, createdUserRole)
	require.NoError(t, err)

	userRoleToCompare, err := client.UserRoles.GetByID(createdUserRole.GetID())
	require.NotNil(t, userRoleToCompare)
	require.NoError(t, err)

	AssertEqualUserRoles(t, createdUserRole, userRoleToCompare)

	return createdUserRole
}

func DeleteTestUserRole(t *testing.T, client *client.Client, userRole *userroles.UserRole) {
	require.NotNil(t, userRole)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.UserRoles.DeleteByID(userRole.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedUserRole, err := client.UserRoles.GetByID(userRole.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedUserRole)
}

func AssertEqualUserRoles(t *testing.T, expected *userroles.UserRole, actual *userroles.UserRole) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, internal.IsLinksEqual(expected.Links, actual.Links))
}

func TestUserRoleServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	userRole := CreateTestUserRole(t, client)
	require.NotNil(t, userRole)
	defer DeleteTestUserRole(t, client, userRole)

	userRoles, err := client.UserRoles.GetAll()
	require.NoError(t, err)
	require.NotNil(t, userRoles)

	for _, userRole := range userRoles {
		query := userroles.UserRolesQuery{
			IDs: []string{userRole.GetID()},
		}
		userRolesToCompare, err := client.UserRoles.Get(query)
		require.NoError(t, err)
		require.NotNil(t, userRolesToCompare)
		for _, userRoleToCompare := range userRolesToCompare.Items {
			AssertEqualUserRoles(t, userRole, userRoleToCompare)
		}

		userRoleToCompare, err := client.UserRoles.GetByID(userRole.GetID())
		require.NoError(t, err)
		require.NotNil(t, userRoleToCompare)
		AssertEqualUserRoles(t, userRole, userRoleToCompare)
	}
}

func TestUserRoleServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	userRoles, err := client.UserRoles.GetAll()
	require.NoError(t, err)
	require.NotNil(t, userRoles)

	for _, userRole := range userRoles {
		if userRole.CanBeDeleted {
			defer DeleteTestUserRole(t, client, userRole)
		}
	}
}

func TestUserRoleServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	userRoles, err := client.UserRoles.GetAll()
	require.NotNil(t, userRoles)
	require.NoError(t, err)

	for _, userRole := range userRoles {
		userRole, err := client.UserRoles.GetByID(userRole.GetID())
		require.NoError(t, err)
		require.NotNil(t, userRole)
	}
}
