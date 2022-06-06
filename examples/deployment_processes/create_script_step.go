package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateScriptStepExample() {
	var (
		// Declare working variables
		octopusURL  string = "https://your_octopus_url"
		apiKey      string = "API-YOUR_API_KEY"
		projectName string = "project-name"
		roleName    string = "role-name"
		scriptBody  string = "Write-Host \"Hello world\""
		spaceID     string = "space-id"
		stepName    string = "Run a script"
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

	// Get project
	projects, err := client.Projects.Get(octopusdeploy.ProjectsQuery{
		Name: projectName,
	})

	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}

	// sub-optimal; iterate through collection
	project := *projects.Items[0]

	// Get the deployment process
	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}

	// Create new step object
	newStep := octopusdeploy.NewDeploymentStep(stepName)
	newStep.Condition = "Success"
	newStep.Properties["Octopus.Action.TargetRoles"] = octopusdeploy.NewPropertyValue(roleName, false)

	// Create new script action
	stepAction := octopusdeploy.NewDeploymentAction(stepName, "Octopus.Script")
	stepAction.Properties["Octopus.Action.Script.ScriptBody"] = octopusdeploy.NewPropertyValue(scriptBody, false)

	// Add step action and step to process
	newStep.Actions = append(newStep.Actions, *stepAction)
	deploymentProcess.Steps = append(deploymentProcess.Steps, *newStep)

	// Update process
	_, err = client.DeploymentProcesses.Update(deploymentProcess)
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
	}
}
