package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

func GetArtifactByIDExample() {
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

	// get artifact by its ID
	artifact, err := client.Artifacts.GetByID(artifactID)
	if err != nil {
		_ = fmt.Errorf("error getting artifact: %v", err)
		return
	}

	fmt.Printf("artifact: (%s)\n", artifact.GetID())
}
