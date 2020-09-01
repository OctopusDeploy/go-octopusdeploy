package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

// CommunityActionTemplateService handles communication with Account-related methods of the
// Octopus API.
type CommunityActionTemplateService struct {
	sling *sling.Sling
	path  string
}

// NewCommunityActionTemplateService returns an CommunityActionTemplateService with a preconfigured client.
func NewCommunityActionTemplateService(sling *sling.Sling) *CommunityActionTemplateService {
	return &CommunityActionTemplateService{
		sling: sling,
		path:  "communityactiontemplates",
	}
}

// Get returns an Account that matches the input ID.
func (s *CommunityActionTemplateService) Get(id string) (*model.CommunityActionTemplate, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.CommunityActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

// GetAll returns all of the Accounts for a Space.
func (s *CommunityActionTemplateService) GetAll() (*[]model.CommunityActionTemplate, error) {
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

// GetByName returns an CommunityActionTemplate that matches the input name.
func (s *CommunityActionTemplateService) GetByName(name string) (*model.CommunityActionTemplate, error) {
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

// Add creates a new Account.
func (s *CommunityActionTemplateService) Add(resource *model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	resp, err := apiAdd(s.sling, resource, new(model.CommunityActionTemplate), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

// Delete removes the CommunityActionTemplate that matches the input ID.
func (s *CommunityActionTemplateService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an CommunityActionTemplate based on the one provided as input.
func (s *CommunityActionTemplateService) Update(resource *model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.CommunityActionTemplate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}
