package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// CreateDynamicWorkerPoolExample provides an example of how to create a
// dynamic worker pool Octopus Deploy through the Go API client.
func CreateDynamicWorkerPoolExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// worker pool values
		name       = "worker-pool-name"
		workerType = octopusdeploy.WorkerTypeUbuntu1804
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

	// create dynamic worker pool
	dynamicWorkerPool, err := octopusdeploy.NewDynamicWorkerPool(name, workerType)
	if err != nil {
		_ = fmt.Errorf("error creating dynamic worker pool: %v", err)
		return
	}

	createdDynamicWorkerPool, err := client.WorkerPools.Add(dynamicWorkerPool)
	if err != nil {
		_ = fmt.Errorf("error creating dynamic worker pool: %v", err)
		return
	}

	fmt.Printf("dynamic worker pool created: (%s)\n", createdDynamicWorkerPool.GetID())
}
