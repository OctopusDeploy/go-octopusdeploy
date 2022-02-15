package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
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

	azureAccount, err := accounts.NewAzureSubscriptionAccount(name, subscriptionID)
	if err != nil {
		_ = fmt.Errorf("error creating Azure subscription accountV1: %v", err)
		return
	}

	// fill in accountV1 details
	azureAccount.Description = accountDescription
	azureAccount.AzureEnvironment = azureEnvironmentName
	azureAccount.ManagementEndpoint = azureManagementEndpoint
	azureAccount.StorageEndpointSuffix = azureStorageEndpointSuffix

	// create accountV1
	createdAccount, err := client.Accounts.Add(azureAccount)
	if err != nil {
		_ = fmt.Errorf("error adding accountV1: %v", err)
	}

	// type conversion required to access Azure-specific fields
	azureAccount = createdAccount.(*accounts.AzureSubscriptionAccount)

	// work with created accountV1
	fmt.Printf("accountV1 created: (%s)\n", azureAccount.GetID())
}
