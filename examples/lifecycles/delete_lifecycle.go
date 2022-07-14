package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteLifecycleExample provides an example of how to delete a lifecycle from
// Octopus Deploy through the Go API client.
func DeleteLifecycleExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// lifecycle values
		lifecycleID string = "lifecycle-id"
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

	// delete lifecycle
	err = client.Lifecycles.DeleteByID(lifecycleID)
	if err != nil {
		_ = fmt.Errorf("error deleting lifecycle: %v", err)
		return
	}

	fmt.Printf("lifecycle deleted: (%s)\n", lifecycleID)
}
