package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// GetSpaceByIDExample provides an example of how to get a space from Octopus
// Deploy by its ID through the Go API client.
func GetSpaceByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// space values
		spaceID string = "space-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, "")
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// get space by its ID
	space, err := client.Spaces.GetByID(spaceID)
	if err != nil {
		_ = fmt.Errorf("error getting space: %v", err)
		return
	}

	fmt.Printf("space: (%s)\n", space.GetID())
}
