package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
)

func DeleteLibraryVariableSetExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// library variable set values
		libraryVariableSetID string = "library-variable-set-id"
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

	// delete library variable set
	err = client.LibraryVariableSets.DeleteByID(libraryVariableSetID)
	if err != nil {
		_ = fmt.Errorf("error deleting library variable set: %v", err)
		return
	}

	fmt.Printf("library variable set deleted: (%s)\n", libraryVariableSetID)
}
