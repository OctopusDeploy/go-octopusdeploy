package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

// certificateService handles communication with Certificate-related methods of the Octopus API.
type certificateService struct {
	canDeleteService
}

// newCertificateService returns an certificateService with a preconfigured client.
func newCertificateService(sling *sling.Sling, uriTemplate string) *certificateService {
	certificateService := &certificateService{}
	certificateService.service = newService(ServiceCertificateService, sling, uriTemplate)

	return certificateService
}

// Add creates a new certificate.
func (s certificateService) Add(resource *CertificateResource) (*CertificateResource, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(CertificateResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CertificateResource), nil
}

// Get returns a collection of certificates based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s certificateService) Get(certificatesQuery CertificatesQuery) (*CertificateResources, error) {
	path, err := s.getURITemplate().Expand(certificatesQuery)
	if err != nil {
		return &CertificateResources{}, err
	}

	response, err := apiGet(s.getClient(), new(CertificateResources), path)
	if err != nil {
		return &CertificateResources{}, err
	}

	return response.(*CertificateResources), nil
}

// GetAll returns all certificates. If none are found or an error occurs, it
// returns an empty collection.
func (s certificateService) GetAll() ([]*CertificateResource, error) {
	items := []*CertificateResource{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the certificate that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s certificateService) GetByID(id string) (*CertificateResource, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(CertificateResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CertificateResource), nil
}

// Update modifies a Certificate based on the one provided as input.
func (s certificateService) Update(resource CertificateResource) (*CertificateResource, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(CertificateResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CertificateResource), nil
}

func (s certificateService) Replace(certificateID string, replacementCertificate *ReplacementCertificate) (*CertificateResource, error) {
	if isEmpty(certificateID) {
		return nil, createInvalidParameterError(OperationReplace, ParameterCertificateID)
	}

	if replacementCertificate == nil {
		return nil, createInvalidParameterError(OperationReplace, ParameterReplacementCertificate)
	}

	if err := validateInternalState(s); err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/replace", certificateID)

	if _, err := apiPost(s.getClient(), replacementCertificate, new(CertificateResource), path); err != nil {
		return nil, err
	}

	// The API endpoint /certificates/id/replace returns the old cert, we need
	// to re-query to get the updated one
	return s.GetByID(certificateID)
}
