package examples

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
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

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// Get projects
	projects, err := client.Projects.GetAll()
	if err != nil {
		// TODO: handle error
		return
	}

	// Loop through list
	for _, project := range projects {
		deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)

		if err != nil {
			// TODO: handle error
		}

		for _, step := range deploymentProcess.Steps {
			propertyValue := step.Properties["Octopus.Action.TargetRoles"]
			if len(propertyValue) > 0 {
				for _, role := range strings.Split(propertyValue, ",") {
					if strings.ToLower(role) == strings.ToLower(roleName) {
						fmt.Printf("Step [%s] from project [%s] is using the role [%s]\n", step.Name, project.Name, roleName)
					}
				}
			}
		}
	}
}
