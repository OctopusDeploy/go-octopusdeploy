package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/githubconnections"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/platformhubgithubconnections"
)

// ListPlatformHubGitHubConnectionsExample provides an example of how to list the GitHub App
// connections available to Platform Hub, paging through the results.
func ListPlatformHubGitHubConnectionsExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// paging values; both skip and take are required by the API
		take int = 30
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

	settings, err := githubconnections.GetSettings(octopusClient)
	if err != nil {
		_ = fmt.Errorf("error getting GitHub App settings: %v", err)
		return
	}
	if !settings.CanUseGitHubApp {
		fmt.Println("this Octopus instance cannot use GitHub App connections")
		return
	}

	var connections []*githubconnections.Connection
	for {
		page, err := platformhubgithubconnections.List(octopusClient, len(connections), take)
		if err != nil {
			_ = fmt.Errorf("error listing GitHub connections: %v", err)
			return
		}

		connections = append(connections, page.Connections...)

		if len(page.Connections) == 0 || len(connections) >= page.TotalResults {
			break
		}
	}

	for _, connection := range connections {
		fmt.Printf("connection: (%s) %s %s [%s]\n", connection.ID, connection.Installation.AccountType, connection.Installation.AccountLogin, connection.Status)

		if connection.Status != githubconnections.ConnectionStatusConnected {
			continue
		}

		repositories, err := platformhubgithubconnections.GetRepositories(octopusClient, connection.ID)
		if err != nil {
			_ = fmt.Errorf("error getting repositories for connection: %v", err)
			return
		}

		for _, repository := range repositories {
			fmt.Printf("  repository: %s (%s), default branch: %s\n", repository.RepositoryName, repository.GitURL, repository.DefaultBranch)
		}
	}
}
