package accounts

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *AccountResource:
		return v == nil
	case *AmazonWebServicesAccount:
		return v == nil
	case *AzureServicePrincipalAccount:
		return v == nil
	case *AzureOIDCAccount:
		return v == nil
	case *AzureSubscriptionAccount:
		return v == nil
	case *GoogleCloudPlatformAccount:
		return v == nil
	case *SSHKeyAccount:
		return v == nil
	case *TokenAccount:
		return v == nil
	case *UsernamePasswordAccount:
		return v == nil
	default:
		return v == nil
	}
}
