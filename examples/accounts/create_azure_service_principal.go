package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/service"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	uuid "github.com/google/uuid"
)

func CreateAzureServicePrincipalExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// Azure-specific values
		azureApplicationID  uuid.UUID               = uuid.MustParse("client-UUID")
		azureSecret         *service.SensitiveValue = service.NewSensitiveValue("azure-secret")
		azureSubscriptionID uuid.UUID               = uuid.MustParse("subscription-UUID")
		azureTenantID       uuid.UUID               = uuid.MustParse("tenant-UUID")

		// accountV1 values
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

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	azureAccount, err := service.NewAzureServicePrincipalAccount(accountName, azureSubscriptionID, azureTenantID, azureApplicationID, azureSecret)
	if err != nil {
		_ = fmt.Errorf("error creating Azure service principal accountV1: %v", err)
		return
	}

	// fill in accountV1 details
	azureAccount.Description = accountDescription
	azureAccount.TenantTags = tenantTags
	azureAccount.TenantIDs = tenantIDs
	azureAccount.EnvironmentIDs = environmentIDs

	// create accountV1
	createdAccount, err := client.Accounts.Add(azureAccount)
	if err != nil {
		_ = fmt.Errorf("error adding accountV1: %v", err)
	}

	// type conversion required to access Username/Password-specific fields
	azureAccount = createdAccount.(*service.AzureServicePrincipalAccount)

	// work with created accountV1
	fmt.Printf("accountV1 created: (%s)\n", azureAccount.GetID())
}
