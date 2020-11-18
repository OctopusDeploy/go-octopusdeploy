package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteTagSetExample provides an example of how to delete a tag set from
// Octopus Deploy through the Go API client.
func DeleteTagSetExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// tag set values
		tagSetID string = "tagSet-id"
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

	// delete tag set
	err = client.TagSets.DeleteByID(tagSetID)
	if err != nil {
		_ = fmt.Errorf("error deleting tag set: %v", err)
		return
	}

	fmt.Printf("tag set deleted: (%s)\n", tagSetID)
}
