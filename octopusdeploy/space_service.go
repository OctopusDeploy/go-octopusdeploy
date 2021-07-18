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
	spaceService.service = newService(ServiceSpaceService, sling, uriTemplate)

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
func (s spaceService) Add(space *Space) (*Space, error) {
	if space == nil {
		return nil, createInvalidParameterError("Add", "space")
	}

	path, err := getAddPath(s, space)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), space, new(Space), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}

// Get returns a collection of spaces based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s spaceService) Get(spacesQuery SpacesQuery) (*Spaces, error) {
	path, err := s.getURITemplate().Expand(spacesQuery)
	if err != nil {
		return &Spaces{}, err
	}

	response, err := apiGet(s.getClient(), new(Spaces), path)
	if err != nil {
		return &Spaces{}, err
	}

	return response.(*Spaces), nil
}

// GetByID returns the space that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s spaceService) GetByID(id string) (*Space, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Space), path)
	if err != nil {
		return nil, err
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
		return nil, createInvalidParameterError(OperationGetByName, ParameterName)
	}

	if err := validateInternalState(s); err != nil {
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

// GetByPartialName performs a lookup and returns spaces with a matching partial name.
func (s spaceService) GetByPartialName(name string) ([]*Space, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []*Space{}, err
	}

	return s.getPagedResponse(path)
}

// Update modifies a space based on the one provided as input.
func (s spaceService) Update(space *Space) (*Space, error) {
	if space == nil {
		return nil, createRequiredParameterIsEmptyOrNilError("space")
	}

	path, err := getUpdatePath(s, space)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), space, new(Space), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}
