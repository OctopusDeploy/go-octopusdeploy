package examples

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

// CreateCertificateExample shows how to create a certificate using go-octopusdeploy.
func CreateCertificateExample() {
	var (
		apiKey          string = "API-YOUR_API_KEY"
		certificateName string = "MyCertificate"
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
		// TODO: handle error
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		// TODO: handle error
	}

	// Convert file to base64
	base64Certificate := base64.StdEncoding.EncodeToString(data)

	// Create certificate object
	certificateData := model.NewSensitiveValue(base64Certificate)
	password := model.NewSensitiveValue(pfxFilePassword)
	octopusCertificate := model.NewCertificate(certificateName, certificateData, password)

	_, err = client.Certificates.Add(octopusCertificate)

	if err != nil {
		// TODO: handle error
	}
}
