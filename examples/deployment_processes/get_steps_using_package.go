package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

func GetStepsUsingPackageExample() {
	var (
		// Declare working variables
		octopusURL string = "https://your_octopus_url"
		apiKey     string = "API-YOUR_API_KEY"
		spaceID    string = "space-id"
		packageID  string = "package-id"
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

	// Get projects
	projects, err := client.Projects.GetAll()
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
		return
	}

	// Loop through list
	for _, project := range projects {
		deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
		if err != nil {
			_ = fmt.Errorf("error: %w", err)
		}

		for _, step := range deploymentProcess.Steps {
			for _, action := range step.Actions {
				for _, pkg := range action.Packages {
					if pkg.ID == packageID {
						fmt.Printf("Step [%s] from project [%s] is using the package [%s]\n", step.Name, project.Name, packageID)
					}
				}
			}
		}
	}
}
