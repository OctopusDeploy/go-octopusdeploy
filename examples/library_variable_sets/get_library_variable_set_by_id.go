package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func GetLibraryVariableSetByIDExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// library variable set values
		libraryVariableSetID string = "libraryVariableSet-id"
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

	// get library variable set by its ID
	libraryVariableSet, err := client.LibraryVariableSets.GetByID(libraryVariableSetID)
	if err != nil {
		_ = fmt.Errorf("error getting library variable set: %v", err)
		return
	}

	fmt.Printf("library variable set: (%s)\n", libraryVariableSet.GetID())
}
