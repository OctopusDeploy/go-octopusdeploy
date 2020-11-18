package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteRunbookRunExample provides an example of how to delete a runbook run
// from Octopus Deploy through the Go API client.
func DeleteRunbookRunExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// runbook run values
		runbookRunID string = "runbook-run-id"
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

	// delete runbook run
	err = client.RunbookRuns.DeleteByID(runbookRunID)
	if err != nil {
		_ = fmt.Errorf("error deleting runbook run: %v", err)
		return
	}

	fmt.Printf("runbook run deleted: (%s)\n", runbookRunID)
}
