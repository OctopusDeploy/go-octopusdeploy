package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

// DeleteScopedUserRoleExample provides an example of how to delete a scoped
// user role from Octopus Deploy through the Go API client.
func DeleteScopedUserRoleExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// scoped user role values
		scopedUserRoleID string = "scoped-user-role-id"
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

	// delete scoped user role
	err = client.ScopedUserRoles.DeleteByID(scopedUserRoleID)
	if err != nil {
		_ = fmt.Errorf("error deleting scoped user role: %v", err)
		return
	}

	fmt.Printf("scoped user role deleted: (%s)\n", scopedUserRoleID)
}
