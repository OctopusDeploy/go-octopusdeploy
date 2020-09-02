package model

type AzureWebAppMachineEndpoint struct {
	ResourceGroupName string `json:"ResourceGroupName,omitempty"`
	WebAppName        string `json:"WebAppName,omitempty"`
	WebAppSlotName    int    `json:"WebAppSlotName"`
}
