package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// CertificateService handles communication with Certificate-related methods of
// the Octopus API.
type CertificateService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

// NewCertificateService returns an CertificateService with a preconfigured
// client.
func NewCertificateService(sling *sling.Sling, uriTemplate string) *CertificateService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &CertificateService{
		sling: sling,
		path:  path,
	}
}

func (s *CertificateService) Get(id string) (*model.Certificate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("CertificateService: invalid parameter, id")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

// GetAll returns all instances of a Certificate.
func (s *CertificateService) GetAll() (*[]model.Certificate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Certificate), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Certificate), nil
}

// GetByName performs a lookup and returns the Certificate with a matching name.
func (s *CertificateService) GetByName(name string) (*model.Certificate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("CertificateService: invalid parameter, name")
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

// Add creates a new Certificate.
func (s *CertificateService) Add(certificate *model.Certificate) (*model.Certificate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if certificate == nil {
		return nil, errors.New("CertificateService: invalid parameter, certificate")
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
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("CertificateService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *CertificateService) Update(certificate model.Certificate) (*model.Certificate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = certificate.Validate()

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
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if len(strings.Trim(certificateID, " ")) == 0 {
		return nil, errors.New("CertificateService: invalid parameter, certificateID")
	}

	if replacementCertificate == nil {
		return nil, errors.New("CertificateService: invalid parameter, replacementCertificate")
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
		return fmt.Errorf("CertificateService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("CertificateService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &CertificateService{}
