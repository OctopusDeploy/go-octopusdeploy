package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// CreateRunbookSnapshotExample provides an example of how to create a runbook
// snapshot in Octopus Deploy through the Go API client.
func CreateRunbookExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		name string = "runbook-name"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, "")
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// NOTE: a project may be obtained through the Projects service API
	//
	// projects, err = client.Projects.Get(...)
	projectID := "project-id"

	// NOTE: a runbook may be obtained through the Runbooks service API
	//
	// runbooks, err = client.Runbooks.Get(...)
	runbookID := "runbook-id"

	// create runbook
	runbookSnapshot := octopusdeploy.NewRunbookSnapshot(name, projectID, runbookID)

	// update any additional project fields here...

	// create runbook through Add(); returns error if fails
	createdRunbookSnapshot, err := client.RunbookSnapshots.Add(runbookSnapshot)
	if err != nil {
		_ = fmt.Errorf("error creating runbook snapshot: %v", err)
		return
	}

	fmt.Printf("runbook snapshot created: (%s)\n", createdRunbookSnapshot.GetID())
}
