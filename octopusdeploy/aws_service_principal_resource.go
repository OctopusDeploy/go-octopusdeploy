package octopusdeploy

type AwsServicePrincipalResource struct {
	AccessKey string          `json:"AccessKey,omitempty"`
	SecretKey *SensitiveValue `json:"SecretKey,omitempty"`
}
