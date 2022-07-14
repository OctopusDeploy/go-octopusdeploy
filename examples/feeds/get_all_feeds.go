package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

func GetAllFeedsExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceName  string = "Default"
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

	// Get all feeds
	feeds, err := client.Feeds.GetAll()
	if err != nil {
		_ = fmt.Errorf("error getting feeds: %v", err)
		return
	}

	for _, feed := range feeds {
		fmt.Printf("Feed ID: %s\n", feed.GetID())
		fmt.Printf("Feed Name: %s\n", feed.GetName())
		fmt.Printf("Feed Type: %s\n", feed.GetFeedType())
		fmt.Println()
	}
}
