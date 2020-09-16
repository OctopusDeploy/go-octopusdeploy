package examples

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

func CreateCertificateExample() {
	var (
		// Declare working variables
		octopusURL      string = "https://youroctourl"
		octopusAPIKey   string = "API-YOURAPIKEY"
		pfxFilePath     string = "path\\to\\pfxfile.pfx"
		pfxFilePassword string = "PFX-file-password"
		certificateName string = "MyCertificate"
		spaceName       string = "default"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)

	if err != nil {
		// TODO: handle error
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
	octopusCertificate, err := model.NewCertificate(certificateName, certificateData, password)

	if err != nil {
		// TODO: handle error
	}

	_, err = client.Certificates.Add(octopusCertificate)

	if err != nil {
		// TODO: handle error
	}
}
