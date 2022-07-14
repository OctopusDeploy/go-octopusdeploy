package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteChannelExample provides an example of how to delete a channel from
// Octopus Deploy through the Go API client.
func DeleteChannelExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// channel values
		channelID string = "channel-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := client.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// delete channel
	err = client.Channels.DeleteByID(channelID)
	if err != nil {
		_ = fmt.Errorf("error deleting channel: %v", err)
		return
	}

	fmt.Printf("channel deleted: (%s)\n", channelID)
}
