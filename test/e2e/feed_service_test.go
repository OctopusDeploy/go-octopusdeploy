package e2e

import (
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/feeds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualFeeds(t *testing.T, expected feeds.IFeed, actual feeds.IFeed) {
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
	fmt.Println(expected.GetLinks())
	fmt.Println(actual.GetLinks())
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

	// IFeed
	assert.Equal(t, expected.GetFeedType(), actual.GetFeedType())
	assert.Equal(t, expected.GetName(), actual.GetName())
	assert.Equal(t, expected.GetPackageAcquisitionLocationOptions(), actual.GetPackageAcquisitionLocationOptions())
	assert.Equal(t, expected.GetPassword(), actual.GetPassword())
	assert.Equal(t, expected.GetSpaceID(), actual.GetSpaceID())
	assert.Equal(t, expected.GetUsername(), actual.GetUsername())

	// TODO: compare remaining values
}

func CreateTestAwsElasticContainerRegistry(t *testing.T, client *client.Client) feeds.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	// the feed client validates the input parameters and attempts to make a
	// connection to the Elastic Container Registry (ECR) -- therefore, a valid
	// set of credentials (i.e. access key, secret key) must be provided along
	// with a valid region (i.e. "ap-southeast-2")

	accessKey := "access-key"
	secretKey := core.NewSensitiveValue("secret-key")
	region := "ap-southeast-2"

	feed, err := feeds.NewAwsElasticContainerRegistry(internal.GetRandomName(), accessKey, secretKey, region)
	require.NoError(t, err)

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestGitHubRepositoryFeed(t *testing.T, client *client.Client) feeds.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed, err := feeds.NewGitHubRepositoryFeed(internal.GetRandomName())
	require.NoError(t, err)

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestHelmFeed(t *testing.T, client *client.Client) feeds.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed, err := feeds.NewHelmFeed(internal.GetRandomName())
	require.NoError(t, err)

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestMavenFeed(t *testing.T, client *client.Client) feeds.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed, err := feeds.NewMavenFeed(internal.GetRandomName())
	require.NoError(t, err)

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestNuGetFeed(t *testing.T, client *client.Client) feeds.IFeed {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	feed, err := feeds.NewNuGetFeed(internal.GetRandomName(), "https://api.nuget.org/v3/index.json")
	require.NoError(t, err)

	resource, err := client.Feeds.Add(feed)
	require.NoError(t, err)

	return resource
}

func DeleteTestFeed(t *testing.T, client *client.Client, feed feeds.IFeed) {
	require.NotNil(t, feed)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	// built-in feeds cannot be deleted
	if feed.GetFeedType() != feeds.FeedTypeBuiltIn && feed.GetFeedType() != feeds.FeedTypeOctopusProject {
		err := client.Feeds.DeleteByID(feed.GetID())
		require.NoError(t, err)

		// verify the delete operation was successful
		deletedFeed, err := client.Feeds.GetByID(feed.GetID())
		require.Error(t, err)
		require.Nil(t, deletedFeed)
	}
}

func TestFeedServiceAdd(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// the following code is commented out due to the validation conducted by
	// the feed client

	// feed := CreateTestAwsElasticContainerRegistry(t, client)
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

func TestFeedServiceCRUD(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	gitHubRepositoryFeed := CreateTestGitHubRepositoryFeed(t, client).(*feeds.GitHubRepositoryFeed)
	require.NotNil(t, gitHubRepositoryFeed)
	defer DeleteTestFeed(t, client, gitHubRepositoryFeed)

	actual, err := client.Feeds.GetByID(gitHubRepositoryFeed.GetID())
	require.NoError(t, err)
	AssertEqualFeeds(t, gitHubRepositoryFeed, actual)

	name := internal.GetRandomName()
	gitHubRepositoryFeed.SetName(name)

	actual, err = client.Feeds.Update(gitHubRepositoryFeed)
	require.NoError(t, err)
	AssertEqualFeeds(t, gitHubRepositoryFeed, actual)

	gitHubRepositoryFeed = actual.(*feeds.GitHubRepositoryFeed)
	require.NotNil(t, gitHubRepositoryFeed)

	expected := CreateTestHelmFeed(t, client)
	require.NotNil(t, expected)
	defer DeleteTestFeed(t, client, expected)

	actual, err = client.Feeds.GetByID(expected.GetID())
	require.NoError(t, err)
	AssertEqualFeeds(t, expected, actual)

	name = internal.GetRandomName()
	expected.SetName(name)

	actual, err = client.Feeds.Update(expected)
	require.NoError(t, err)
	AssertEqualFeeds(t, expected, actual)

	expected = CreateTestMavenFeed(t, client)
	require.NotNil(t, expected)
	defer DeleteTestFeed(t, client, expected)

	actual, err = client.Feeds.GetByID(expected.GetID())
	require.NoError(t, err)
	AssertEqualFeeds(t, expected, actual)

	name = internal.GetRandomName()
	expected.SetName(name)

	actual, err = client.Feeds.Update(expected)
	require.NoError(t, err)
	AssertEqualFeeds(t, expected, actual)

	expected = CreateTestNuGetFeed(t, client)
	require.NotNil(t, expected)
	defer DeleteTestFeed(t, client, expected)

	actual, err = client.Feeds.GetByID(expected.GetID())
	require.NoError(t, err)
	AssertEqualFeeds(t, expected, actual)

	name = internal.GetRandomName()
	expected.SetName(name)

	actual, err = client.Feeds.Update(expected)
	require.NoError(t, err)
	AssertEqualFeeds(t, expected, actual)
}

// TODO: fix test
// func TestFeedServiceDeleteAll(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	feeds, err := client.Feeds.GetAll()
// 	require.NotNil(t, feeds)
// 	require.NoError(t, err)

// 	for _, feed := range feeds {
// 		defer DeleteTestFeed(t, client, feed)
// 	}
// }

func TestFeedServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	count := 10

	for i := 0; i < count; i++ {
		feed := CreateTestNuGetFeed(t, client)
		require.NotNil(t, feed)
		defer DeleteTestFeed(t, client, feed)
	}

	feeds, err := client.Feeds.GetAll()
	require.NotNil(t, feeds)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(feeds), count)
}

