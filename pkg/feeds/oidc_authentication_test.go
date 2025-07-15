package feeds

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOidcAuthentication_NewAzureOidcAuthentication(t *testing.T) {
	oidc := NewAzureOidcAuthentication("client-id", "tenant-id", "audience", []string{"subject1", "subject2"})
	
	assert.Equal(t, OidcAuthenticationTypeAzure, oidc.Type)
	assert.Equal(t, "client-id", oidc.ClientId)
	assert.Equal(t, "tenant-id", oidc.TenantId)
	assert.Equal(t, "audience", oidc.Audience)
	assert.Equal(t, []string{"subject1", "subject2"}, oidc.SubjectKeys)
	
	azure, ok := oidc.GetAzure()
	assert.True(t, ok)
	assert.Equal(t, "client-id", azure.ClientId)
	assert.Equal(t, "tenant-id", azure.TenantId)
	assert.Equal(t, "audience", azure.Audience)
	assert.Equal(t, []string{"subject1", "subject2"}, azure.SubjectKeys)
}

func TestOidcAuthentication_NewAwsOidcAuthentication(t *testing.T) {
	oidc := NewAwsOidcAuthentication("3600", "audience", "role-arn", []string{"subject1", "subject2"})
	
	assert.Equal(t, OidcAuthenticationTypeAWS, oidc.Type)
	assert.Equal(t, "3600", oidc.SessionDuration)
	assert.Equal(t, "audience", oidc.Audience)
	assert.Equal(t, "role-arn", oidc.RoleArn)
	assert.Equal(t, []string{"subject1", "subject2"}, oidc.SubjectKeys)
	
	aws, ok := oidc.GetAWS()
	assert.True(t, ok)
	assert.Equal(t, "3600", aws.SessionDuration)
	assert.Equal(t, "audience", aws.Audience)
	assert.Equal(t, "role-arn", aws.RoleArn)
	assert.Equal(t, []string{"subject1", "subject2"}, aws.SubjectKeys)
}

func TestOidcAuthentication_NewGoogleOidcAuthentication(t *testing.T) {
	oidc := NewGoogleOidcAuthentication("audience", []string{"subject1", "subject2"})
	
	assert.Equal(t, OidcAuthenticationTypeGoogle, oidc.Type)
	assert.Equal(t, "audience", oidc.Audience)
	assert.Equal(t, []string{"subject1", "subject2"}, oidc.SubjectKeys)
	
	google, ok := oidc.GetGoogle()
	assert.True(t, ok)
	assert.Equal(t, "audience", google.Audience)
	assert.Equal(t, []string{"subject1", "subject2"}, google.SubjectKeys)
}

func TestOidcAuthentication_MarshalJSON_Azure(t *testing.T) {
	oidc := NewAzureOidcAuthentication("client-id", "tenant-id", "audience", []string{"subject1"})
	
	data, err := json.Marshal(oidc)
	assert.NoError(t, err)
	
	// Verify the JSON structure is flattened
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)
	
	assert.Equal(t, "Azure", result["Type"])
	assert.Equal(t, "client-id", result["ClientId"])
	assert.Equal(t, "tenant-id", result["TenantId"])
	assert.Equal(t, "audience", result["Audience"])
	assert.Equal(t, []interface{}{"subject1"}, result["SubjectKeys"])
	
	// Should not have nested Data field
	assert.Nil(t, result["Data"])
}

func TestOidcAuthentication_MarshalJSON_AWS(t *testing.T) {
	oidc := NewAwsOidcAuthentication("3600", "audience", "role-arn", []string{"subject1"})
	
	data, err := json.Marshal(oidc)
	assert.NoError(t, err)
	
	// Verify the JSON structure is flattened
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)
	
	assert.Equal(t, "AWS", result["Type"])
	assert.Equal(t, "3600", result["SessionDuration"])
	assert.Equal(t, "audience", result["Audience"])
	assert.Equal(t, "role-arn", result["RoleArn"])
	assert.Equal(t, []interface{}{"subject1"}, result["SubjectKeys"])
	
	// Should not have nested Data field
	assert.Nil(t, result["Data"])
}

func TestOidcAuthentication_MarshalJSON_Google(t *testing.T) {
	oidc := NewGoogleOidcAuthentication("audience", []string{"subject1"})
	
	data, err := json.Marshal(oidc)
	assert.NoError(t, err)
	
	// Verify the JSON structure is flattened
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)
	
	assert.Equal(t, "Google", result["Type"])
	assert.Equal(t, "audience", result["Audience"])
	assert.Equal(t, []interface{}{"subject1"}, result["SubjectKeys"])
	
	// Should not have nested Data field
	assert.Nil(t, result["Data"])
}

