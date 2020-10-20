package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func DeleteFeedExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		feedName   string = "nuget to delete"
		octopusURL string = "https://your_octopus_url"
		spaceName  string = "Default"
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

	// Get Feed instances that match the name provided
	feeds, err := client.Feeds.GetByPartialName(feedName)
	if err != nil {
		_ = fmt.Errorf("error getting feed: %v", err)
		return
	}

	// select a specific Feed instance
	feed := feeds[0]

	// Delete feed
	err = client.Feeds.DeleteByID(feed.GetID())
	if err != nil {
		_ = fmt.Errorf("error deleting feed: %v", err)
		return
	}
}
