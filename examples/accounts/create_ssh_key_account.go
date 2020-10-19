package examples

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

func CreateSSHKeyAccountExample() {
	var (
		octopusURL    string = "https://your_octopus_url"
		octopusAPIKey string = "API-YOUR_API_KEY"

		// SSH key-specific values
		username       string               = "account-username"
		privateKeyFile model.SensitiveValue = model.NewSensitiveValue("private-key")

		// Octopus Account values
		accountName string = "SSH Key Account"
		spaceName   string = "default"
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
	}

	sshKeyAccount := model.NewSSHKeyAccount(accountName, username, privateKeyFile)

	// create account
	createdAccount, err := client.Accounts.Add(sshKeyAccount)
	if err != nil {
		_ = fmt.Errorf("error adding account: %v", err)
	}

	// type conversion required to access SSH key-specific fields
	sshKeyAccount = createdAccount.(*model.SSHKeyAccount)

	// work with created account
	fmt.Printf("account created: (%s)\n", sshKeyAccount.ID)
}
