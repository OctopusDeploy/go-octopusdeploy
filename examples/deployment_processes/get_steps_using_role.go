package examples

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

func GetStepsUsingRoleExample() {
	var (
		// Declare working variables
		octopusURL string = "https://your_octopus_url"
		apiKey     string = "API-YOUR_API_KEY"
		spaceName  string = "space-id"
		roleName   string = "role-name"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := client.NewClient(nil, apiURL, apiKey, spaceName)
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
			propertyValue := step.Properties["Octopus.Action.TargetRoles"]
			if len(propertyValue.Value) > 0 {
				for _, role := range strings.Split(propertyValue.Value, ",") {
					if strings.EqualFold(role, roleName) {
						fmt.Printf("Step [%s] from project [%s] is using the role [%s]\n", step.Name, project.Name, roleName)
					}
				}
			}
		}
	}
}
