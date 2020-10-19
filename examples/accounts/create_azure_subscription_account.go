package examples

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	uuid "github.com/google/uuid"
)

func CreateAzureSubscriptionAccountExample() {
	var (
		octopusURL    string = "https://your_octopus_url"
		octopusAPIKey string = "API-YOUR_API_KEY"

		// Azure-specific values
		subscriptionID             uuid.UUID = uuid.MustParse("subscription UUID")
		azureEnvironmentName       string    = "AzureCloud"
		azureManagementEndpoint    string    = "https://management.core.windows.net/"
		azureStorageEndpointSuffix string    = "core.windows.net"

		// Octopus Account values
		spaceName          string = "default"
		accountName        string = "Azure Account"
		accountDescription string = "My Azure Account"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
	}

	azureAccount := model.NewAzureSubscriptionAccount(accountName, subscriptionID)

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
	azureAccount = createdAccount.(*model.AzureSubscriptionAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", azureAccount.ID)
}
