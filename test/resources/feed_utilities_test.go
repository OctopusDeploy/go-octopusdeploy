package resources

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/feeds"
	"github.com/stretchr/testify/require"
)

func TestToFeed(t *testing.T) {
	var feedResource *feeds.FeedResource
	feed, err := feeds.ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feedResource = feeds.NewFeedResource("", "")
	feed, err = feeds.ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feedResource = feeds.NewFeedResource(internal.GetRandomName(), "")
	feed, err = feeds.ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feedResource = feeds.NewFeedResource(internal.GetRandomName(), feeds.FeedTypeAwsElasticContainerRegistry)
	feed, err = feeds.ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feed, err = feeds.ToFeedResource(feedResource)
	require.NotNil(t, feed)
	require.NoError(t, err)
	require.EqualValues(t, feed.GetFeedType(), "AwsElasticContainerRegistry")
}
