package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteUserRoleExample provides an example of how to delete a user role from
// Octopus Deploy through the Go API client.
func DeleteUserRoleExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// user role values
		userRoleID string = "user-role-id"
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

	// delete user role
	err = client.UserRoles.DeleteByID(userRoleID)
	if err != nil {
		_ = fmt.Errorf("error deleting user role: %v", err)
		return
	}

	fmt.Printf("user role deleted: (%s)\n", userRoleID)
}
