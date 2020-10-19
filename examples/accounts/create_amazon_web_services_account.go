package examples

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

func CreateAmazonWebServicesAccountExample() {
	var (
		octopusURL    string = "https://your_octopus_url"
		octopusAPIKey string = "API-YOUR_API_KEY"

		// AWS-specific values
		accessKey string               = "access-key"
		secretKey model.SensitiveValue = model.NewSensitiveValue("secret-key")

		// Octopus Account values
		spaceName          string = "default"
		accountName        string = "AWS Account"
		accountDescription string = "My AWS Account"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
	}

	awsAccount := model.NewAmazonWebServicesAccount(accountName, accessKey, secretKey)

	// fill in account details
	awsAccount.Description = accountDescription

	// create account
	createdAccount, err := client.Accounts.Add(awsAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access AWS-specific fields
	awsAccount = createdAccount.(*model.AmazonWebServicesAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", awsAccount.ID)
}
