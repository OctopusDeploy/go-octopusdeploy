package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/platformhubgithubconnections"
)

// GetPlatformHubGitHubConnectionByIDExample provides an example of how to get a single
// Platform Hub GitHub App connection and the repositories it grants access to.
func GetPlatformHubGitHubConnectionByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// GitHub connection values
		connectionID string = "GitHubAppConnections-1"
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

	connection, err := platformhubgithubconnections.GetByID(octopusClient, connectionID)
	if err != nil {
		_ = fmt.Errorf("error getting GitHub connection: %v", err)
		return
	}

	fmt.Printf("connection: (%s) %s\n", connection.ID, connection.Installation.AccountLogin)
	fmt.Printf("status: %s %s\n", connection.Status, connection.StatusUserMessage)

	for _, repository := range connection.Repositories {
		fmt.Printf("  repository: %s (%s)\n", repository.RepositoryName, repository.GitURL)
	}

	// repositories configured on the connection that GitHub no longer returns; they may have
	// been deleted, renamed, or had access revoked
	for _, repository := range connection.UnknownRepositories {
		fmt.Printf("  unknown repository: %s (%s)\n", repository.RepositoryName, repository.RepositoryID)
	}
}
