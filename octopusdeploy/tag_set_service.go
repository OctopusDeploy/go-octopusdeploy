package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type tagSetService struct {
	sortOrderPath string

	canDeleteService
}

func newTagSetService(sling *sling.Sling, uriTemplate string, sortOrderPath string) *tagSetService {
	tagSetService := &tagSetService{
		sortOrderPath: sortOrderPath,
	}
	tagSetService.service = newService(ServiceTagSetService, sling, uriTemplate)

	return tagSetService
}

// Add creates a new tag set.
func (s tagSetService) Add(resource *TagSet) (*TagSet, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

// Get returns a collection of tag sets based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s tagSetService) Get(tagSetsQuery TagSetsQuery) (*TagSets, error) {
	path, err := s.getURITemplate().Expand(tagSetsQuery)
	if err != nil {
		return &TagSets{}, err
	}

	response, err := apiGet(s.getClient(), new(TagSets), path)
	if err != nil {
		return &TagSets{}, err
	}

	return response.(*TagSets), nil
}

// GetByID returns the tag set that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s tagSetService) GetByID(id string) (*TagSet, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}

// GetAll returns all tag sets. If none can be found or an error occurs, it
// returns an empty collection.
func (s tagSetService) GetAll() ([]*TagSet, error) {
	items := []*TagSet{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the TagSet with a matching name.
func (s tagSetService) GetByName(name string) (*TagSet, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError(OperationGetByName, ParameterName)
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

	return nil, createItemNotFoundError(s.getName(), OperationGetByName, name)
}

// Update modifies a tag set based on the one provided as input.
func (s tagSetService) Update(tagSet *TagSet) (*TagSet, error) {
	if tagSet == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterTagSet)
	}

	path, err := getUpdatePath(s, tagSet)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), tagSet, new(TagSet), path)
	if err != nil {
		return nil, err
	}

	return resp.(*TagSet), nil
}
