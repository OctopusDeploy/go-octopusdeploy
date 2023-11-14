package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	uuid "github.com/google/uuid"
)

func CreateAzureOIDCExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// Azure-specific values
		azureApplicationID  uuid.UUID = uuid.MustParse("client-UUID")
		azureSubscriptionID uuid.UUID = uuid.MustParse("subscription-UUID")
		azureTenantID       uuid.UUID = uuid.MustParse("tenant-UUID")

		// Subject claims
		deploymentSubjectKeys  []string = nil
		healthCheckSubjectKeys []string = nil
		accountTestSubjectKeys []string = nil

		// Other claims
		audience string = ""

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

	azureAccount, err := accounts.NewAzureOIDCAccount(accountName, azureSubscriptionID, azureTenantID, azureApplicationID)
	if err != nil {
		_ = fmt.Errorf("error creating Azure service principal account: %v", err)
		return
	}

	// fill in claims
	azureAccount.DeploymentSubjectKeys = deploymentSubjectKeys
	azureAccount.HealthCheckSubjectKeys = healthCheckSubjectKeys
	azureAccount.AccountTestSubjectKeys = accountTestSubjectKeys
	azureAccount.Audience = audience

	// fill in account details
	azureAccount.Description = accountDescription
	azureAccount.TenantTags = tenantTags
	azureAccount.TenantIDs = tenantIDs
	azureAccount.EnvironmentIDs = environmentIDs

	// create account
	createdAccount, err := accounts.Add(client, azureAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access Username/Password-specific fields
	azureAccount = createdAccount.(*accounts.AzureOIDCAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", azureAccount.GetID())
}
