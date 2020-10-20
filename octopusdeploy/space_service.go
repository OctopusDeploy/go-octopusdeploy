package octopusdeploy

import (
	"github.com/dghubble/sling"
)

type spaceService struct {
	homePath string

	canDeleteService
}

func newSpaceService(sling *sling.Sling, uriTemplate string, homePath string) *spaceService {
	spaceService := &spaceService{
		homePath: homePath,
	}
	spaceService.service = newService(serviceSpaceService, sling, uriTemplate, new(Space))

	return spaceService
}

func (s spaceService) getPagedResponse(path string) ([]*Space, error) {
	resources := []*Space{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Spaces), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Spaces)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

// Add creates a new space.
func (s spaceService) Add(resource *Space) (*Space, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}

// GetByID returns the space that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s spaceService) GetByID(id string) (*Space, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(*Space), nil
}

// GetAll returns all spaces. If none can be found or an error occurs, it
// returns an empty collection.
func (s spaceService) GetAll() ([]*Space, error) {
	items := []*Space{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByName performs a lookup and returns the Space with a matching name.
func (s spaceService) GetByName(name string) (*Space, error) {
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

	return nil, createItemNotFoundError(s.getName(), operationGetByName, name)
}

// GetByPartialName performs a lookup and returns spaces with a matching partial name.
func (s spaceService) GetByPartialName(name string) ([]*Space, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Space{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a space based on the one provided as input.
func (s spaceService) Update(resource *Space) (*Space, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}
