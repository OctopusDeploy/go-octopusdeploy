package examples

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/certificates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// ReplaceCertificateExample provides an example of how to replace an existing
// certificate in Octopus Deploy through the Go API client.
func ReplaceCertificateExample() {
	var (
		apiKey          string = "API-YOUR_API_KEY"
		certificateName string = "MyCertificate"
		octopusURL      string = "https://your_octopus_url"
		pfxFilePath     string = "path\\to\\pfxfile.pfx"
		pfxFilePassword string = "PFX-file-password"
		spaceID         string = "default"
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

	// construct query
	query := certificates.CertificatesQuery{
		PartialName: certificateName,
	}

	// find the certificate with a specific name
	certificateResources, err := certificates.Get(client, spaceID, query)
	if err != nil {
		_ = fmt.Errorf("error matching certificate(s): %v", err)
		return
	}

	// NOTE: this is lazy and should be replaced by something more robust
	certificate := certificateResources.Items[0]

	file, err := os.Open(pfxFilePath)
	if err != nil {
		_ = fmt.Errorf("error opening file path: %v", err)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		_ = fmt.Errorf("error reading file: %v", err)
		return
	}

	// encode to Base64
	base64Certificate := base64.StdEncoding.EncodeToString(data)

	// replace certificate
	replacementCertificate := certificates.NewReplacementCertificate(base64Certificate, pfxFilePassword)
	if _, err = client.Certificates.Replace(certificate.GetID(), replacementCertificate); err != nil {
		_ = fmt.Errorf("error replacing certificate: %v", err)
	}
}
