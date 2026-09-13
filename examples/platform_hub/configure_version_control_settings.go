package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/platformhubversioncontrolsettings"
)

// ConfigurePlatformHubVersionControlSettingsExample provides an example of how to point
// Platform Hub at a Git repository using username/password credentials.
//
// If the repository is already configured, changing the URL will repoint Platform Hub at the new repository.
func ConfigurePlatformHubVersionControlSettingsExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// version control values
		gitURL        string = "https://github.com/your-org/your-repo.git"
		gitUsername   string = "your-username"
		gitPassword   string = "your-personal-access-token"
		defaultBranch string = "main"
		basePath      string = ".octopus/"
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

	gitCredentials := credentials.NewUsernamePassword(gitUsername, core.NewSensitiveValue(gitPassword))
	settings := platformhubversioncontrolsettings.NewResource(gitURL, gitCredentials, defaultBranch, basePath)

	updatedSettings, err := platformhubversioncontrolsettings.Update(octopusClient, settings)
	if err != nil {
		_ = fmt.Errorf("error updating Platform Hub version control settings: %v", err)
		return
	}

	fmt.Printf("Platform Hub configured against: (%s)\n", updatedSettings.URL)
}
