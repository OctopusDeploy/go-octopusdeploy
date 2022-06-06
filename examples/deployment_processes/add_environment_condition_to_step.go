package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func AddEnvironmentConditionToStepExample() {
	var (
		// Declare working variables
		octopusURL       string   = "https://your_octopus_url"
		apiKey           string   = "API-YOUR_API_KEY"
		spaceID          string   = "space-id"
		projectName      string   = "project-name"
		environmentNames []string = []string{"Development", "Test"}
		stepName         string   = "Run a script"
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

	environmentIDs := []string{}

	for _, environmentName := range environmentNames {
		environments, err := client.Environments.GetByName(environmentName)
		if err != nil {
			_ = fmt.Errorf("error: %w", err)
		}

		environmentIDs = append(environmentIDs, environments[0].GetID())
	}

	projects, err := client.Projects.Get(octopusdeploy.ProjectsQuery{
		Name: projectName,
	})
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}

	// sub-optimal; iterate through collection
	project := projects.Items[0]

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}

	for _, step := range deploymentProcess.Steps {
		if step.Name == stepName {
			for _, action := range step.Actions {
				action.Environments = append(action.Environments, environmentIDs...)
			}
		}
	}

	_, err = client.DeploymentProcesses.Update(deploymentProcess)
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}
}
