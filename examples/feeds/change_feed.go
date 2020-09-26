package examples

import "github.com/OctopusDeploy/go-octopusdeploy/client"

func ChangeFeedExample() {
	var (
		// Declare working variables
		octopusURL    string = "https://youroctourl"
		octopusAPIKey string = "API-YOURAPIKEY"

		spaceName   string = "Default"
		feedName    string = "nuget.org 3"
		newFeedName string = "nuget.org feed"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)
	if err != nil {
		// TODO: handle error
	}

	// Get Feed instances that match the name provided
	feeds, err := client.Feeds.GetByPartialName(feedName)
	if err != nil {
		// TODO: handle error
	}

	// select a specific Feed instance
	feed := feeds[0]

	// Change feed name
	feed.Name = newFeedName

	// Update feed
	_, err = client.Feeds.Update(feed)
	if err != nil {
		// TODO: handle error
	}
}
