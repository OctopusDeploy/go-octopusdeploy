package examples

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/service"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateTokenExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// accountV1 values
		name  string                  = "Token Account"
		token *service.SensitiveValue = service.NewSensitiveValue("password-value")
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

	// option 1: create a token accountV1 and assign values to fields
	account, err := accounts.NewTokenAccount(name, token)
	if err != nil {
		_ = fmt.Errorf("error creating token accountV1: %v", err)
	}
	account.Description = "This is the description."

	// option 2: create a token accountV1 and assign values to fields using the
	// variadic configuration option
	options := func(t *accounts.TokenAccount) {
		t.Description = "This is the description."
	}

	account, err = accounts.NewTokenAccount(name, token, options)
	if err != nil {
		_ = fmt.Errorf("error creating token accountV1: %v", err)
	}

	// create accountV1
	createdAccount, err := client.Accounts.Add(account)
	if err != nil {
		_ = fmt.Errorf("error adding accountV1: %v", err)
	}

	// type conversion required to access token-specific fields
	account = createdAccount.(*accounts.TokenAccount)

	// work with created accountV1
	fmt.Printf("accountV1 created: (%s)\n", account.GetID())
}
