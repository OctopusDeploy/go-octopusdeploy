package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateRunbookExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// runbook values
		name      string = "runbook-name"
		projectID string = "project-id"
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

	// create runbook
	runbook := octopusdeploy.NewRunbook(name, projectID)
	createdRunbook, err := client.Runbooks.Add(runbook)
	if err != nil {
		_ = fmt.Errorf("error creating runbook: %v", err)
		return
	}

	fmt.Printf("runbook: (%s)\n", createdRunbook.GetID())
}
