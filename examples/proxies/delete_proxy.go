package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteProxyExample provides an example of how to delete a proxy from
// Octopus Deploy through the Go API client.
func DeleteProxyExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// proxy values
		proxyID string = "proxy-id"
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

	// delete proxy
	err = client.Proxies.DeleteByID(proxyID)
	if err != nil {
		_ = fmt.Errorf("error deleting proxy: %v", err)
		return
	}

	fmt.Printf("proxy deleted: (%s)\n", proxyID)
}
