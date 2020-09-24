package examples

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

// ReplaceCertificateExample shows how to replace an existing certificate using go-octopusdeploy.
func ReplaceCertificateExample() {
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

	// get certificates
	certificateList, err := client.Certificates.GetByPartialName(certificateName)
	if err != nil {
		// TODO: handle error
	}

	// find the certificate with a specific name
	certificate := certificateList[0]

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

	// Replace certificate
	replacementCertificate := model.NewReplacementCertificate(base64Certificate, pfxFilePassword)
	_, err = client.Certificates.Replace(certificate.ID, replacementCertificate)
	if err != nil {
		// TODO: handle error
	}
}
