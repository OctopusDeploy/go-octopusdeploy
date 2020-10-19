package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	uuid "github.com/google/uuid"
)

func CreateAzureServicePrincipalExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// Azure-specific values
		azureApplicationID  uuid.UUID            = uuid.MustParse("client-UUID")
		azureSecret         model.SensitiveValue = model.NewSensitiveValue("azure-secret")
		azureSubscriptionID uuid.UUID            = uuid.MustParse("subscription-UUID")
		azureTenantID       uuid.UUID            = uuid.MustParse("tenant-UUID")

		// account values
		accountName        string   = "Azure Account"
		accountDescription string   = "My Azure Account"
		tenantTags         []string = nil
		tenantIDs          []string = nil
		environmentIDs     []string = nil
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

	azureAccount := model.NewAzureServicePrincipalAccount(accountName, azureSubscriptionID, azureTenantID, azureApplicationID, azureSecret)

	// fill in account details
	azureAccount.Description = accountDescription
	azureAccount.TenantTags = tenantTags
	azureAccount.TenantIDs = tenantIDs
	azureAccount.EnvironmentIDs = environmentIDs

	// create account
	createdAccount, err := client.Accounts.Add(azureAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access Username/Password-specific fields
	azureAccount = createdAccount.(*model.AzureServicePrincipalAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", azureAccount.ID)
}
