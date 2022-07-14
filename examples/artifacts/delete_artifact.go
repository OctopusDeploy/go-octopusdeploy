package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

// DeleteArtifactExample provides an example of how to delete an artifact from
// Octopus Deploy through the Go API client.
func DeleteArtifactExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// artifact values
		artifactID string = "artifact-id"
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

	// delete artifact
	err = client.Artifacts.DeleteByID(artifactID)
	if err != nil {
		_ = fmt.Errorf("error deleting artifact: %v", err)
		return
	}

	fmt.Printf("artifact deleted: (%s)\n", artifactID)
}
