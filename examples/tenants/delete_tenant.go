package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteTenantExample provides an example of how to delete a tenant from
// Octopus Deploy through the Go API client.
func DeleteTenantExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// tenant values
		tenantID string = "tenant-id"
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

	// delete tenant
	err = client.Tenants.DeleteByID(tenantID)
	if err != nil {
		_ = fmt.Errorf("error deleting tenant: %v", err)
		return
	}

	fmt.Printf("tenant deleted: (%s)\n", tenantID)
}
