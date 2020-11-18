package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteBuildInformationExample provides an example of how to delete build
// information from Octopus Deploy through the Go API client.
func DeleteBuildInformationExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// build information values
		buildInformationID string = "build-information-id"
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

	// delete build information
	err = client.BuildInformation.DeleteByID(buildInformationID)
	if err != nil {
		_ = fmt.Errorf("error deleting build information: %v", err)
		return
	}

	fmt.Printf("build information deleted: (%s)\n", buildInformationID)
}
