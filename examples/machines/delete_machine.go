package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

// DeleteMachineExample provides an example of how to delete a machine from
// Octopus Deploy through the Go API client.
func DeleteMachineExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// machine values
		machineID string = "machine-id"
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

	// delete machine
	err = client.Machines.DeleteByID(machineID)
	if err != nil {
		_ = fmt.Errorf("error deleting machine: %v", err)
		return
	}

	fmt.Printf("machine deleted: (%s)\n", machineID)
}