func TestFeedServiceGetBuiltInFeedStatistics(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	builtInFeedStatistics, err := client.Feeds.GetBuiltInFeedStatistics()
	require.NotNil(t, builtInFeedStatistics)
	require.NoError(t, err)
}

func TestFeedServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	feed, err := client.Feeds.GetByID(id)
	require.Error(t, err)
	require.Nil(t, feed)
}

func TestFeedServiceSearchPackages(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	feed := CreateTestGitHubRepositoryFeed(t, client)
	require.NotNil(t, feed)
	defer DeleteTestFeed(t, client, feed)

	searchPackagesQuery := feeds.SearchPackagesQuery{
		Term: "ngnix",
		Take: 10,
	}

	packageDescriptions, err := client.Feeds.SearchPackages(feed, searchPackagesQuery)
	require.NotNil(t, packageDescriptions)
	require.NoError(t, err)
}

func TestFeedServiceSearchPackageVersions(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	feed := CreateTestNuGetFeed(t, client)
	require.NotNil(t, feed)

	packageDescriptions, err := client.Feeds.SearchPackages(feed, feeds.SearchPackagesQuery{
		Term: "OctopusTools",
	})
	require.NotNil(t, packageDescriptions)
	require.NoError(t, err)

	for _, v := range packageDescriptions.Items {
		packageVersions, err := client.Feeds.SearchPackageVersions(v, feeds.SearchPackageVersionsQuery{
			FeedID:    feed.GetID(),
			PackageID: v.ID,
		})
		require.NotNil(t, packageVersions)
		require.NoError(t, err)
	}
}
