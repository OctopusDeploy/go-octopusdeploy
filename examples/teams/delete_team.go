package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteTeamExample provides an example of how to delete a team from Octopus
// Deploy through the Go API client.
func DeleteTeamExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// team values
		teamID string = "team-id"
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

	// delete team
	err = client.Teams.DeleteByID(teamID)
	if err != nil {
		_ = fmt.Errorf("error deleting team: %v", err)
		return
	}

	fmt.Printf("team deleted: (%s)\n", teamID)
}
