package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/githubconnections"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/platformhubgithubconnections"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/platformhubversioncontrolsettings"
)

// ConfigurePlatformHubVersionControlSettingsWithGitHubConnectionExample provides an example
// of how to point Platform Hub at a repository reachable through a GitHub App connection.
//
// The repository listing supplies both the Git URL and the repository's own default branch,
// so neither needs to be provided.
func ConfigurePlatformHubVersionControlSettingsWithGitHubConnectionExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// version control values
		connectionID   string = "GitHubAppConnections-1"
		repositoryName string = "your-org/your-repo"
		basePath       string = ".octopus/"
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

	repositories, err := platformhubgithubconnections.GetRepositories(octopusClient, connectionID)
	if err != nil {
		_ = fmt.Errorf("error getting repositories for connection: %v", err)
		return
	}

	var repository *githubconnections.Repository
	for _, r := range repositories {
		if r.RepositoryName == repositoryName {
			repository = r
			break
		}
	}
	if repository == nil {
		_ = fmt.Errorf("repository (%s) is not accessible through connection (%s)", repositoryName, connectionID)
		return
	}

	gitCredentials := credentials.NewGitHubApp(connectionID)
	settings := platformhubversioncontrolsettings.NewResource(repository.GitURL, gitCredentials, repository.DefaultBranch, basePath)

	updatedSettings, err := platformhubversioncontrolsettings.Update(octopusClient, settings)
	if err != nil {
		_ = fmt.Errorf("error updating Platform Hub version control settings: %v", err)
		return
	}

	fmt.Printf("Platform Hub configured against: (%s)\n", updatedSettings.URL)
}
