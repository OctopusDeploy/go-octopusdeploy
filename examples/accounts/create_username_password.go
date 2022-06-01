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
		password string = "password-value"
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

	// option 1: create a username/password account and assign values to fields
	usernamePasswordAccount, err := octopusdeploy.NewUsernamePasswordAccount(name)
	if err != nil {
		_ = fmt.Errorf("error creating username/password account: %v", err)
	}
	usernamePasswordAccount.SetPassword(octopusdeploy.NewSensitiveValue(password))
	usernamePasswordAccount.SetUsername(username)

	// option 2: create a username/password account and assign values to fields
	// using the variadic configuration option
	options := func(u *octopusdeploy.UsernamePasswordAccount) {
		u.Password = octopusdeploy.NewSensitiveValue(password)
		u.Username = username
	}

	usernamePasswordAccount, err = octopusdeploy.NewUsernamePasswordAccount(name, options)
	if err != nil {
		_ = fmt.Errorf("error creating username/password account: %v", err)
	}

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
