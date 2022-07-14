package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteWorkerPoolExample provides an example of how to delete a dynamic
// worker pool from Octopus Deploy through the Go API client.
func DeleteWorkerPoolExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// worker pool values
		workerPoolID string = "worker-pool-id"
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

	// delete worker pool
	err = client.WorkerPools.DeleteByID(workerPoolID)
	if err != nil {
		_ = fmt.Errorf("error deleting worker pool: %v", err)
		return
	}

	fmt.Printf("worker pool deleted: (%s)\n", workerPoolID)
}
