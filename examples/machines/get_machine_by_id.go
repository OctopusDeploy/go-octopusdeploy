package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

func GetMachineByIDExample() {
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

	// get machine by its ID
	machine, err := client.Machines.GetByID(machineID)
	if err != nil {
		_ = fmt.Errorf("error getting machine: %v", err)
		return
	}

	fmt.Printf("machine: (%s)\n", machine.GetID())
}
