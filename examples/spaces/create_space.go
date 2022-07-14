package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/spaces"
)

// CreateSpaceExample provides an example of how to create a space in Octopus
// Deploy through the Go API client.
func CreateSpaceExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// space values
		name   string = "space-name"
		userID string = "user-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := client.NewClient(nil, apiURL, apiKey, "")
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// create space
	space := spaces.NewSpace(name)
	space.SpaceManagersTeamMembers = []string{userID}
	createdSpace, err := client.Spaces.Add(space)
	if err != nil {
		_ = fmt.Errorf("error creating space: %v", err)
		return
	}

	fmt.Printf("space: (%s)\n", createdSpace.GetID())
}
