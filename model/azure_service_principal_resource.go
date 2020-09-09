package model

import uuid "github.com/google/uuid"

type AzureServicePrincipalResource struct {
	AzureEnvironment               string     `json:"AzureEnvironment,omitempty"`
	ActiveDirectoryEndpointBase    string     `json:"ActiveDirectoryEndpointBaseUri,omitempty"`
	ClientID                       *uuid.UUID `json:"ClientId,omitempty"`
	ResourceManagementEndpointBase string     `json:"ResourceManagementEndpointBaseUri,omitempty"`
	SubscriptionNumber             *uuid.UUID `json:"SubscriptionNumber,omitempty"`
	TenantID                       *uuid.UUID `json:"TenantId,omitempty"`
}