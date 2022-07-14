package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

// DeleteReleaseExample provides an example of how to delete a release from
// Octopus Deploy through the Go API client.
func DeleteReleaseExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// release values
		releaseID string = "release-id"
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

	// delete release
	err = client.Releases.DeleteByID(releaseID)
	if err != nil {
		_ = fmt.Errorf("error deleting release: %v", err)
		return
	}

	fmt.Printf("release deleted: (%s)\n", releaseID)
}
