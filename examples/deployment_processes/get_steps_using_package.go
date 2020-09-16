package examples

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
)

func GetStepsUsingPackageExample() {
	var (
		// Declare working variables
		octopusURL    string = "https://youroctourl"
		octopusAPIKey string = "API-YOURAPIKEY"
		spaceName     string = "default"
		packageID     string = "PackageId"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)

	if err != nil {
		// TODO: handle error
	}

	// Get projects
	projects, err := client.Projects.GetAll()

	if err != nil {
		// TODO: handle error
	}

	// Loop through list
	for _, project := range projects {
		deploymentProcess, err := client.DeploymentProcesses.Get(project.DeploymentProcessID)

		if err != nil {
			// TODO: handle error
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
