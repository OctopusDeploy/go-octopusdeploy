package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/processtemplates"
)

// GetProcessTemplateBySlugExample provides an example of how to get a single process template
// from Platform Hub through the Go API client.
func GetProcessTemplateBySlugExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// process template values
		gitRef string = "refs/heads/main"
		slug   string = "your-process-template-slug"
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

	processTemplate, err := processtemplates.GetBySlug(octopusClient, gitRef, slug)
	if err != nil {
		_ = fmt.Errorf("error getting process template: %v", err)
		return
	}

	fmt.Printf("process template: (%s) %s\n", processTemplate.Slug, processTemplate.Name)
	fmt.Printf("description: %s\n", processTemplate.Description)

	for _, step := range processTemplate.Steps {
		fmt.Printf("  step: %s\n", step.Name)
	}

	for _, parameter := range processTemplate.Parameters {
		fmt.Printf("  parameter: %s, optional: %t\n", parameter.Name, parameter.IsOptional)

		for _, value := range parameter.Values {
			// a parameter value is sensitive when it is backed by a sensitive value rather
			// than a plain string
			if value.Value.IsSensitive {
				fmt.Println("    default value: (sensitive)")
				continue
			}
			fmt.Printf("    default value: %s\n", value.Value.Value)
		}
	}
}
