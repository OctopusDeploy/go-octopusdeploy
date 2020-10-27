package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
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

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	query := octopusdeploy.FeedsQuery{
		PartialName: feedName,
	}

	// get feeds that match the name provided
	feeds, err := client.Feeds.Get(query)
	if err != nil {
		_ = fmt.Errorf("error getting feed: %v", err)
		return
	}

	// select a specific feed
	feed := feeds.Items[0]

	// change feed name
	feed.SetName(newFeedName)

	// update feed
	_, err = client.Feeds.Update(feed)
	if err != nil {
		_ = fmt.Errorf("error updating feed: %v", err)
		return
	}
}
