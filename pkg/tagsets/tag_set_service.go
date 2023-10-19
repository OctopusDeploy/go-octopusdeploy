package tagsets

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type TagSetService struct {
	sortOrderPath string

	services.CanDeleteService
}

func NewTagSetService(sling *sling.Sling, uriTemplate string, sortOrderPath string) *TagSetService {
	return &TagSetService{
		sortOrderPath: sortOrderPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceTagSetService, sling, uriTemplate),
		},
	}
}

// Add creates a new tag set.
//
// Deprecated: Use tagsets.Add
func (s *TagSetService) Add(tagSet *TagSet) (*TagSet, error) {
	if IsNil(tagSet) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterTagSet)
	}

	path, err := services.GetAddPath(s, tagSet)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), tagSet, new(TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

// Get returns a collection of tag sets based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
//
// Deprecated: Use tagsets.Get
func (s *TagSetService) Get(tagSetsQuery TagSetsQuery) (*resources.Resources[*TagSet], error) {
	path, err := s.GetURITemplate().Expand(tagSetsQuery)
	if err != nil {
		return &resources.Resources[*TagSet]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*TagSet]), path)
	if err != nil {
		return &resources.Resources[*TagSet]{}, err
	}

	return response.(*resources.Resources[*TagSet]), nil
}

// GetByID returns the tag set that matches the input ID. If one cannot be
// found, it returns nil and an error.
//
// Deprecated: Use tagsets.GetByID
func (s *TagSetService) GetByID(id string) (*TagSet, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

// GetAll returns all tag sets. If none can be found or an error occurs, it
// returns an empty collection.
//
// Deprecates: use tagsets.GetAll
func (s *TagSetService) GetAll() ([]*TagSet, error) {
	items := []*TagSet{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the TagSet with a matching name.
func (s *TagSetService) GetByName(name string) (*TagSet, error) {
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
		if strings.EqualFold(item.Name, name) {
			return item, nil
		}
	}

	return nil, internal.CreateItemNotFoundError(s.GetName(), "GetByName", name)
}

// Update modifies a tag set based on the one provided as input.
//
// Deprecated: Use tagsets.Update
func (s *TagSetService) Update(tagSet *TagSet) (*TagSet, error) {
	if tagSet == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterTagSet)
	}

	path, err := services.GetUpdatePath(s, tagSet)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), tagSet, new(TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

// --- new ---

const template = "/api/{spaceId}/tagsets{/id}{?skip,take,ids,partialName}"

// Add creates a new tag set.
func Add(client newclient.Client, tagSet *TagSet) (*TagSet, error) {
	return newclient.Add[TagSet](client, template, tagSet.SpaceID, tagSet)
}

// Get returns a collection of tag sets based on the criteria defined by its
// input query parameter.
func Get(client newclient.Client, spaceID string, tagSetsQuery TagSetsQuery) (*resources.Resources[*TagSet], error) {
	return newclient.GetByQuery[TagSet](client, template, spaceID, tagSetsQuery)
}

// GetByID returns the tag set that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (*TagSet, error) {
	return newclient.GetByID[TagSet](client, template, spaceID, ID)
}

// Update modifies a tag set based on the one provided as input.
func Update(client newclient.Client, tagSet *TagSet) (*TagSet, error) {
	return newclient.Update[TagSet](client, template, tagSet.SpaceID, tagSet.ID, tagSet)
}

// DeleteByID deletes the tag set that matches the provided ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetAll returns all tag sets. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*TagSet, error) {
	return newclient.GetAll[TagSet](client, template, spaceID)
}
