package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// certificateService handles communication with Certificate-related methods of the Octopus API.
type certificateService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

// newCertificateService returns an certificateService with a preconfigured client.
func newCertificateService(sling *sling.Sling, uriTemplate string) *certificateService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &certificateService{
		name:        serviceCertificateService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s certificateService) getClient() *sling.Sling {
	return s.sling
}

func (s certificateService) getName() string {
	return s.name
}

func (s certificateService) getPagedResponse(path string) ([]model.Certificate, error) {
	resources := []model.Certificate{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Certificates), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Certificates)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s certificateService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new certificate.
func (s certificateService) Add(resource *model.Certificate) (*model.Certificate, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Certificate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

// DeleteByID deletes the certificate that matches the input ID. If the ID
// cannot be found, it returns nil and an error.
func (s certificateService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("certificate", "ID", id)
	}

	return err
}

// GetAll returns all certificates. If none can be found or an error occurs, it
// returns an empty collection.
func (s certificateService) GetAll() ([]model.Certificate, error) {
	items := []model.Certificate{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the certificate that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s certificateService) GetByID(id string) (*model.Certificate, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Certificate), path)
	if err != nil {
		return nil, createResourceNotFoundError("certificate", "ID", id)
	}

	return resp.(*model.Certificate), nil
}

// GetByPartialName performs a lookup and returns instances of a Certificate with a matching partial name.
func (s certificateService) GetByPartialName(name string) ([]model.Certificate, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.Certificate{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a Certificate based on the one provided as input.
func (s certificateService) Update(resource model.Certificate) (*model.Certificate, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Certificate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Certificate), nil
}

func (s certificateService) Replace(certificateID string, replacementCertificate *model.ReplacementCertificate) (*model.Certificate, error) {
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

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s/replace", certificateID)

	_, err = apiPost(s.getClient(), replacementCertificate, new(model.ReplacementCertificate), path)
	if err != nil {
		return nil, err
	}

	//The API endpoint /certificates/id/replace returns the old cert, we need to re-query to get the updated one.
	return s.GetByID(certificateID)
}

var _ ServiceInterface = &certificateService{}
