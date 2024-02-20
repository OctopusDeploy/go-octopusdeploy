package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/users"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteUserExample provides an example of how to delete a user from Octopus
// Deploy through the Go API client.
func DeleteUserExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// user values
		userID string = "user-id"
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

	// get the current user
	user, err := client.Users.GetMe()
	if err != nil {
		_ = fmt.Errorf("error getting user: %v", err)
		return
	}

	// A user attempting to delete itself will result in an error. The field,
	// IsRequestor may be checked to see if this situation exists.
	if user.IsRequestor {
		return
	}

	// delete user
	err = users.DeleteByID(client, userID)
	if err != nil {
		_ = fmt.Errorf("error deleting user: %v", err)
		return
	}

	fmt.Printf("user deleted: (%s)\n", userID)
}
