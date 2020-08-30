package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type SpaceService struct {
	sling *sling.Sling
}

func NewSpaceService(sling *sling.Sling) *SpaceService {
	return &SpaceService{
		sling: sling,
	}
}

type Spaces struct {
	Items []Space `json:"Items"`
	PagedResults
}

type Space struct {
	ID                 string   `json:"Id"`
	Name               string   `json:"Name"`
	Description        string   `json:"Description"`
	IsDefault          bool     `json:"IsDefault"`
	TaskQueueStopped   bool     `json:"TaskQueueStopped"`
	SpaceManagersTeams []string `json:"SpaceManagersTeams"`
}

func (t *Space) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewSpace(name string) *Space {
	return &Space{
		Name: name,
	}
}

func (s *SpaceService) Get(spaceID string) (*Space, error) {
	path := fmt.Sprintf("spaces/%s", spaceID)
	resp, err := apiGet(s.sling, new(Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}

func (s *SpaceService) GetAll() (*[]Space, error) {
	var p []Space

	path := "spaces"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Spaces), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Spaces)

		p = append(p, r.Items...)

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *SpaceService) GetByName(spaceName string) (*Space, error) {
	var foundSpace Space
	spaces, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, space := range *spaces {
		if space.Name == spaceName {
			return &space, nil
		}
	}

	return &foundSpace, fmt.Errorf("no space found with space name %s", spaceName)
}

func (s *SpaceService) Add(space *Space) (*Space, error) {
	resp, err := apiAdd(s.sling, space, new(Space), "spaces")

	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}

func (s *SpaceService) Delete(spaceID string) error {
	path := fmt.Sprintf("spaces/%s", spaceID)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *SpaceService) Update(space *Space) (*Space, error) {
	path := fmt.Sprintf("spaces/%s", space.ID)
	resp, err := apiUpdate(s.sling, space, new(Space), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Space), nil
}
