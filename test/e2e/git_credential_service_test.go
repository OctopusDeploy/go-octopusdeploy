package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestGitCredentialResource(t *testing.T, client *client.Client) *credentials.Resource {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"

	username := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())
	usernamePassword := credentials.NewUsernamePassword(username, password)

	resource := credentials.NewResource(name, usernamePassword)
	resource.Description = description

	require.NoError(t, resource.Validate())

	createdResource, err := client.GitCredentials.Add(resource)
	require.NoError(t, err)
	require.NotNil(t, createdResource)
	require.NotEmpty(t, createdResource.GetID())

	return createdResource
}

func DeleteTestGitCredentialResource(t *testing.T, client *client.Client, resource *credentials.Resource) {
	require.NotNil(t, resource)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.GitCredentials.DeleteByID(resource.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedResource, err := client.GitCredentials.GetByID(resource.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedResource)
}

func IsEqualCredentials(t *testing.T, expected *credentials.Resource, actual *credentials.Resource) {
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

	// Resource
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
}

func UpdateTestGitCredentialResource(t *testing.T, client *client.Client, resource *credentials.Resource) *credentials.Resource {
	require.NotNil(t, resource)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedResource, err := client.GitCredentials.Update(resource)
	assert.NoError(t, err)
	require.NotNil(t, updatedResource)

	return updatedResource
}

func TestCredentialServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	createdResource := CreateTestGitCredentialResource(t, client)
	defer DeleteTestGitCredentialResource(t, client, createdResource)

	resource, err := client.GitCredentials.GetByID(createdResource.GetID())
	require.NoError(t, err)
	require.NotNil(t, resource)

	credentials, err := client.GitCredentials.Get(credentials.Query{Name: resource.GetName()})
	require.NoError(t, err)
	require.NotNil(t, credentials)

	resourceToCompare := credentials.Items[0]
	IsEqualCredentials(t, resource, resourceToCompare)
}

func TestCredentialServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 30 test credentials (to be deleted)
	for i := 0; i < 30; i++ {
		environment := CreateTestGitCredentialResource(t, client)
		require.NotNil(t, environment)
		defer DeleteTestGitCredentialResource(t, client, environment)
	}

	allCredentials, err := client.GitCredentials.Get(credentials.Query{})
	require.NoError(t, err)
	require.NotNil(t, allCredentials)
	require.True(t, len(allCredentials.Items) >= 30)

	for _, credential := range allCredentials.Items {
		credentialToCompare, err := client.GitCredentials.GetByID(credential.GetID())
		require.NoError(t, err)
		require.NotNil(t, credential)
		require.NotEmpty(t, credential.GetID())
		IsEqualCredentials(t, credential, credentialToCompare)
	}
}

func TestCredentialServiceUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resource := CreateTestGitCredentialResource(t, client)
	defer DeleteTestGitCredentialResource(t, client, resource)

	resource, err := client.GitCredentials.GetByID(resource.GetID())
	require.NotNil(t, resource)
	require.NoError(t, err)

	newDescription := internal.GetRandomName()
	newName := internal.GetRandomName()

	resource.Description = newDescription
	resource.Name = newName

	updatedCredential := UpdateTestGitCredentialResource(t, client, resource)
	require.NotNil(t, updatedCredential)
	require.NotEmpty(t, updatedCredential.GetID())
	require.Equal(t, updatedCredential.GetID(), updatedCredential.GetID())
	require.Equal(t, newDescription, updatedCredential.Description)
	require.Equal(t, newName, updatedCredential.Name)
}

func TestCredentialServiceAddGetDelete_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	createdResource := CreateTestGitCredentialResource_NewClient(t, client)
	defer DeleteTestGitCredentialResource_NewClient(t, client, createdResource)

	resource, err := credentials.GetByID(client, createdResource.SpaceID, createdResource.GetID())
	require.NoError(t, err)
	require.NotNil(t, resource)

	credentials, err := credentials.Get(client, resource.SpaceID, credentials.Query{Name: resource.GetName()})
	require.NoError(t, err)
	require.NotNil(t, credentials)

	resourceToCompare := credentials.Items[0]
	IsEqualCredentials(t, resource, resourceToCompare)
}

func CreateTestGitCredentialResource_NewClient(t *testing.T, client *client.Client) *credentials.Resource {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"

	username := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())
	usernamePassword := credentials.NewUsernamePassword(username, password)

	resource := credentials.NewResource(name, usernamePassword)
	resource.Description = description

	require.NoError(t, resource.Validate())

	createdResource, err := credentials.Add(client, resource)
	require.NoError(t, err)
	require.NotNil(t, createdResource)
	require.NotEmpty(t, createdResource.GetID())

	return createdResource
}

func DeleteTestGitCredentialResource_NewClient(t *testing.T, client *client.Client, resource *credentials.Resource) {
	require.NotNil(t, resource)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := credentials.DeleteByID(client, resource.SpaceID, resource.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedResource, err := credentials.GetByID(client, resource.SpaceID, resource.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedResource)
}
