package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/platformhubversioncontrolsettings"
)

// GetPlatformHubVersionControlSettingsExample provides an example of how to read the
// Platform Hub version control settings from Octopus Deploy through the Go API client.
func GetPlatformHubVersionControlSettingsExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	// Platform Hub is system-scoped, so no space ID is required
	octopusClient, err := client.NewClient(nil, apiURL, apiKey, "")
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	settings, err := platformhubversioncontrolsettings.Get(octopusClient)
	if err != nil {
		_ = fmt.Errorf("error getting Platform Hub version control settings: %v", err)
		return
	}

	fmt.Printf("URL: %s\n", settings.URL)
	fmt.Printf("default branch: %s\n", settings.DefaultBranch)
	fmt.Printf("base path: %s\n", settings.BasePath)

	switch creds := settings.Credentials.(type) {
	case *credentials.Anonymous:
		fmt.Println("credentials: anonymous")
	case *credentials.UsernamePassword:
		fmt.Printf("credentials: username/password (%s)\n", creds.Username)
	case *credentials.GitHubApp:
		fmt.Printf("credentials: GitHub App connection (%s)\n", creds.ID)
	case nil:
		fmt.Println("Platform Hub is not configured for version control")
	}
}
