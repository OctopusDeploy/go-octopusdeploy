package platformhubaccounts

type PlatformHubAccountType string

const (
	AccountTypePlatformHubAwsAccount     = PlatformHubAccountType("AmazonWebServicesAccount")
	AccountTypePlatformHubAwsOIDCAccount = PlatformHubAccountType("AmazonWebServicesOidcAccount")
)
