package accounts

type AccountType string

const (
	AccountTypeNone                       = AccountType("None")
	AccountTypeAmazonWebServicesAccount   = AccountType("AmazonWebServicesAccount")
	AccountTypeAzureServicePrincipal      = AccountType("AzureServicePrincipal")
	AccountTypeAzureOIDC                  = AccountType("AzureOidc")
	AccountTypeAwsOIDC                    = AccountType("AmazonWebServicesOidcAccount")
	AccountTypeAzureSubscription          = AccountType("AzureSubscription")
	AccountTypeGenericOIDCAccount         = AccountType("GenericOidcAccount")
	AccountTypeGoogleCloudPlatformAccount = AccountType("GoogleCloudAccount")
	AccountTypeSSHKeyPair                 = AccountType("SshKeyPair")
	AccountTypeToken                      = AccountType("Token")
	AccountTypeUsernamePassword           = AccountType("UsernamePassword")
)
