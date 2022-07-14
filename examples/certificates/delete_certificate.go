package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteCertificateExample provides an example of how to delete a certificate
// from Octopus Deploy through the Go API client.
func DeleteCertificateExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// certificate values
		certificateID string = "certificate-id"
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

	// delete certificate
	err = client.Certificates.DeleteByID(certificateID)
	if err != nil {
		_ = fmt.Errorf("error deleting certificate: %v", err)
		return
	}

	fmt.Printf("certificate deleted: (%s)\n", certificateID)
}
