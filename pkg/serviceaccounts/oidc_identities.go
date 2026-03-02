package serviceaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type OIDCIdentityQuery struct {
	ServiceAccountId string `uri:"serviceAccountId" url:"serviceAccountId"`
	Skip             int    `uri:"skip" url:"skip"`
	Take             int    `uri:"take" url:"take"`
}

type OIDCIdentity struct {
	Audience         string `json:"Audience"`
	Issuer           string `json:"Issuer"`
	Name             string `json:"Name"`
	ServiceAccountID string `json:"ServiceAccountId"`
	Subject          string `json:"Subject"`
	resources.Resource
}

type ServiceAccountOIDCIdentitiesResponse struct {
	ServerUrl      string          `json:"ServerUrl"`
	ExternalId     string          `json:"ExternalId"`
	OidcIdentities []*OIDCIdentity `json:"OidcIdentities"`
	Count          int             `json:"Count"`
}

// NewOIDCIdentity initializes a Service Account with required fields.
func NewOIDCIdentity(serviceAccountID string, name string, issuer string, subject string) *OIDCIdentity {
	return &OIDCIdentity{
		ServiceAccountID: serviceAccountID,
		Name:             name,
		Issuer:           issuer,
		Subject:          subject,
	}
}

const (
	serviceAccountOIDCIDQueryTemplate = "api/serviceaccounts/{serviceAccountId}/oidcidentities/v1{?skip,take}"
	serviceAccountOIDCCreate          = "api/serviceaccounts/{serviceAccountId}/oidcidentities/create/v1"
	serviceAccountOIDC                = "api/serviceaccounts{/serviceAccountId}/oidcidentities{/id}/v1"
)

// AddOIDCIdentity creates a new OIDC Identity for the service account
func AddOIDCIdentity(client newclient.Client, identity *OIDCIdentity) (*OIDCIdentity, error) {
	if identity == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterResource)
	}
	path, err := client.URITemplateCache().Expand(serviceAccountOIDCCreate, map[string]any{
		"serviceAccountId": identity.ServiceAccountID,
	})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Post[OIDCIdentity](client.HttpSession(), path, identity)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetOIDCIdentities queries all OIDC identities for the provided service account ID
func GetOIDCIdentities(client newclient.Client, query OIDCIdentityQuery) (*resources.Resources[*OIDCIdentity], error) {
	values, _ := uritemplates.Struct2map(query)
	if values == nil {
		values = map[string]any{}
	}
	path, err := client.URITemplateCache().Expand(serviceAccountOIDCIDQueryTemplate, values)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[resources.Resources[*OIDCIdentity]](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetServiceAccountOIDCData queries the service account and identities for the provided service account ID
func GetServiceAccountOIDCData(client newclient.Client, query OIDCIdentityQuery) (*ServiceAccountOIDCIdentitiesResponse, error) {
	if internal.IsEmpty(query.ServiceAccountId) {
		return nil, internal.CreateInvalidParameterError("GetServiceAccountOIDCData", "query.ServiceAccountId")
	}

	values, _ := uritemplates.Struct2map(query)
	if values == nil {
		values = map[string]any{}
	}

	path, err := client.URITemplateCache().Expand(serviceAccountOIDCIDQueryTemplate, values)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[ServiceAccountOIDCIdentitiesResponse](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetOIDCIdentityByID queries OIDC identities by ID for the provided service account ID
func GetOIDCIdentityByID(client newclient.Client, serviceAccountID string, ID string) (*OIDCIdentity, error) {
	path, err := client.URITemplateCache().Expand(serviceAccountOIDC, map[string]any{
		"serviceAccountId": serviceAccountID,
		"id":               ID,
	})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[OIDCIdentity](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateOIDCIdentity will update the OIDC identity, update does not return the updated OIDC identity
func UpdateOIDCIdentity(client newclient.Client, identity *OIDCIdentity) error {
	// update returns nothing
	path, err := client.URITemplateCache().Expand(serviceAccountOIDC, map[string]any{
		"serviceAccountId": identity.ServiceAccountID,
		"id":               identity.ID,
	})
	if err != nil {
		return err
	}

	_, err = newclient.Put[OIDCIdentity](client.HttpSession(), path, identity)
	if err != nil {
		return err
	}

	return nil
}

// DeleteOIDCIdentityByID remove an OIDC identity by ID for the provided service account ID
func DeleteOIDCIdentityByID(client newclient.Client, serviceAccountID string, ID string) error {
	path, err := client.URITemplateCache().Expand(serviceAccountOIDC, map[string]any{
		"serviceAccountId": serviceAccountID,
		"id":               ID,
	})
	if err != nil {
		return err
	}

	return newclient.Delete(client.HttpSession(), path)
}
