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
	certificateService.service = newService(serviceCertificateService, sling, uriTemplate, new(Certificate))

	return certificateService
}

func (s certificateService) getPagedResponse(path string) ([]*Certificate, error) {
	resources := []*Certificate{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Certificates), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Certificates)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new certificate.
func (s certificateService) Add(resource *Certificate) (*Certificate, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Certificate), nil
}

// GetAll returns all certificates. If none can be found or an error occurs, it
// returns an empty collection.
func (s certificateService) GetAll() ([]*Certificate, error) {
	items := []*Certificate{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the certificate that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s certificateService) GetByID(id string) (*Certificate, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Certificate), nil
}

// GetByPartialName performs a lookup and returns instances of a Certificate with a matching partial name.
func (s certificateService) GetByPartialName(name string) ([]*Certificate, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Certificate{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a Certificate based on the one provided as input.
func (s certificateService) Update(resource Certificate) (*Certificate, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Certificate), nil
}

func (s certificateService) Replace(certificateID string, replacementCertificate *ReplacementCertificate) (*Certificate, error) {
	if isEmpty(certificateID) {
		return nil, createInvalidParameterError(operationReplace, parameterCertificateID)
	}

	if replacementCertificate == nil {
		return nil, createInvalidParameterError(operationReplace, parameterReplacementCertificate)
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/replace", certificateID)

	_, err = apiPost(s.getClient(), replacementCertificate, new(Certificate), path)
	if err != nil {
		return nil, err
	}

	//The API endpoint /certificates/id/replace returns the old cert, we need to re-query to get the updated one.
	return s.GetByID(certificateID)
}
