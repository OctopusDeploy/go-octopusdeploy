package accounts

type AccountType string

const (
	AccountTypeNone                       = AccountType("None")
	AccountTypeAmazonWebServicesAccount   = AccountType("AmazonWebServicesAccount")
	AccountTypeAzureServicePrincipal      = AccountType("AzureServicePrincipal")
	AccountTypeAzureOIDC                  = AccountType("AzureOIDC")
	AccountTypeAzureSubscription          = AccountType("AzureSubscription")
	AccountTypeGoogleCloudPlatformAccount = AccountType("GoogleCloudAccount")
	AccountTypeSSHKeyPair                 = AccountType("SshKeyPair")
	AccountTypeToken                      = AccountType("Token")
	AccountTypeUsernamePassword           = AccountType("UsernamePassword")
)
