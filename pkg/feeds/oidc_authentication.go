package feeds

import (
	"encoding/json"
)

// OidcAuthentication represents a union type that can hold any of the three OIDC authentication types
// The properties are flattened at the top level for API compatibility
type OidcAuthentication struct {
	// Type indicates which OIDC authentication type this represents
	Type OidcAuthenticationType `json:"Type"`

	// Common fields across all OIDC types
	Audience    string   `json:"Audience,omitempty"`
	SubjectKeys []string `json:"SubjectKeys,omitempty"`

	// Azure-specific fields
	ClientId string `json:"ClientId,omitempty"`
	TenantId string `json:"TenantId,omitempty"`

	// AWS-specific fields
	SessionDuration string `json:"SessionDuration,omitempty"`
	RoleArn         string `json:"RoleArn,omitempty"`
}

// OidcAuthenticationType represents the type of OIDC authentication
type OidcAuthenticationType string

const (
	OidcAuthenticationTypeAzure  OidcAuthenticationType = "Azure"
	OidcAuthenticationTypeAWS    OidcAuthenticationType = "AWS"
	OidcAuthenticationTypeGoogle OidcAuthenticationType = "Google"
)

// NewAzureOidcAuthentication creates a new OIDC authentication for Azure
func NewAzureOidcAuthentication(clientId, tenantId, audience string, subjectKeys []string) *OidcAuthentication {
	return &OidcAuthentication{
		Type:        OidcAuthenticationTypeAzure,
		ClientId:    clientId,
		TenantId:    tenantId,
		Audience:    audience,
		SubjectKeys: subjectKeys,
	}
}

// NewAwsOidcAuthentication creates a new OIDC authentication for AWS
func NewAwsOidcAuthentication(sessionDuration, audience, roleArn string, subjectKeys []string) *OidcAuthentication {
	return &OidcAuthentication{
		Type:            OidcAuthenticationTypeAWS,
		SessionDuration: sessionDuration,
		Audience:        audience,
		RoleArn:         roleArn,
		SubjectKeys:     subjectKeys,
	}
}

// NewGoogleOidcAuthentication creates a new OIDC authentication for Google
func NewGoogleOidcAuthentication(audience string, subjectKeys []string) *OidcAuthentication {
	return &OidcAuthentication{
		Type:        OidcAuthenticationTypeGoogle,
		Audience:    audience,
		SubjectKeys: subjectKeys,
	}
}

// GetAzure returns the Azure OIDC authentication data if this is an Azure type
func (o *OidcAuthentication) GetAzure() (*AzureContainerRegistryOidcAuthentication, bool) {
	if o.Type == OidcAuthenticationTypeAzure {
		return &AzureContainerRegistryOidcAuthentication{
			ClientId:    o.ClientId,
			TenantId:    o.TenantId,
			Audience:    o.Audience,
			SubjectKeys: o.SubjectKeys,
		}, true
	}
	return nil, false
}

// GetAWS returns the AWS OIDC authentication data if this is an AWS type
func (o *OidcAuthentication) GetAWS() (*AwsElasticContainerRegistryOidcAuthentication, bool) {
	if o.Type == OidcAuthenticationTypeAWS {
		return &AwsElasticContainerRegistryOidcAuthentication{
			SessionDuration: o.SessionDuration,
			Audience:        o.Audience,
			SubjectKeys:     o.SubjectKeys,
			RoleArn:         o.RoleArn,
		}, true
	}
	return nil, false
}

// GetGoogle returns the Google OIDC authentication data if this is a Google type
func (o *OidcAuthentication) GetGoogle() (*GoogleContainerRegistryOidcAuthentication, bool) {
	if o.Type == OidcAuthenticationTypeGoogle {
		return &GoogleContainerRegistryOidcAuthentication{
			Audience:    o.Audience,
			SubjectKeys: o.SubjectKeys,
		}, true
	}
	return nil, false
}

// MarshalJSON implements custom JSON marshaling to handle the union type
func (o *OidcAuthentication) MarshalJSON() ([]byte, error) {
	// Create a map to hold the fields we want to serialize
	result := make(map[string]interface{})

	// Always include the Type field
	result["Type"] = o.Type

	// Add common fields
	if o.Audience != "" {
		result["Audience"] = o.Audience
	}
	if len(o.SubjectKeys) > 0 {
		result["SubjectKeys"] = o.SubjectKeys
	}

	// Add type-specific fields based on the type
	switch o.Type {
	case OidcAuthenticationTypeAzure:
		if o.ClientId != "" {
			result["ClientId"] = o.ClientId
		}
		if o.TenantId != "" {
			result["TenantId"] = o.TenantId
		}
	case OidcAuthenticationTypeAWS:
		if o.SessionDuration != "" {
			result["SessionDuration"] = o.SessionDuration
		}
		if o.RoleArn != "" {
			result["RoleArn"] = o.RoleArn
		}
	case OidcAuthenticationTypeGoogle:
		// Google only has common fields, no additional ones
	}

	return json.Marshal(result)
}

// UnmarshalJSON implements custom JSON unmarshaling to handle the union type
func (o *OidcAuthentication) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*o = OidcAuthentication{}
		return nil
	}

	// First, unmarshal into a map to inspect the fields
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Discriminate based on properties. This is hacky, but we would have to add a "Type" property to the API
	// To handle this properly
	if _, hasClientId := raw["ClientId"]; hasClientId {
		o.Type = OidcAuthenticationTypeAzure
	} else if _, hasRoleArn := raw["RoleArn"]; hasRoleArn {
		o.Type = OidcAuthenticationTypeAWS
	} else {
		o.Type = OidcAuthenticationTypeGoogle
	}

	// Now unmarshal into the appropriate struct to get all fields
	switch o.Type {
	case OidcAuthenticationTypeAzure:
		var azure AzureContainerRegistryOidcAuthentication
		o.Type = OidcAuthenticationTypeAzure
		o.ClientId = azure.ClientId
		o.TenantId = azure.TenantId
		o.Audience = azure.Audience
		o.SubjectKeys = azure.SubjectKeys

	case OidcAuthenticationTypeAWS:
		var aws AwsElasticContainerRegistryOidcAuthentication
		o.Type = OidcAuthenticationTypeAWS
		o.SessionDuration = aws.SessionDuration
		o.Audience = aws.Audience
		o.SubjectKeys = aws.SubjectKeys
		o.RoleArn = aws.RoleArn

	case OidcAuthenticationTypeGoogle:
		var google GoogleContainerRegistryOidcAuthentication
		o.Type = OidcAuthenticationTypeGoogle
		o.Audience = google.Audience
		o.SubjectKeys = google.SubjectKeys
	}

	return nil
}
