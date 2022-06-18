package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

func GetEnvironmentByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// environment values
		environmentID string = "environment-id"
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

	// get environment by its ID
	environment, err := client.Environments.GetByID(environmentID)
	if err != nil {
		_ = fmt.Errorf("error getting environment: %v", err)
		return
	}

	fmt.Printf("environment: (%s)\n", environment.GetID())
}
