package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
)

func CreateTokenExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// account values
		name  string               = "Token Account"
		token *core.SensitiveValue = core.NewSensitiveValue("password-value")
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

	// create a token account and assign values to fields
	account, err := accounts.NewTokenAccount(name, token)
	if err != nil {
		_ = fmt.Errorf("error creating token account: %v", err)
	}
	account.Description = "This is the description."

	// create account
	createdAccount, err := client.Accounts.Add(account)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access token-specific fields
	account = createdAccount.(*accounts.TokenAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", account.GetID())
}
