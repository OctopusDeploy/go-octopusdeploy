package spaces

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type SpaceService struct {
	homePath string

	services.CanDeleteService
}

func NewSpaceService(sling *sling.Sling, uriTemplate string, homePath string) *SpaceService {
	return &SpaceService{
		homePath: homePath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceSpaceService, sling, uriTemplate),
		},
	}
}

// Add creates a new space.
func (s *SpaceService) Add(space *Space) (*Space, error) {
	if IsNil(space) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterSpace)
	}

	path, err := services.GetAddPath(s, space)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), space, new(Space), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}

// Get returns a collection of spaces based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
//
// Deprecated: Use spaces.Get
func (s *SpaceService) Get(spacesQuery SpacesQuery) (*resources.Resources[*Space], error) {
	path, err := s.GetURITemplate().Expand(spacesQuery)
	if err != nil {
		return &resources.Resources[*Space]{}, err
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*Space]), path)
	if err != nil {
		return &resources.Resources[*Space]{}, err
	}

	return response.(*resources.Resources[*Space]), nil
}

// GetByID returns the space that matches the input ID. If one cannot be found,
// it returns nil and an error.
//
// Deprecated: Use spaces.GetByID
func (s *SpaceService) GetByID(id string) (*Space, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Space), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}

// GetByName returns the space that matches the input ID or name. If one
// cannot be found, it returns nil and an error.
func (s *SpaceService) GetByName(name string) (*Space, error) {
	spaces, err := s.Get(SpacesQuery{
		PartialName: name,
	})
	if err != nil {
		return nil, err
	}

	for _, space := range spaces.Items {
		if strings.EqualFold(space.Name, name) {
			return space, nil
		}
	}

	return nil, services.ErrItemNotFound
}

// GetByIDOrName returns the space that matches the input ID or name. If one
// cannot be found, it returns nil and an error.
func (s *SpaceService) GetByIDOrName(idOrName string) (*Space, error) {
	space, err := s.GetByID(idOrName)
	if err != nil {
		apiError, ok := err.(*core.APIError)
		if ok && apiError.StatusCode != 404 {
			return nil, err
		}
	} else {
		if space != nil {
			return space, nil
		}
	}

	return s.GetByName(idOrName)
}

// GetAll returns all spaces. If none can be found or an error occurs, it
// returns an empty collection.
//
// Deprecated: Use spaces.GetAll
func (s *SpaceService) GetAll() ([]*Space, error) {
	items := []*Space{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// Update modifies a space based on the one provided as input.
//
// Deprecated: Use spaces.Update
func (s *SpaceService) Update(space *Space) (*Space, error) {
	if space == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("space")
	}

	path, err := services.GetUpdatePath(s, space)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), space, new(Space), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}

// --- new ---

const (
	spacesTemplate = "/api/spaces{/id}{?skip,ids,take,partialName}"
)

// Get returns a collection of spaces based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func Get(client newclient.Client, spacesQuery SpacesQuery) (*resources.Resources[*Space], error) {
	path, err := client.URITemplateCache().Expand(spacesTemplate, spacesQuery)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[resources.Resources[*Space]](client.HttpSession(), path)
	if err != nil {
		return &resources.Resources[*Space]{}, err
	}

	return res, nil
}

// GetAll returns all spaces. If none can be found or an error occurs, it
// returns an empty collection.
func GetAll(client newclient.Client) ([]*Space, error) {
	path, err := client.URITemplateCache().Expand(spacesTemplate, map[string]any{})
	if err != nil {
		return nil, err
	}
	spaces := make([]*Space, 0)
	res, err := newclient.Get[resources.Resources[*Space]](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}
	spaces = append(spaces, res.Items...)
	for res.Links.PageNext != "" {
		nextPagePath, err := client.URITemplateCache().Expand(res.Links.PageNext, map[string]any{})
		if err != nil {
			return nil, err
		}
		res, err = newclient.Get[resources.Resources[*Space]](client.HttpSession(), nextPagePath)
		if err != nil {
			return nil, err
		}
		spaces = append(spaces, res.Items...)
	}
	return spaces, err
}

// GetByID returns the space that matches the input ID. If one cannot be found,
// it returns nil and an error.
func GetByID(client newclient.Client, id string) (*Space, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateRequiredParameterIsEmptyError(constants.ParameterID)
	}

	path, err := client.URITemplateCache().Expand(spacesTemplate, map[string]any{
		"id": id,
	})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Get[Space](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Update modifies a space based on the one provided as input.
func Update(client newclient.Client, space *Space) (*Space, error) {
	if space == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("space")
	}

	path, err := client.URITemplateCache().Expand(spacesTemplate, map[string]any{
		"id": space.ID,
	})
	if err != nil {
		return nil, err
	}

	res, err := newclient.Put[Space](client.HttpSession(), path, space)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetDefaultSpace tries to find default space. Returns nil if a default space can not be found.
func GetDefaultSpace(client newclient.Client) (*Space, error) {
	// TODO: this should change to return a custom error (can't find default space)
	spaces, err := GetAll(client)
	if err != nil {
		return nil, err
	}

	for _, space := range spaces {
		if space.IsDefault {
			return space, nil
		}
	}
	return nil, nil
}
