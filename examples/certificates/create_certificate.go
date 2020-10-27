package examples

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
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

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	file, err := os.Open(pfxFilePath)
	if err != nil {
		_ = fmt.Errorf("error opening private key: %v", err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		_ = fmt.Errorf("error reading file: %v", err)
		return
	}

	// certificate properties
	base64Certificate := base64.StdEncoding.EncodeToString(data)
	certificateData := octopusdeploy.NewSensitiveValue(base64Certificate)
	password := octopusdeploy.NewSensitiveValue(pfxFilePassword)

	// create certificate
	certificate := octopusdeploy.NewCertificateResource(certificateName, certificateData, password)

	// create certificate through Add(); returns error if fails
	createdCertificate, err := client.Certificates.Add(certificate)
	if err != nil {
		_ = fmt.Errorf("error adding certificate: %v", err)
	}

	// work with created certificate
	fmt.Printf("certificate created: (%s)\n", createdCertificate.GetID())
}
