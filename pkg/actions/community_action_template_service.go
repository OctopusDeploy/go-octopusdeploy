package actions

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

// CommunityActionTemplateService handles communication with Account-related methods of the Octopus API.
type CommunityActionTemplateService struct {
	services.Service
}

// NewCommunityActionTemplateService returns an CommunityActionTemplateService with a preconfigured client.
func NewCommunityActionTemplateService(sling *sling.Sling, uriTemplate string) *CommunityActionTemplateService {
	return &CommunityActionTemplateService{
		Service: services.NewService(constants.ServiceCommunityActionTemplateService, sling, uriTemplate),
	}
}

func (s *CommunityActionTemplateService) getInstallationPath(resource CommunityActionTemplate) (string, error) {
	if err := resource.Validate(); err != nil {
		return "", internal.CreateValidationFailureError(constants.OperationInstall, err)
	}

	if err := services.ValidateInternalState(s); err != nil {
		return "", err
	}

	values := make(map[string]interface{})
	values[constants.ParameterID] = resource.GetID()

	path, err := s.GetURITemplate().Expand(values)
	path = path + "/installation"

	return path, err
}

func (s *CommunityActionTemplateService) getPagedResponse(path string) ([]*CommunityActionTemplate, error) {
	resources := []*CommunityActionTemplate{}
	loadNextPage := true

	for loadNextPage {
		resp, err := services.ApiGet(s.GetClient(), new(CommunityActionTemplates), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*CommunityActionTemplates)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = services.LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Get returns a collection of community action templates based on the criteria
// defined by its input query parameter. If an error occurs, an empty
// collection is returned along with the associated error.
func (s *CommunityActionTemplateService) Get(communityActionTemplatesQuery CommunityActionTemplatesQuery) (*CommunityActionTemplates, error) {
	v, _ := query.Values(communityActionTemplatesQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	resp, err := services.ApiGet(s.GetClient(), new(CommunityActionTemplates), path)
	if err != nil {
		return &CommunityActionTemplates{}, err
	}

	return resp.(*CommunityActionTemplates), nil
}

// GetAll returns all community action templates. If none can be found or an
// error occurs, it returns an empty collection.
func (s *CommunityActionTemplateService) GetAll() ([]*CommunityActionTemplate, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return []*CommunityActionTemplate{}, err
	}

	return s.getPagedResponse(path)
}

// GetByID returns the community action template that matches the input ID. If
// one cannot be found, it returns nil and an error.
func (s *CommunityActionTemplateService) GetByID(id string) (*CommunityActionTemplate, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(CommunityActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CommunityActionTemplate), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s *CommunityActionTemplateService) GetByIDs(ids []string) ([]*CommunityActionTemplate, error) {
	if len(ids) == 0 {
		return []*CommunityActionTemplate{}, nil
	}

	path, err := services.GetByIDsPath(s, ids)
	if err != nil {
		return []*CommunityActionTemplate{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName performs a lookup and returns the community action template with a
// matching name.
func (s *CommunityActionTemplateService) GetByName(name string) (*CommunityActionTemplate, error) {
	if internal.IsEmpty(name) {
		return nil, internal.CreateInvalidParameterError("GetByName", "name")
	}

	if err := services.ValidateInternalState(s); err != nil {
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

	return nil, internal.CreateResourceNotFoundError(s.GetName(), constants.ParameterName, name)
}

// Install installs a community step template.
func (s *CommunityActionTemplateService) Install(resource CommunityActionTemplate) (*CommunityActionTemplate, error) {
	path, err := s.getInstallationPath(resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiPost(s.GetClient(), resource, new(CommunityActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CommunityActionTemplate), nil
}

// Update modifies a community action template based on the one provided as
// input.
func (s *CommunityActionTemplateService) Update(resource CommunityActionTemplate) (*CommunityActionTemplate, error) {
	path, err := s.getInstallationPath(resource)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), resource, new(CommunityActionTemplate), path)
	if err != nil {
		return nil, err
	}

	return resp.(*CommunityActionTemplate), nil
}
