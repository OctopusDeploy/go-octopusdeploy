package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tagsets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestTagSet(t *testing.T, client *client.Client) *tagsets.TagSet {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	tagSet := tagsets.NewTagSet(name)
	require.NoError(t, tagSet.Validate())

	createdTagSet, err := client.TagSets.Add(tagSet)
	require.NoError(t, err)
	require.NotNil(t, createdTagSet)
	require.NotEmpty(t, createdTagSet.GetID())
	require.Equal(t, name, createdTagSet.Name)

	return createdTagSet
}

func DeleteTestTagSet(t *testing.T, client *client.Client, tagSet *tagsets.TagSet) {
	require.NotNil(t, tagSet)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.TagSets.DeleteByID(tagSet.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedTagSet, err := client.TagSets.GetByID(tagSet.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedTagSet)
}

func IsEqualTagSets(t *testing.T, expected *tagsets.TagSet, actual *tagsets.TagSet) {
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

	// TagSet
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func UpdateTagSet(t *testing.T, client *client.Client, tagSet *tagsets.TagSet) *tagsets.TagSet {
	require.NotNil(t, tagSet)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedTagSet, err := client.TagSets.Update(tagSet)
	require.NoError(t, err)
	require.NotNil(t, updatedTagSet)

	return updatedTagSet
}

func TestTagSetServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	tagSet := CreateTestTagSet(t, client)
	require.NotNil(t, tagSet)
	defer DeleteTestTagSet(t, client, tagSet)

	tagSetToCompare, err := client.TagSets.GetByID(tagSet.GetID())
	require.NoError(t, err)
	require.NotNil(t, tagSetToCompare)
	IsEqualTagSets(t, tagSet, tagSetToCompare)
}

func TestTagSetServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 30 test tag sets (to be deleted)
	for i := 0; i < 30; i++ {
		tagSet := CreateTestTagSet(t, client)
		require.NotNil(t, tagSet)
		defer DeleteTestTagSet(t, client, tagSet)
	}

	allTagSets, err := client.TagSets.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allTagSets)
	require.True(t, len(allTagSets) >= 30)
}

func TestTagSetServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	tagSet, err := client.TagSets.GetByID(id)
	require.Error(t, err)
	require.Nil(t, tagSet)

	tagSets, err := client.TagSets.GetAll()
	require.NoError(t, err)
	require.NotNil(t, tagSets)

	for _, tagSet := range tagSets {
		tagSetToCompare, err := client.TagSets.GetByID(tagSet.GetID())
		require.NoError(t, err)
		IsEqualTagSets(t, tagSet, tagSetToCompare)
	}
}

func TestTagSetServiceUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	createdTagSet := CreateTestTagSet(t, client)
	updatedTagSet := UpdateTagSet(t, client, createdTagSet)
	IsEqualTagSets(t, createdTagSet, updatedTagSet)
	defer DeleteTestTagSet(t, client, updatedTagSet)
}
