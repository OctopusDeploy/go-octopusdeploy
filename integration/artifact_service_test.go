package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestArtifact(t *testing.T, client *octopusdeploy.Client) *octopusdeploy.Artifact {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	filename := "output.log"

	artifact := octopusdeploy.NewArtifact(filename)
	require.NotNil(t, artifact)

	createdArtifact, err := client.Artifacts.Add(artifact)
	require.NoError(t, err)
	require.NotNil(t, createdArtifact)
	require.NotEmpty(t, createdArtifact.GetID())

	return createdArtifact
}

func DeleteTestArtifact(t *testing.T, client *octopusdeploy.Client, artifact *octopusdeploy.Artifact) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Artifacts.DeleteByID(artifact.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedArtifact, err := client.Artifacts.GetByID(artifact.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedArtifact)

}

func AssertEqualArtifacts(t *testing.T, expected *octopusdeploy.Artifact, actual *octopusdeploy.Artifact) {
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

	// Artifact
	assert.Equal(t, expected.Created, actual.Created)
	assert.Equal(t, expected.Filename, actual.Filename)
	assert.Equal(t, expected.LogCorrelationID, actual.LogCorrelationID)
	assert.Equal(t, expected.ServerTaskID, actual.ServerTaskID)
	assert.Equal(t, expected.Source, actual.Source)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func TestArtifactServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	accounts, err := client.Artifacts.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		defer DeleteTestArtifact(t, client, account)
	}
}

// TODO: fix test
// func TestArtifactServiceGetAll(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	// create 30 test artifacts (to be deleted)
// 	for i := 0; i < 30; i++ {
// 		artifact := CreateTestArtifact(t, client)
// 		require.NotNil(t, artifact)
// 		defer DeleteTestArtifact(t, client, artifact)
// 	}

// 	allArtifacts, err := client.Artifacts.GetAll()
// 	require.NoError(t, err)
// 	require.NotNil(t, allArtifacts)
// 	require.True(t, len(allArtifacts) >= 30)
// }

func TestArtifactServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resources, err := client.Artifacts.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	if len(resources) > 0 {
		resourceToCompare, err := client.Artifacts.GetByID(resources[0].GetID())
		require.NoError(t, err)
		AssertEqualArtifacts(t, resources[0], resourceToCompare)
	}
}
