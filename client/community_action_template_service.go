package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// CommunityActionTemplateService handles communication with Account-related methods of the
// Octopus API.
type CommunityActionTemplateService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

// NewCommunityActionTemplateService returns an CommunityActionTemplateService with a preconfigured client.
func NewCommunityActionTemplateService(sling *sling.Sling) *CommunityActionTemplateService {
	if sling == nil {
		return nil
	}

	return &CommunityActionTemplateService{
		sling: sling,
		path:  "communityactiontemplates",
	}
}

// Get returns an Account that matches the input ID.
func (s *CommunityActionTemplateService) Get(id string) (*model.CommunityActionTemplate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("CommunityActionTemplateService: invalid parameter, id")
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
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(name) {
		return nil, errors.New("CommunityActionTemplateService: invalid parameter, name")
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

// Add creates a new CommunityActionTemplate.
func (s *CommunityActionTemplateService) Add(communityActionTemplate *model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if communityActionTemplate == nil {
		return nil, errors.New("CommunityActionTemplateService: invalid parameter, communityActionTemplate")
	}

	err = communityActionTemplate.Validate()

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
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	if isEmpty(id) {
		return errors.New("CommunityActionTemplateService: invalid parameter, id")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an CommunityActionTemplate based on the one provided as input.
func (s *CommunityActionTemplateService) Update(communityActionTemplate model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = communityActionTemplate.Validate()

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
		return fmt.Errorf("CommunityActionTemplateService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("CommunityActionTemplateService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &CommunityActionTemplateService{}
