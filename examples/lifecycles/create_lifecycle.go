package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/lifecycles"
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

	client, err := client.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// create lifecycle
	lifecycle := lifecycles.NewLifecycle(name)

	// update any additional lifecycle fields here...

	// create lifecycle through Add(); returns error if fails
	createdLifecycle, err := client.Lifecycles.Add(lifecycle)
	if err != nil {
		_ = fmt.Errorf("error creating lifecycle: %v", err)
		return
	}

	fmt.Printf("lifecycle created: (%s)\n", createdLifecycle.GetID())
}
