package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
)

func CreateAmazonWebServicesAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// AWS-specific values
		accessKey string                   = "access-key"
		secretKey *services.SensitiveValue = services.NewSensitiveValue("secret-key")

		// account values
		accountName        string = "AWS Account"
		accountDescription string = "My AWS Account"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	octoEndpoint, err := services.NewOctopusServerEndpoint(apiURL, apiKey)
	client, err := services.NewSpaceScopedClient(octoEndpoint, spaceID, nil)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	awsAccount, err := accounts.NewAmazonWebServicesAccount(accountName, accessKey, secretKey)
	if err != nil {
		_ = fmt.Errorf("error creating AWS account: %v", err)
		return
	}

	// fill in account details
	awsAccount.Description = accountDescription

	// create account
	createdAccount, err := client.Accounts.Add(awsAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access AWS-specific fields
	awsAccount = createdAccount.(*accounts.AmazonWebServicesAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", awsAccount.GetID())
}
