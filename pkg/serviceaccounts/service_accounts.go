package serviceaccounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type OIDCIdentityQuery struct {
	Skip int `uri:"skip,omitempty" url:"skip,omitempty"`
	Take int `uri:"take,omitempty" url:"take,omitempty"`
}

type OIDCIdentity struct {
	Audience         string `json:"Audience"`
	Issuer           string `json:"Issuer"`
	Name             string `json:"Name"`
	ServiceAccountID string `json:"ServiceAccountId"`
	Subject          string `json:"Subject"`
	resources.Resource
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

// Add creates a new OIDC Identity for the service account
func Add(client newclient.Client, identity *OIDCIdentity) (*OIDCIdentity, error) {
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

// Get queries all OIDC identities for the provided service account ID
func Get(client newclient.Client, query OIDCIdentityQuery) (*resources.Resources[*OIDCIdentity], error) {
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

// GetByID queries OIDC identities by ID for the provided service account ID
func GetByID(client newclient.Client, serviceAccountID string, ID string) (*OIDCIdentity, error) {
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

// Update will update the OIDC identity, update does not return the updated OIDC identity
func Update(client newclient.Client, identity *OIDCIdentity) error {
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

// DeleteByID remove an OIDC identity by ID for the provided service account ID
func DeleteByID(client newclient.Client, serviceAccountID string, ID string) error {
	path, err := client.URITemplateCache().Expand(serviceAccountOIDC, map[string]any{
		"serviceAccountId": serviceAccountID,
		"id":               ID,
	})
	if err != nil {
		return err
	}

	return newclient.Delete(client.HttpSession(), path)
}
