package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

// DeleteAccountExample provides an example of how to delete an account from
// Octopus Deploy through the Go API client.
func DeleteAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// account values
		accountID string = "account-id"
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

	// delete the account, empty spaceID will revert to using the spaceID provided on the client
	err = accounts.DeleteByID(client, "", accountID)
	if err != nil {
		_ = fmt.Errorf("error deleting account: %v", err)
		return
	}

	fmt.Printf("account deleted: (%s)\n", accountID)
}
