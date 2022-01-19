package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateProjectExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		name string = "project-name"
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

	// NOTE: a lifecycle is obtained through the Lifecycles service API
	//
	// lifecycles, err = client.Lifecycles.GetAll()
	// lifecycle, err = client.Lifecycles.GetByID(id)
	// lifecycles, err = client.Lifecycles.GetByPartialName(name)
	//
	// the lifecycle ID value (below) is obtained via GetID()

	lifecycleID := "lifecycle-id"

	// NOTE: a project group is obtained through the ProjectGroups service API
	//
	// projectGroups, err = client.ProjectGroups.GetAll()
	// projectGroup, err = client.ProjectGroups.GetByID(id)
	// projectGroups, err = client.ProjectGroups.GetByPartialName(name)
	//
	// the project group ID value (below) is obtained via GetID()

	projectGroupID := "project-group-id"

	// create project
	project := octopusdeploy.NewProject(spaceID, name, lifecycleID, projectGroupID)

	// update any additional project fields here...

	// create project through Add(); returns error if fails
	createdProject, err := client.Projects.Add(project)
	if err != nil {
		_ = fmt.Errorf("error creating project: %v", err)
		return
	}

	fmt.Printf("project created: (%s)\n", createdProject.GetID())
}
