package octopusdeploy

import (
	"github.com/dghubble/sling"
)

// communityActionTemplateService handles communication with Account-related methods of the Octopus API.
type communityActionTemplateService struct {
	service
}

// newCommunityActionTemplateService returns an communityActionTemplateService with a preconfigured client.
func newCommunityActionTemplateService(sling *sling.Sling, uriTemplate string) *communityActionTemplateService {
	return &communityActionTemplateService{
		service: newService(serviceCommunityActionTemplateService, sling, uriTemplate, new(CommunityActionTemplate)),
	}
}

func (s communityActionTemplateService) getInstallationPath(resource CommunityActionTemplate) (string, error) {
	err := resource.Validate()
	if err != nil {
		return emptyString, createValidationFailureError(operationInstall, err)
	}

	err = validateInternalState(s)
	if err != nil {
		return emptyString, err
	}

	values := make(map[string]interface{})
	values[parameterID] = resource.GetID()

	path, err := s.getURITemplate().Expand(values)
	path = path + "/installation"

	return path, err
}

func (s communityActionTemplateService) getPagedResponse(path string) ([]*CommunityActionTemplate, error) {
	resources := []*CommunityActionTemplate{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(CommunityActionTemplates), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*CommunityActionTemplates)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// GetAll returns all community action templates. If none can be found or an
// error occurs, it returns an empty collection.
func (s communityActionTemplateService) GetAll() ([]*CommunityActionTemplate, error) {
	path, err := getPath(s)
	if err != nil {
		return []*CommunityActionTemplate{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the community action template that matches the input ID. If
// one cannot be found, it returns nil and an error.
func (s communityActionTemplateService) GetByID(id string) (*CommunityActionTemplate, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*CommunityActionTemplate), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s communityActionTemplateService) GetByIDs(ids []string) ([]*CommunityActionTemplate, error) {
	if len(ids) == 0 {
		return []*CommunityActionTemplate{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []*CommunityActionTemplate{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName performs a lookup and returns the CommunityActionTemplate with a matching name.
func (s communityActionTemplateService) GetByName(name string) (*CommunityActionTemplate, error) {
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

	for _, item := range collection {
		if item.Name == name {
			return item, nil
		}
	}

	return nil, createResourceNotFoundError(s.getName(), parameterName, name)
}

// Install installs a community step template.
func (s communityActionTemplateService) Install(resource CommunityActionTemplate) (*CommunityActionTemplate, error) {
	path, err := s.getInstallationPath(resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiPost(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*CommunityActionTemplate), nil
}

// Update modifies a community action template based on the one provided as
// input.
func (s communityActionTemplateService) Update(resource CommunityActionTemplate) (*CommunityActionTemplate, error) {
	path, err := s.getInstallationPath(resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*CommunityActionTemplate), nil
}
