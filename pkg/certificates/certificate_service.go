package certificates

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
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
func (s *CertificateService) Get(certificatesQuery CertificatesQuery) (*CertificateResources, error) {
	path, err := s.GetURITemplate().Expand(certificatesQuery)
	if err != nil {
		return &CertificateResources{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(CertificateResources), path)
	if err != nil {
		return &CertificateResources{}, err
	}

	return response.(*CertificateResources), nil
}

// GetAll returns all certificates. If none are found or an error occurs, it
// returns an empty collection.
func (s *CertificateService) GetAll() ([]*CertificateResource, error) {
	items := []*CertificateResource{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the certificate that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *CertificateService) GetByID(id string) (*CertificateResource, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(CertificateResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CertificateResource), nil
}

// Update modifies a Certificate based on the one provided as input.
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
func (s *CertificateService) Unarchive(resource *CertificateResource) (*CertificateResource, error) {
	path := resource.Links["Unarchive"]

	_, err := services.ApiPost(s.GetClient(), resource, new(CertificateResource), path)
	if err != nil {
		return resource, err
	}

	return s.GetByID(resource.GetID())
}
