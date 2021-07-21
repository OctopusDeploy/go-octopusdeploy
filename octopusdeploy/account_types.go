package octopusdeploy

type AccountType string

const (
	AccountTypeNone                     = AccountType("None")
	AccountTypeAmazonWebServicesAccount = AccountType("AmazonWebServicesAccount")
	AccountTypeAzureServicePrincipal    = AccountType("AzureServicePrincipal")
	AccountTypeAzureSubscription        = AccountType("AzureSubscription")
	AccountTypeGoogleCloudAccount       = AccountType("GoogleCloudAccount")
	AccountTypeSSHKeyPair               = AccountType("SshKeyPair")
	AccountTypeToken                    = AccountType("Token")
	AccountTypeUsernamePassword         = AccountType("UsernamePassword")
)