func TestOidcAuthentication_UnmarshalJSON_Azure(t *testing.T) {
	// Create Azure OIDC using constructor
	expectedOidc := NewAzureOidcAuthentication("client-id", "tenant-id", "audience", []string{"subject1", "subject2"})
	
	// Marshal to JSON
	data, err := json.Marshal(expectedOidc)
	assert.NoError(t, err)
	
	// Unmarshal back
	var oidc OidcAuthentication
	err = json.Unmarshal(data, &oidc)
	assert.NoError(t, err)
	
	// Verify the unmarshaled data matches the original
	assert.Equal(t, expectedOidc.Type, oidc.Type)
	assert.Equal(t, expectedOidc.ClientId, oidc.ClientId)
	assert.Equal(t, expectedOidc.TenantId, oidc.TenantId)
	assert.Equal(t, expectedOidc.Audience, oidc.Audience)
	assert.Equal(t, expectedOidc.SubjectKeys, oidc.SubjectKeys)
}

func TestOidcAuthentication_UnmarshalJSON_AWS(t *testing.T) {
	// Create AWS OIDC using constructor
	expectedOidc := NewAwsOidcAuthentication("3600", "audience", "role-arn", []string{"subject1", "subject2"})
	
	// Marshal to JSON
	data, err := json.Marshal(expectedOidc)
	assert.NoError(t, err)
	
	// Unmarshal back
	var oidc OidcAuthentication
	err = json.Unmarshal(data, &oidc)
	assert.NoError(t, err)
	
	// Verify the unmarshaled data matches the original
	assert.Equal(t, expectedOidc.Type, oidc.Type)
	assert.Equal(t, expectedOidc.SessionDuration, oidc.SessionDuration)
	assert.Equal(t, expectedOidc.Audience, oidc.Audience)
	assert.Equal(t, expectedOidc.RoleArn, oidc.RoleArn)
	assert.Equal(t, expectedOidc.SubjectKeys, oidc.SubjectKeys)
}

func TestOidcAuthentication_UnmarshalJSON_Google(t *testing.T) {
	// Create Google OIDC using constructor
	expectedOidc := NewGoogleOidcAuthentication("audience", []string{"subject1", "subject2"})
	
	// Marshal to JSON
	data, err := json.Marshal(expectedOidc)
	assert.NoError(t, err)
	
	// Unmarshal back
	var oidc OidcAuthentication
	err = json.Unmarshal(data, &oidc)
	assert.NoError(t, err)
	
	// Verify the unmarshaled data matches the original
	assert.Equal(t, expectedOidc.Type, oidc.Type)
	assert.Equal(t, expectedOidc.Audience, oidc.Audience)
	assert.Equal(t, expectedOidc.SubjectKeys, oidc.SubjectKeys)
}

func TestOidcAuthentication_GetWrongType(t *testing.T) {
	oidc := NewAzureOidcAuthentication("client-id", "tenant-id", "audience", []string{"subject1"})
	
	// Try to get AWS data from Azure OIDC
	aws, ok := oidc.GetAWS()
	assert.False(t, ok)
	assert.Nil(t, aws)
	
	// Try to get Google data from Azure OIDC
	google, ok := oidc.GetGoogle()
	assert.False(t, ok)
	assert.Nil(t, google)
}

func TestOidcAuthentication_UnmarshalJSON_EmptyClientId_IsAzure(t *testing.T) {
	jsonData := `{
		"ClientId": "",
		"TenantId": "tenant-id",
		"Audience": "audience",
		"SubjectKeys": ["subject1"]
	}`
	var oidc OidcAuthentication
	err := json.Unmarshal([]byte(jsonData), &oidc)
	assert.NoError(t, err)
	assert.Equal(t, OidcAuthenticationTypeAzure, oidc.Type)
	assert.Equal(t, "", oidc.ClientId)
	assert.Equal(t, "tenant-id", oidc.TenantId)
	assert.Equal(t, "audience", oidc.Audience)
	assert.Equal(t, []string{"subject1"}, oidc.SubjectKeys)
}

func TestOidcAuthentication_UnmarshalJSON_EmptyRoleArn_IsAWS(t *testing.T) {
	jsonData := `{
		"SessionDuration": "3600",
		"RoleArn": "",
		"Audience": "audience",
		"SubjectKeys": ["subject1"]
	}`
	var oidc OidcAuthentication
	err := json.Unmarshal([]byte(jsonData), &oidc)
	assert.NoError(t, err)
	assert.Equal(t, OidcAuthenticationTypeAWS, oidc.Type)
	assert.Equal(t, "", oidc.RoleArn)
	assert.Equal(t, "3600", oidc.SessionDuration)
	assert.Equal(t, "audience", oidc.Audience)
	assert.Equal(t, []string{"subject1"}, oidc.SubjectKeys)
}

func TestOidcAuthentication_UnmarshalJSON_OnlyAudienceAndSubjectKeys_IsGoogle(t *testing.T) {
	jsonData := `{
		"Audience": "audience",
		"SubjectKeys": ["subject1"]
	}`
	var oidc OidcAuthentication
	err := json.Unmarshal([]byte(jsonData), &oidc)
	assert.NoError(t, err)
	assert.Equal(t, OidcAuthenticationTypeGoogle, oidc.Type)
	assert.Equal(t, "audience", oidc.Audience)
	assert.Equal(t, []string{"subject1"}, oidc.SubjectKeys)
} 