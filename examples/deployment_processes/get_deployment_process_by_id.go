package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func GetDeploymentProcessByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// deployment process values
		deploymentProcessID string = "deployment-process-id"
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

	// get deployment process by its ID
	deploymentProcess, err := client.DeploymentProcesses.GetByID(deploymentProcessID)
	if err != nil {
		_ = fmt.Errorf("error getting deployment process: %v", err)
		return
	}

	fmt.Printf("deployment process: (%s)\n", deploymentProcess.GetID())
}
