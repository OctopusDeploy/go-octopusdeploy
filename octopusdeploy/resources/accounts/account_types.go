package accounts

type AccountType string

const (
	AccountTypeNone                       = AccountType("None")
	AccountTypeAmazonWebServicesAccount   = AccountType("AmazonWebServicesAccount")
	AccountTypeAzureServicePrincipal      = AccountType("AzureServicePrincipal")
	AccountTypeAzureSubscription          = AccountType("AzureSubscription")
	AccountTypeGoogleCloudPlatformAccount = AccountType("GoogleCloudAccount")
	AccountTypeSSHKeyPair                 = AccountType("SshKeyPair")
	AccountTypeToken                      = AccountType("Token")
	AccountTypeUsernamePassword           = AccountType("UsernamePassword")
)
