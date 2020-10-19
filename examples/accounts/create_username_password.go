package examples

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

func CreateUsernamePasswordExample() {
	var (
		octopusURL    string = "https://your_octopus_url"
		octopusAPIKey string = "API-YOUR_API_KEY"

		// Username/Password-specific values
		username string = "account-username"

		// Octopus Account values
		accountName string = "Username/Password Account"
		spaceName   string = "default"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
	}

	usernamePasswordAccount := model.NewUsernamePasswordAccount(accountName)

	// fill in account details
	usernamePasswordAccount.Username = username

	// create account
	createdAccount, err := client.Accounts.Add(usernamePasswordAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access Username/Password-specific fields
	usernamePasswordAccount = createdAccount.(*model.UsernamePasswordAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", usernamePasswordAccount.ID)
}
