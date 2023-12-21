package certificates

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

// certificateService handles communication with Certificate-related methods of the Octopus API.
type CertificateService struct {
	services.CanDeleteService
}

// NewCertificateService returns an certificateService with a preconfigured client.
func NewCertificateService(sling *sling.Sling, uriTemplate string) *CertificateService {
	return &CertificateService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceCertificateService, sling, uriTemplate),
		},
	}
}

// Add creates a new certificate.
//
// Deprecated: use certificates.Add
func (s *CertificateService) Add(certificate *CertificateResource) (*CertificateResource, error) {
	if IsNil(certificate) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterCertificate)
	}

	if err := certificate.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	path, err := services.GetAddPath(s, certificate)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), certificate, new(CertificateResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CertificateResource), nil
}

// Archive sets the status of a certificate to an archived state.
//
// Deprecated: use certificates.Archive
func (s *CertificateService) Archive(resource *CertificateResource) (*CertificateResource, error) {
	path := resource.Links["Archive"]

	_, err := services.ApiPost(s.GetClient(), resource, new(CertificateResource), path)
	if err != nil {
		return resource, err
	}

	return s.GetByID(resource.GetID())
}

// Get returns a collection of certificates based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
//
// Deprecated: use certificates.Get
func (s *CertificateService) Get(certificatesQuery CertificatesQuery) (*resources.Resources[*CertificateResource], error) {
	path, err := s.GetURITemplate().Expand(certificatesQuery)
	if err != nil {
		return &resources.Resources[*CertificateResource]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*CertificateResource]), path)
	if err != nil {
		return &resources.Resources[*CertificateResource]{}, err
	}

	return response.(*resources.Resources[*CertificateResource]), nil
}

// GetAll returns all certificates. If none are found or an error occurs, it
// returns an empty collection.
//
// Deprecated: use certificates.GetAll
func (s *CertificateService) GetAll() ([]*CertificateResource, error) {
	items := []*CertificateResource{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the certificate that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: use certificates.GetByID
func (s *CertificateService) GetByID(id string) (*CertificateResource, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(CertificateResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CertificateResource), nil
}

// Update modifies a Certificate based on the one provided as input.
//
// Deprecated: use certificates.Update
func (s *CertificateService) Update(resource CertificateResource) (*CertificateResource, error) {
	path, err := services.GetUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), resource, new(CertificateResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CertificateResource), nil
}

// Deprecated: use certificates.Replace
func (s *CertificateService) Replace(certificateID string, replacementCertificate *ReplacementCertificate) (*CertificateResource, error) {
	if internal.IsEmpty(certificateID) {
		return nil, internal.CreateInvalidParameterError("Replace", "certificateID")
	}

	if replacementCertificate == nil {
		return nil, internal.CreateInvalidParameterError("Replace", "replacementCertificate")
	}

	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	path := internal.TrimTemplate(s.GetPath())
	path = fmt.Sprintf(path+"/%s/replace", certificateID)

	if _, err := services.ApiPost(s.GetClient(), replacementCertificate, new(CertificateResource), path); err != nil {
		return nil, err
	}

	// The API endpoint /certificates/id/replace returns the old cert, we need
	// to re-query to get the updated one
	return s.GetByID(certificateID)
}

// Unarchive resets the status of an archived certificate.
//
// Deprecated: use certificates.Unarchive
func (s *CertificateService) Unarchive(resource *CertificateResource) (*CertificateResource, error) {
	path := resource.Links["Unarchive"]

	_, err := services.ApiPost(s.GetClient(), resource, new(CertificateResource), path)
	if err != nil {
		return resource, err
	}

	return s.GetByID(resource.GetID())
}

// --- new ---

const (
	template                   = "/api/{spaceId}/certificates{/id}{?skip,take,search,archived,tenant,firstResult,orderBy,ids,partialName}"
	certificateReplaceTemplate = "/api/{spaceID}/certificates/{id}/replace"
)

// Get returns a collection of certificates based on the criteria defined by its input
// query parameter. If an error occurs, a nil is returned along
// with the associated error.
func Get(client newclient.Client, spaceID string, certificatesQuery CertificatesQuery) (*resources.Resources[*CertificateResource], error) {
	return newclient.GetByQuery[CertificateResource](client, template, spaceID, certificatesQuery)
}

// Add creates a new certificate.
func Add(client newclient.Client, certificate *CertificateResource) (*CertificateResource, error) {
	return newclient.Add[CertificateResource](client, template, certificate.SpaceID, certificate)
}

// DeleteByID deletes a certificate based on the provided ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetByID returns the certificate that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*CertificateResource, error) {
	return newclient.GetByID[CertificateResource](client, template, spaceID, ID)
}

// Update modifies a Certificate based on the one provided as input.
func Update(client newclient.Client, resource *CertificateResource) (*CertificateResource, error) {
	return newclient.Update[CertificateResource](client, template, resource.SpaceID, resource.ID, resource)
}

// GetAll returns all certificates. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*CertificateResource, error) {
	return newclient.GetAll[CertificateResource](client, template, spaceID)
}

// Archive sets the status of a certificate to an archived state.
func Archive(client newclient.Client, spaceID string, resource *CertificateResource) (*CertificateResource, error) {
	path := resource.Links["Archive"]

	_, err := newclient.Post[CertificateResource](client.HttpSession(), path, resource)
	if err != nil {
		return resource, err
	}

	return newclient.GetByID[CertificateResource](client, template, spaceID, resource.GetID())
}

func Replace(client newclient.Client, spaceID string, certificateID string, replacementCertificate *ReplacementCertificate) (*CertificateResource, error) {
	if internal.IsEmpty(certificateID) {
		return nil, internal.CreateInvalidParameterError("Replace", "certificateID")
	}

	if replacementCertificate == nil {
		return nil, internal.CreateInvalidParameterError("Replace", "replacementCertificate")
	}

	templateParams := map[string]any{"spaceId": spaceID, "id": certificateID}
	expandedUri, err := client.URITemplateCache().Expand(certificateReplaceTemplate, templateParams)
	if err != nil {
		return nil, err
	}

	if _, err := newclient.Post[CertificateResource](client.HttpSession(), expandedUri, replacementCertificate); err != nil {
		return nil, err
	}

	// The API endpoint /certificates/id/replace returns the old cert, we need
	// to re-query to get the updated one
	return newclient.GetByID[CertificateResource](client, template, spaceID, certificateID)
}

// Unarchive resets the status of an archived certificate.
func Unarchive(client newclient.Client, spaceID string, resource *CertificateResource) (*CertificateResource, error) {
	path := resource.Links["Unarchive"]

	_, err := newclient.Post[CertificateResource](client.HttpSession(), path, resource)
	if err != nil {
		return resource, err
	}

	return newclient.GetByID[CertificateResource](client, template, spaceID, resource.GetID())
}
