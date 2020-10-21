package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	uuid "github.com/google/uuid"
)

func CreateAzureSubscriptionAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// Azure-specific values
		accountDescription         string    = "My Azure Account"
		azureEnvironmentName       string    = "AzureCloud"
		azureManagementEndpoint    string    = "https://management.core.windows.net/"
		azureStorageEndpointSuffix string    = "core.windows.net"
		name                       string    = "Azure Account"
		subscriptionID             uuid.UUID = uuid.MustParse("subscription-UUID")
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

	azureAccount := octopusdeploy.NewAzureSubscriptionAccount(name, subscriptionID)

	// fill in account details
	azureAccount.Description = accountDescription
	azureAccount.AzureEnvironment = azureEnvironmentName
	azureAccount.ManagementEndpoint = azureManagementEndpoint
	azureAccount.StorageEndpointSuffix = azureStorageEndpointSuffix

	// create account
	createdAccount, err := client.Accounts.Add(azureAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access Azure-specific fields
	azureAccount = createdAccount.(*octopusdeploy.AzureSubscriptionAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", azureAccount.GetID())
}
