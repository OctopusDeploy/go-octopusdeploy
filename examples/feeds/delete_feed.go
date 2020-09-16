package examples

import "github.com/OctopusDeploy/go-octopusdeploy/client"

func DeleteFeedExample() {
	var (
		// Declare working variables
		octopusURL    string = "https://youroctourl"
		octopusAPIKey string = "API-YOURAPIKEY"

		spaceName string = "Default"
		feedName  string = "nuget to delete"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)

	if err != nil {
		// TODO: handle error
	}

	// Get Feed
	feed, err := client.Feeds.GetByName(feedName)

	if err != nil {
		// TODO: handle error
	}

	// Delete feed
	err = client.Feeds.Delete(feed.ID)

	if err != nil {
		// TODO: handle error
	}
}
