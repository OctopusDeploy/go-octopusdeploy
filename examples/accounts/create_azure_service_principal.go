package examples

import (
	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	uuid "github.com/google/uuid"
)

func CreateAzureServicePrincipalExample() {
	var (
		octopusURL    string = "https://youroctourl"
		octopusAPIKey string = "API-YOURAPIKEY"

		// Azure specific details
		azureApplicationID  uuid.UUID            = uuid.MustParse("Client UUID")
		azureSecret         model.SensitiveValue = model.NewSensitiveValue("Secret")
		azureSubscriptionID uuid.UUID            = uuid.MustParse("Subscription UUID")
		azureTenantID       uuid.UUID            = uuid.MustParse("Tenant UUID")

		// Octopus Account details
		spaceName           string   = "default"
		accountName         string   = "Azure Account"
		accountDescription  string   = "My Azure Account"
		tenantParticipation string   = "Untenanted"
		tenantTags          []string = nil
		tenantIDs           []string = nil
		environmentIDs      []string = nil
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)

	if err != nil {
		// TODO: handle error
	}

	azureAccount := model.NewAzureServicePrincipalAccount(accountName, azureSubscriptionID, azureTenantID, azureApplicationID, azureSecret)

	if err != nil {
		// TODO: handle error
	}

	// Fill in account details
	azureAccount.Description = accountDescription
	azureAccount.TenantedDeploymentMode = tenantParticipation
	azureAccount.TenantTags = tenantTags
	azureAccount.TenantIDs = tenantIDs
	azureAccount.EnvironmentIDs = environmentIDs

	// Create account
	_, err = client.Accounts.Add(azureAccount)

	if err != nil {
		// TODO: handle error
	}
}
