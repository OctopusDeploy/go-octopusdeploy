package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// CertificateService handles communication with Certificate-related methods of
// the Octopus API.
type CertificateService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

// NewCertificateService returns an CertificateService with a preconfigured
// client.
func NewCertificateService(sling *sling.Sling, uriTemplate string) *CertificateService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &CertificateService{
		name:  "CertificateService",
		path:  path,
		sling: sling,
	}
}

func (s *CertificateService) Get(id string) (*model.Certificate, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

// GetAll returns all instances of a Certificate.
func (s *CertificateService) GetAll() ([]model.Certificate, error) {
	err := s.validateInternalState()

	items := new([]model.Certificate)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the Certificate with a matching name.
func (s *CertificateService) GetByName(name string) (*model.Certificate, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("GetByName", "name")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// Add creates a new Certificate.
func (s *CertificateService) Add(certificate *model.Certificate) (*model.Certificate, error) {
	if certificate == nil {
		return nil, createInvalidParameterError("Add", "certificate")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = certificate.Validate()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, certificate, new(model.Certificate), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

// Delete removes the Certificate that matches the input ID.
func (s *CertificateService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *CertificateService) Update(certificate model.Certificate) (*model.Certificate, error) {
	err := certificate.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", certificate.ID)
	resp, err := apiUpdate(s.sling, certificate, new(model.Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

func (s *CertificateService) Replace(certificateID string, replacementCertificate *model.ReplacementCertificate) (*model.Certificate, error) {
	if isEmpty(certificateID) {
		return nil, createInvalidParameterError("Replace", "certificateID")
	}

	if replacementCertificate == nil {
		return nil, createInvalidParameterError("Replace", "replacementCertificate")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s/replace", certificateID)
	_, err = apiPost(s.sling, replacementCertificate, new(model.ReplacementCertificate), path)

	if err != nil {
		return nil, err
	}

	//The API endpoint /certificates/id/replace returns the old cert, we need to re-query to get the updated one.
	return s.Get(certificateID)
}

func (s *CertificateService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &CertificateService{}
