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

	// NOTE: a project is obtained through the Projects service API
	//
	// projects, err = client.Projects.GetAll()
	// project, err = client.Projects.GetByID(id)
	// project, err = client.Projects.GetByName(name)
	//
	// this ID value (below) is obtained via GetID()

	projectID := "project-id"

	// create runbook
	runbook := octopusdeploy.NewRunbook(name, projectID)

	// update any additional project fields here...

	// create runbook through Add(); returns error if fails
	createdRunbook, err := client.Runbooks.Add(runbook)
	if err != nil {
		_ = fmt.Errorf("error creating runbook: %v", err)
		return
	}

	fmt.Printf("runbook created: (%s)\n", createdRunbook.GetID())
}
