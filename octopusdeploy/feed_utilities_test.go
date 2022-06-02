package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToFeed(t *testing.T) {
	var feedResource *FeedResource
	feed, err := ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feedResource = NewFeedResource("", "")
	feed, err = ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feedResource = NewFeedResource(getRandomName(), "")
	feed, err = ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feedResource = NewFeedResource(getRandomName(), FeedTypeAwsElasticContainerRegistry)
	feed, err = ToFeed(feedResource)
	require.Nil(t, feed)
	require.Error(t, err)

	feed, err = ToFeedResource(feedResource)
	require.NotNil(t, feed)
	require.NoError(t, err)
	require.EqualValues(t, feed.GetFeedType(), "AwsElasticContainerRegistry")
}
