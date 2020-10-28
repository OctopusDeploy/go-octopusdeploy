package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func DeleteMachinePolicyExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// machine policy values
		machinePolicyID string = "machine-policy-id"
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

	// delete machine policy
	err = client.MachinePolicies.DeleteByID(machinePolicyID)
	if err != nil {
		_ = fmt.Errorf("error deleting machine policy: %v", err)
		return
	}

	fmt.Printf("machine policy deleted: (%s)\n", machinePolicyID)
}
