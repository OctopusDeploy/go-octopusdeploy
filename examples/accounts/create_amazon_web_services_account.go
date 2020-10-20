package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateAmazonWebServicesAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// AWS-specific values
		accessKey string                       = "access-key"
		secretKey octopusdeploy.SensitiveValue = octopusdeploy.NewSensitiveValue("secret-key")

		// account values
		accountName        string = "AWS Account"
		accountDescription string = "My AWS Account"
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

	awsAccount := octopusdeploy.NewAmazonWebServicesAccount(accountName, accessKey, secretKey)

	// fill in account details
	awsAccount.Description = accountDescription

	// create account
	createdAccount, err := client.Accounts.Add(awsAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access AWS-specific fields
	awsAccount = createdAccount.(*octopusdeploy.AmazonWebServicesAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", awsAccount.ID)
}
