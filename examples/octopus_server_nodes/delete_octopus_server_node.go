package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteOctopusServerNodeExample provides an example of how to delete an
// Octopus server node from Octopus Deploy through the Go API client.
func DeleteOctopusServerNodeExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// octopus server node values
		octopusServerNodeID string = "octopus-server-node-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// delete octopus server node
	err = client.OctopusServerNodes.DeleteByID(octopusServerNodeID)
	if err != nil {
		_ = fmt.Errorf("error deleting octopus server node: %v", err)
		return
	}

	fmt.Printf("octopus server node deleted: (%s)\n", octopusServerNodeID)
}
