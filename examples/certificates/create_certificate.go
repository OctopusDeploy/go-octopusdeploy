package examples

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/certificates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
)

// CreateCertificateExample shows how to create a certificate using go-octopusdeploy.
func CreateCertificateExample() {
	var (
		apiKey          string = "API-YOUR_API_KEY"
		certificateName string = "certificate-name"
		octopusURL      string = "https://your_octopus_url"
		pfxFilePath     string = "path\\to\\pfxfile.pfx"
		pfxFilePassword string = "PFX-file-password"
		spaceID         string = "space-id"
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

	file, err := os.Open(pfxFilePath)
	if err != nil {
		_ = fmt.Errorf("error opening private key: %v", err)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		_ = fmt.Errorf("error reading file: %v", err)
		return
	}

	// certificate properties
	base64Certificate := base64.StdEncoding.EncodeToString(data)
	certificateData := core.NewSensitiveValue(base64Certificate)
	password := core.NewSensitiveValue(pfxFilePassword)

	// create certificate
	certificate := certificates.NewCertificateResource(certificateName, certificateData, password)

	// create certificate through Add(); returns error if fails
	createdCertificate, err := certificates.Add(client, certificate)
	if err != nil {
		_ = fmt.Errorf("error adding certificate: %v", err)
	}

	// work with created certificate
	fmt.Printf("certificate created: (%s)\n", createdCertificate.GetID())
}
