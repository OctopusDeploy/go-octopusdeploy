package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteProjectExample provides an example of how to delete a project from
// Octopus Deploy through the Go API client.
func DeleteProjectExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// project values
		projectID string = "project-id"
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

	// delete project
	err = projects.DeleteByID(client, spaceID, projectID)
	if err != nil {
		_ = fmt.Errorf("error deleting project: %v", err)
		return
	}

	fmt.Printf("project deleted: (%s)\n", projectID)
}
