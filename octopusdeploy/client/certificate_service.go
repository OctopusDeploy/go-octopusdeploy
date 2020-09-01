package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

// CertificateService handles communication with Certificate-related methods of
// the Octopus API.
type CertificateService struct {
	sling *sling.Sling
	path  string
}

// NewCertificateService returns an CertificateService with a preconfigured
// client.
func NewCertificateService(sling *sling.Sling) *CertificateService {
	return &CertificateService{
		sling: sling,
		path:  "certificates",
	}
}

func (s *CertificateService) Get(certificateID string) (*model.Certificate, error) {
	path := fmt.Sprintf(s.path+"/%s", certificateID)
	resp, err := apiGet(s.sling, new(model.Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

func (s *CertificateService) GetAll() (*[]model.Certificate, error) {
	var p []model.Certificate
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Certificates), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Certificates)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *CertificateService) GetByName(name string) (*model.Certificate, error) {
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

func (s *CertificateService) Add(resource *model.Certificate) (*model.Certificate, error) {
	resp, err := apiAdd(s.sling, resource, new(model.Certificate), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

// Delete removes the Certificate that matches the input ID.
func (s *CertificateService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *CertificateService) Update(resource *model.Certificate) (*model.Certificate, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

func (s *CertificateService) Replace(certificateID string, certificateReplace *model.CertificateReplace) (*model.Certificate, error) {
	path := fmt.Sprintf(s.path+"/%s/replace", certificateID)
	_, err := apiPost(s.sling, certificateReplace, new(model.Certificate), path)

	if err != nil {
		return nil, err
	}

	//The API endpoint /certificates/id/replace returns the old cert, we need to re-query to get the updated one.
	return s.Get(certificateID)
}
