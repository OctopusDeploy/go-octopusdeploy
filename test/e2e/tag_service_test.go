package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tagsets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestTag(t *testing.T) *tagsets.Tag {
	name := internal.GetRandomName()

	// TODO: randomize color
	createdTag := tagsets.NewTag(name, "#000000")
	require.NotNil(t, createdTag)

	return createdTag
}

func IsEqualTags(t *testing.T, expected *tagsets.Tag, actual *tagsets.Tag) {
	assert.Equal(t, expected.CanonicalTagName, actual.CanonicalTagName)
	assert.Equal(t, expected.Color, actual.Color)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.SortOrder, actual.SortOrder)
}

func TestTagServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	tagSet := CreateTestTagSet(t, client)
	require.NotNil(t, tagSet)
	defer DeleteTestTagSet(t, client, tagSet)

	tag := CreateTestTag(t)
	require.NotNil(t, tag)

	tagSet.Tags = append(tagSet.Tags, tag)
	updatedTagSet, err := client.TagSets.Update(tagSet)
	require.NotNil(t, updatedTagSet)
	require.NoError(t, err)

	tagSetToCompare, err := client.TagSets.GetByID(updatedTagSet.GetID())
	require.NotNil(t, tagSetToCompare)
	require.NoError(t, err)
	IsEqualTags(t, updatedTagSet.Tags[0], tagSetToCompare.Tags[0])
}
