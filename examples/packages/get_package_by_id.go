package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

func GetPackageByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// package values
		packageID string = "package-id"
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

	// get package by its ID
	octopusPackage, err := client.Packages.GetByID(packageID)
	if err != nil {
		_ = fmt.Errorf("error getting package: %v", err)
		return
	}

	fmt.Printf("package: (%s)\n", octopusPackage.GetID())
}
