package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

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

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
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
	err = client.Users.DeleteByID(userID)
	if err != nil {
		_ = fmt.Errorf("error deleting user: %v", err)
		return
	}

	fmt.Printf("user deleted: (%s)\n", userID)
}
