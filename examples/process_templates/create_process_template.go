package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/processtemplates"
)

// CreateProcessTemplateExample provides an example of how to create a process template in
// Platform Hub through the Go API client.
//
// Steps and parameters cannot be set at creation time; the new template is empty, and should be
// configured after in Git or Octopus Deploy.
func CreateProcessTemplateExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// process template values
		gitRef            string = "refs/heads/main"
		name              string = "your-process-template-name"
		description       string = "your-process-template-description"
		changeDescription string = "Add a process template"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	// Platform Hub is system-scoped, so no space ID is required
	octopusClient, err := client.NewClient(nil, apiURL, apiKey, "")
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	processTemplate, err := processtemplates.Add(octopusClient, gitRef, name, description, changeDescription)
	if err != nil {
		_ = fmt.Errorf("error creating process template: %v", err)
		return
	}

	fmt.Printf("process template created: (%s) %s\n", processTemplate.Slug, processTemplate.Name)
}
