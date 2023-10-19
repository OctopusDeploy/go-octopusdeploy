package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/feeds"
)

func ChangeFeedExample() {
	var (
		apiKey      string = "API-YOUR_API_KEY"
		feedName    string = "nuget.org 3"
		newFeedName string = "nuget.org feed"
		octopusURL  string = "https://your_octopus_url"
		spaceName   string = "Default"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := client.NewClient(nil, apiURL, apiKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	query := feeds.FeedsQuery{
		PartialName: feedName,
	}

	// get feedList that match the name provided
	feedList, err := feeds.Get(client, "", query)
	if err != nil {
		_ = fmt.Errorf("error getting feed: %v", err)
		return
	}

	// select a specific feed
	feed := feedList.Items[0]

	// change feed name
	feed.SetName(newFeedName)

	// update feed
	_, err = feeds.Update(client, feed)
	if err != nil {
		_ = fmt.Errorf("error updating feed: %v", err)
		return
	}
}
