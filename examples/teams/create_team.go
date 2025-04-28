package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/teams"
	"net/url"
)

func CreateTeamExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		name string = "team-name"
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

	// create team
	team := teams.NewTeam(name)

	// create team through Add(); returns error if fails
	createdTeam, err := teams.Add(client, team)
	if err != nil {
		_ = fmt.Errorf("error creating team: %v", err)
		return
	}

	fmt.Printf("team created: (%s)\n", createdTeam.GetID())
}
