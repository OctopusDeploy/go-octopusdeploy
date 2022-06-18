package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

// DeleteProjectGroupExample provides an example of how to delete a project
// group from Octopus Deploy through the Go API client.
func DeleteProjectGroupExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// project group values
		projectGroupID string = "project-group-id"
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

	// delete project group
	err = client.ProjectGroups.DeleteByID(projectGroupID)
	if err != nil {
		_ = fmt.Errorf("error deleting project group: %v", err)
		return
	}

	fmt.Printf("project group deleted: (%s)\n", projectGroupID)
}
