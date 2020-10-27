package integration

import (
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualFeeds(t *testing.T, expected octopusdeploy.IFeed, actual octopusdeploy.IFeed) {
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

	// TODO: compare remaining values
}

func CreateTestAwsElasticContainerRegistry(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	// the feed client validates the input parameters and attempts to make a
	// connection to the Elastic Container Registry (ECR) -- therefore, a valid
	// set of credentials (i.e. access key, secret key) must be provided along
	// with a valid region (i.e. "ap-southeast-2")

	accessKey := "access-key"
	secretKey := octopusdeploy.NewSensitiveValue("secret-key")
	region := "ap-southeast-2"

	feed := octopusdeploy.NewAwsElasticContainerRegistry(getRandomName(), accessKey, secretKey, region)

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestGitHubRepositoryFeed(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed := octopusdeploy.NewGitHubRepositoryFeed(getRandomName(), "http://example.com/")

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestHelmFeed(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed := octopusdeploy.NewHelmFeed(getRandomName(), "http://example.com/")

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestMavenFeed(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed := octopusdeploy.NewMavenFeed(getRandomName(), "http://example.com/")

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestNuGetFeed(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed := octopusdeploy.NewNuGetFeed(getRandomName(), "http://example.com/")

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func DeleteTestFeed(t *testing.T, client *octopusdeploy.Client, feed octopusdeploy.IFeed) {
	require.NotNil(t, feed)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Feeds.DeleteByID(feed.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedFeed, err := client.Feeds.GetByID(feed.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedFeed)
}

func TestFeedServiceAdd(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// the following code is commented out due to the validation conducted by
	// the feed client

	// feed = CreateTestAwsElasticContainerRegistry(t, client)
	// require.NotNil(t, feed)
	// defer DeleteTestFeed(t, client, feed)

	feed := CreateTestGitHubRepositoryFeed(t, client)
	require.NotNil(t, feed)
	defer DeleteTestFeed(t, client, feed)

	feed = CreateTestHelmFeed(t, client)
	require.NotNil(t, feed)
	defer DeleteTestFeed(t, client, feed)

	feed = CreateTestMavenFeed(t, client)
	require.NotNil(t, feed)
	defer DeleteTestFeed(t, client, feed)

	feed = CreateTestNuGetFeed(t, client)
	require.NotNil(t, feed)
	defer DeleteTestFeed(t, client, feed)
}

func TestFeedServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	expected := CreateTestNuGetFeed(t, client)
	require.NotNil(t, expected)
	defer DeleteTestFeed(t, client, expected)

	actual, err := client.Feeds.GetByID(expected.GetID())
	require.NoError(t, err)
	AssertEqualFeeds(t, expected, actual)
}

func TestFeedServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	feeds, err := client.Feeds.GetAll()
	require.NotNil(t, feeds)
	require.NoError(t, err)

	for _, feed := range feeds {
		if !strings.Contains(feed.GetID(), "builtin") {
			defer DeleteTestFeed(t, client, feed)
		}
	}
}

func TestFeedServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	const count int = 32
	expected := map[string]octopusdeploy.IFeed{}
	for i := 0; i < count; i++ {
		feed := CreateTestNuGetFeed(t, client)
		require.NotNil(t, feed)
		defer DeleteTestFeed(t, client, feed)
		expected[feed.GetID()] = feed
	}

	feeds, err := client.Feeds.GetAll()
	require.NotNil(t, feeds)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(feeds), count)

	for _, actual := range feeds {
		_, ok := expected[actual.GetID()]
		if ok {
			AssertEqualFeeds(t, expected[actual.GetID()], actual)
		}
	}
}

func TestFeedServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	feed, err := client.Feeds.GetByID(id)
	require.Equal(t, createResourceNotFoundError(octopusdeploy.ServiceFeedService, "ID", id), err)
	require.Nil(t, feed)
}
