package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteSpaceExample provides an example of how to delete a space from Octopus
// Deploy through the Go API client.
func DeleteSpaceExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// space values
		spaceID string = "space-id"
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

	space, err := client.Spaces.GetByID(spaceID)
	if err != nil {
		_ = fmt.Errorf("error getting space: %v", err)
		return
	}

	// A space cannot be deleted with an active task queue. Attempting to do so
	// will result in an error. In order to delete a space its task queue must
	// be stopped. This can be accomplished by setting the field,
	// TaskQueueStopped to true and updating the space.

	if !space.TaskQueueStopped {
		space.TaskQueueStopped = true
		_, err := client.Spaces.Update(space)
		if err != nil {
			_ = fmt.Errorf("error attempting to stop task queue: %v", err)
			return
		}
	}

	// Delete the space. Note: attempting to delete the last space will result
	// in an error.
	err = client.Spaces.DeleteByID(spaceID)
	if err != nil {
		_ = fmt.Errorf("error deleting space: %v", err)
		return
	}

	fmt.Printf("space deleted: (%s)\n", spaceID)
}
