package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateLifecycleExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		name string = "lifecycle-name"
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

	// create lifecycle
	lifecycle := services.NewLifecycle(name)

	// update any additional lifecycle fields here...

	// create lifecycle through Add(); returns error if fails
	createdLifecycle, err := client.Lifecycles.Add(lifecycle)
	if err != nil {
		_ = fmt.Errorf("error creating lifecycle: %v", err)
		return
	}

	fmt.Printf("lifecycle created: (%s)\n", createdLifecycle.GetID())
}
