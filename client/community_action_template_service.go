package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// CommunityActionTemplateService handles communication with Account-related methods of the
// Octopus API.
type CommunityActionTemplateService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

// NewCommunityActionTemplateService returns an CommunityActionTemplateService with a preconfigured client.
func NewCommunityActionTemplateService(sling *sling.Sling, uriTemplate string) *CommunityActionTemplateService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &CommunityActionTemplateService{
		name:  "CommunityActionTemplateService",
		path:  path,
		sling: sling,
	}
}

// Get returns an Account that matches the input ID.
func (s *CommunityActionTemplateService) Get(id string) (*model.CommunityActionTemplate, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.CommunityActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

// GetAll returns all instances of a CommunityActionTemplate.
func (s *CommunityActionTemplateService) GetAll() (*[]model.CommunityActionTemplate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	var p []model.CommunityActionTemplate
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.CommunityActionTemplates), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.CommunityActionTemplates)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName performs a lookup and returns the CommunityActionTemplate with a matching name.
func (s *CommunityActionTemplateService) GetByName(name string) (*model.CommunityActionTemplate, error) {
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

	for _, item := range *collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// Add creates a new CommunityActionTemplate.
func (s *CommunityActionTemplateService) Add(communityActionTemplate *model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	if communityActionTemplate == nil {
		return nil, createInvalidParameterError("Add", "communityActionTemplate")
	}

	err := communityActionTemplate.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, communityActionTemplate, new(model.CommunityActionTemplate), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

// Delete removes the CommunityActionTemplate that matches the input ID.
func (s *CommunityActionTemplateService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an CommunityActionTemplate based on the one provided as input.
func (s *CommunityActionTemplateService) Update(communityActionTemplate model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	err := communityActionTemplate.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", communityActionTemplate.ID)
	resp, err := apiUpdate(s.sling, communityActionTemplate, new(model.CommunityActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

func (s *CommunityActionTemplateService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &CommunityActionTemplateService{}
