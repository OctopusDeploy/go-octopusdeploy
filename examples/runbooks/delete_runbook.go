package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)
// DeleteRunbookExample provides an example of how to delete a runbook from
// Octopus Deploy through the Go API client.
func DeleteRunbookExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// runbook values
		runbookID string = "runbook-id"
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

	// delete runbook
	err = client.Runbooks.DeleteByID(runbookID)
	if err != nil {
		_ = fmt.Errorf("error deleting runbook: %v", err)
		return
	}

	fmt.Printf("runbook deleted: (%s)\n", runbookID)
}
