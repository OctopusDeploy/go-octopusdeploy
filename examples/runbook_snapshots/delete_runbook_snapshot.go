package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteRunbookSnapshotExample provides an example of how to delete a runbook
// snapshot from Octopus Deploy through the Go API client.
func DeleteRunbookSnapshotExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// runbook snapshot values
		runbookSnapshotID string = "runbook-snapshot-id"
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

	// delete runbook snapshot
	err = client.RunbookSnapshots.DeleteByID(runbookSnapshotID)
	if err != nil {
		_ = fmt.Errorf("error deleting runbook snapshot: %v", err)
		return
	}

	fmt.Printf("runbook snapshot deleted: (%s)\n", runbookSnapshotID)
}
