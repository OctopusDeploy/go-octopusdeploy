package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
)

func CreateAmazonWebServicesAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// AWS-specific values
		accessKey string               = "access-key"
		secretKey *core.SensitiveValue = core.NewSensitiveValue("secret-key")

		// account values
		accountName        string = "AWS Account"
		accountDescription string = "My AWS Account"
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
