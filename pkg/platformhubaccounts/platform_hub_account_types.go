package platformhubaccounts

type PlatformHubAccountType string

const (
	AccountTypePlatformHubAwsAccount                   = PlatformHubAccountType("AmazonWebServicesAccount")
	AccountTypePlatformHubAwsOIDCAccount               = PlatformHubAccountType("AmazonWebServicesOidcAccount")
	AccountTypePlatformHubAzureOidcAccount             = PlatformHubAccountType("AzureOidc")
	AccountTypePlatformHubAzureServicePrincipalAccount = PlatformHubAccountType("AzureServicePrincipal")
	AccountTypePlatformHubGcpAccount                   = PlatformHubAccountType("GoogleCloudAccount")
	AccountTypePlatformHubGenericOidcAccount           = PlatformHubAccountType("GenericOidcAccount")
	AccountTypePlatformHubUsernamePasswordAccount      = PlatformHubAccountType("UsernamePassword")
)
