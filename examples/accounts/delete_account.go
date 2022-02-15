package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

// DeleteAccountExample provides an example of how to delete an accountV1 from
// Octopus Deploy through the Go API client.
func DeleteAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// accountV1 values
		accountID string = "accountV1-id"
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

	// delete the accountV1
	err = client.Accounts.DeleteByID(accountID)
	if err != nil {
		_ = fmt.Errorf("error deleting accountV1: %v", err)
		return
	}

	fmt.Printf("accountV1 deleted: (%s)\n", accountID)
}
