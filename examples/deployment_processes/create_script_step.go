package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/projects"
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

	client, err := client.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// Get project
	projects, err := client.Projects.Get(projects.ProjectsQuery{Name: projectName})
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
		return
	}

	// sub-optimal; iterate through collection
	project := *projects.Items[0]

	// Get the deployment process
	deploymentProcess, err := client.DeploymentProcesses.GetByID(project.DeploymentProcessID)
	if err != nil {
		_ = fmt.Errorf("error: %w", err)
		return
	}

	// Create new step object
	newStep := deployments.NewDeploymentStep(stepName)
	newStep.Properties["Octopus.Action.TargetRoles"] = core.NewPropertyValue(roleName, false)

	// Create new script action
	stepAction := deployments.NewDeploymentAction(stepName, "Octopus.Script")
	stepAction.Properties["Octopus.Action.Script.ScriptBody"] = core.NewPropertyValue(scriptBody, false)

	// Add step action and step to process
	newStep.Actions = append(newStep.Actions, stepAction)
	deploymentProcess.Steps = append(deploymentProcess.Steps, newStep)

	// Update process
	if _, err = client.DeploymentProcesses.Update(deploymentProcess); err != nil {
		_ = fmt.Errorf("error: %w", err)
	}
}
