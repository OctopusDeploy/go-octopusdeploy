package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/processtemplates"
)

// ListProcessTemplatesExample provides an example of how to list the process templates on a
// Git reference in Platform Hub through the Go API client.
func ListProcessTemplatesExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"

		// process template values
		gitRef string = "refs/heads/main"
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

	query := processtemplates.ProcessTemplatesQuery{
		GitRef: gitRef,
		Take:   30,
	}

	results, err := processtemplates.List(octopusClient, query)
	if err != nil {
		_ = fmt.Errorf("error listing process templates: %v", err)
		return
	}

	for _, processTemplate := range results.ProcessTemplates {
		fmt.Printf("process template: (%s) %s, %d step(s), %d parameter(s)\n", processTemplate.Slug, processTemplate.Name, len(processTemplate.Steps), len(processTemplate.Parameters))
	}

	fmt.Printf("showing %d of %d process template(s)\n", len(results.ProcessTemplates), results.TotalResults)
}
