package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/certificates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

func GetCertificateByIDExample() {
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

	// get certificate by its ID
	certificate, err := certificates.GetByID(client, spaceID, certificateID)
	if err != nil {
		_ = fmt.Errorf("error getting certificate: %v", err)
		return
	}

	fmt.Printf("certificate: (%s)\n", certificate.GetID())
}
