package model

type AzureWebAppMachineEndpoint struct {
	ResourceGroupName string `json:ResourceGroupName`
	WebAppName        string `json:WebAppName`
	WebAppSlotName    int    `json:WebAppSlotName`
}
