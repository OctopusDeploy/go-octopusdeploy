package tagsets

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
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
func (s *TagSetService) Get(tagSetsQuery TagSetsQuery) (*TagSets, error) {
	path, err := s.GetURITemplate().Expand(tagSetsQuery)
	if err != nil {
		return &TagSets{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(TagSets), path)
	if err != nil {
		return &TagSets{}, err
	}

	return response.(*TagSets), nil
}

// GetByID returns the tag set that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *TagSetService) GetByID(id string) (*TagSet, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

// GetAll returns all tag sets. If none can be found or an error occurs, it
// returns an empty collection.
func (s *TagSetService) GetAll() ([]*TagSet, error) {
	items := []*TagSet{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
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
		if item.Name == name {
			return item, nil
		}
	}

	return nil, internal.CreateItemNotFoundError(s.GetName(), "GetByName", name)
}

// Update modifies a tag set based on the one provided as input.
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
