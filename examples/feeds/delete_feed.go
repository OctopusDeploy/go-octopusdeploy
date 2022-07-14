package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/feeds"
)

func DeleteFeedExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		feedName   string = "nuget (ok to delete)"
		octopusURL string = "https://your_octopus_url"
		spaceName  string = "space-id"
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

	// get feeds that match the name provided
	feeds, err := client.Feeds.Get(query)
	if err != nil {
		_ = fmt.Errorf("error getting feed: %v", err)
		return
	}

	// select a specific feed
	feed := feeds.Items[0]

	// delete feed
	err = client.Feeds.DeleteByID(feed.GetID())
	if err != nil {
		_ = fmt.Errorf("error deleting feed: %v", err)
		return
	}

	fmt.Printf("feed deleted: (%s)\n", feed.GetID())
}
