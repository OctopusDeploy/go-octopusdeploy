package examples

import "github.com/OctopusDeploy/go-octopusdeploy/client"

func AddEnvironmentConditionToStepExample() {
	var (
		// Declare working variables
		octopusURL       string   = "https://youroctourl"
		octopusAPIKey    string   = "API-YOURAPIKEY"
		spaceName        string   = "default"
		projectName      string   = "MyProject"
		environmentNames []string = []string{"Development", "Test"}
		stepName         string   = "Run a script"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)

	if err != nil {
		// TODO: handle error
	}

	environmentIDs := []string{}

	for _, environmentName := range environmentNames {
		environments, err := client.Environments.GetByName(environmentName)

		if err != nil {
			// TODO: handle error
		}

		environmentIDs = append(environmentIDs, environments[0].ID)
	}

	project, err := client.Projects.GetByName(projectName)

	if err != nil {
		// TODO: handle error
	}

	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)

	if err != nil {
		// TODO: handle error
	}

	for _, step := range deploymentProcess.Steps {
		if step.Name == stepName {
			for _, action := range step.Actions {
				for _, environmentID := range environmentIDs {
					action.Environments = append(action.Environments, environmentID)
				}
			}
		}
	}

	_, err = client.DeploymentProcesses.Update(*deploymentProcess)

	if err != nil {
		// TODO: handle error
	}
}
