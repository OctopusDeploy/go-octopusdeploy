package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/service"
	service2 "github.com/OctopusDeploy/go-octopusdeploy/service"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
)

func CreateAmazonWebServicesAccountExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// AWS-specific values
		accessKey string                  = "access-key"
		secretKey *service.SensitiveValue = service.NewSensitiveValue("secret-key")

		// accountV1 values
		accountName        string = "AWS Account"
		accountDescription string = "My AWS Account"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	octoEndpoint, err := service2.NewOctopusServerEndpoint(apiURL, apiKey)
	client, err := service2.NewSpaceScopedClient(octoEndpoint, spaceID, nil)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	awsAccount, err := accounts.NewAmazonWebServicesAccount(accountName, accessKey, secretKey)
	if err != nil {
		_ = fmt.Errorf("error creating AWS accountV1: %v", err)
		return
	}

	// fill in accountV1 details
	awsAccount.Description = accountDescription

	// create accountV1
	createdAccount, err := client.Accounts.Add(awsAccount)
	if err != nil {
		_ = fmt.Errorf("error adding accountV1: %v", err)
	}

	// type conversion required to access AWS-specific fields
	awsAccount = createdAccount.(*accounts.AmazonWebServicesAccount)

	// work with created accountV1
	fmt.Printf("accountV1 created: (%s)\n", awsAccount.GetID())
}
