package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

// DeleteWorkerExample provides an example of how to delete a worker from
// Octopus Deploy through the Go API client.
func DeleteWorkerExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// worker values
		workerID string = "worker-id"
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

	// delete worker
	err = client.Workers.DeleteByID(workerID)
	if err != nil {
		_ = fmt.Errorf("error deleting worker: %v", err)
		return
	}

	fmt.Printf("worker deleted: (%s)\n", workerID)
}
