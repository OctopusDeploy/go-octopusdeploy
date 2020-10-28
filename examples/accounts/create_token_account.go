package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateTokenExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// account values
		name  string                        = "Token Account"
		token *octopusdeploy.SensitiveValue = octopusdeploy.NewSensitiveValue("password-value")
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

	// option 1: create a token account and assign values to fields
	account, err := octopusdeploy.NewTokenAccount(name, token)
	if err != nil {
		_ = fmt.Errorf("error creating token account: %v", err)
	}
	account.Description = "This is the description."

	// option 2: create a token account and assign values to fields using the
	// variadic configuration option
	options := func(t *octopusdeploy.TokenAccount) {
		t.Description = "This is the description."
	}

	account, err = octopusdeploy.NewTokenAccount(name, token, options)
	if err != nil {
		_ = fmt.Errorf("error creating token account: %v", err)
	}

	// create account
	createdAccount, err := client.Accounts.Add(account)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access token-specific fields
	account = createdAccount.(*octopusdeploy.TokenAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", account.GetID())
}
