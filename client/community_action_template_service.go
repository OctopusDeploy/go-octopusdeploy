package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// communityActionTemplateService handles communication with Account-related methods of the Octopus API.
type communityActionTemplateService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

// newCommunityActionTemplateService returns an communityActionTemplateService with a preconfigured client.
func newCommunityActionTemplateService(sling *sling.Sling, uriTemplate string) *communityActionTemplateService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &communityActionTemplateService{
		name:        serviceCommunityActionTemplateService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s communityActionTemplateService) getClient() *sling.Sling {
	return s.sling
}

func (s communityActionTemplateService) getName() string {
	return s.name
}

func (s communityActionTemplateService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// GetByID returns an Account that matches the input ID. If one cannot be found, it returns nil and an error.
func (s communityActionTemplateService) GetByID(id string) (*model.CommunityActionTemplate, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.CommunityActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

// GetAll returns all instances of a CommunityActionTemplate. If none can be found or an error occurs, it returns an empty collection.
func (s communityActionTemplateService) GetAll() (*[]model.CommunityActionTemplate, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	var p []model.CommunityActionTemplate
	path := trimTemplate(s.path)
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.CommunityActionTemplates), path)
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
func (s communityActionTemplateService) GetByName(name string) (*model.CommunityActionTemplate, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError(operationGetByName, parameterName)
	}

	err := validateInternalState(s)
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

	return nil, createItemNotFoundError(s.name, operationGetByName, name)
}

// Add creates a new CommunityActionTemplate.
func (s communityActionTemplateService) Add(communityActionTemplate *model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	if communityActionTemplate == nil {
		return nil, createInvalidParameterError(operationAdd, "communityActionTemplate")
	}

	err := communityActionTemplate.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)

	resp, err := apiAdd(s.getClient(), communityActionTemplate, new(model.CommunityActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

// DeleteByID deletes the CommunityActionTemplate that matches the input ID.
func (s communityActionTemplateService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// Update modifies an CommunityActionTemplate based on the one provided as input.
func (s communityActionTemplateService) Update(communityActionTemplate model.CommunityActionTemplate) (*model.CommunityActionTemplate, error) {
	err := communityActionTemplate.Validate()

	if err != nil {
		return nil, err
	}

	err = validateInternalState(s)

	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s", communityActionTemplate.ID)

	resp, err := apiUpdate(s.getClient(), communityActionTemplate, new(model.CommunityActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.CommunityActionTemplate), nil
}

var _ ServiceInterface = &communityActionTemplateService{}
