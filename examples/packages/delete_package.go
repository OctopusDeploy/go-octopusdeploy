package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeletePackageExample provides an example of how to delete a package from
// Octopus Deploy through the Go API client.
func DeletePackageExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// package values
		packageID string = "package-id"
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

	// delete package
	err = client.Packages.DeleteByID(packageID)
	if err != nil {
		_ = fmt.Errorf("error deleting package: %v", err)
		return
	}

	fmt.Printf("package deleted: (%s)\n", packageID)
}
