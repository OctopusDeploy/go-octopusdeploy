package examples

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
)

func GetAllFeedsExample() {
	var (
		// Declare working variables
		octopusURL    string = "https://youroctourl"
		octopusAPIKey string = "API-YOURAPIKEY"

		spaceName string = "Default"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)
	if err != nil {
		// TODO: handle error
	}

	// Get all Feeds
	feeds, err := client.Feeds.GetAll()
	if err != nil {
		// TODO: handle error
	}

	for _, feed := range feeds {
		fmt.Printf("Feed ID: %s\n", feed.GetID())
		fmt.Printf("Feed Name: %s\n", feed.GetName())
		fmt.Printf("Feed Type: %s\n", feed.GetFeedType())
		fmt.Println()
	}
}
