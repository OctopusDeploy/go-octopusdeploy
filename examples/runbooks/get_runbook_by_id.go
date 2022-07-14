package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

// GetRunbookByIDExample provides an example of how to get a runbook from
// Octopus Deploy by its ID through the Go API client.
func GetRunbookByIDExample() {
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

	client, err := client.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// get runbook by its ID
	runbook, err := client.Runbooks.GetByID(runbookID)
	if err != nil {
		_ = fmt.Errorf("error getting runbook: %v", err)
		return
	}

	fmt.Printf("runbook: (%s)\n", runbook.GetID())
}
