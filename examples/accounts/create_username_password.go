package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateUsernamePasswordExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// account values
		name     string = "Username/Password Account"
		username string = "account-username"
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

	usernamePasswordAccount := octopusdeploy.NewUsernamePasswordAccount(name)

	// fill in account details
	usernamePasswordAccount.Username = username

	// create account
	createdAccount, err := client.Accounts.Add(usernamePasswordAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access Username/Password-specific fields
	usernamePasswordAccount = createdAccount.(*octopusdeploy.UsernamePasswordAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", usernamePasswordAccount.GetID())
}
