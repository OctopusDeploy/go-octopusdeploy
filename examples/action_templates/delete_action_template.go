package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteActionTemplateExample provides an example of how to delete an action
// template from Octopus Deploy through the Go API client.
func DeleteActionTemplateExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// action template values
		actionTemplateID string = "action-template-id"
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

	// delete action template
	err = client.ActionTemplates.DeleteByID(actionTemplateID)
	if err != nil {
		_ = fmt.Errorf("error deleting action template: %v", err)
		return
	}

	fmt.Printf("action template deleted: (%s)\n", actionTemplateID)
}
