package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

type ProjectGroupService struct {
	sling *sling.Sling
	path  string
}

func NewProjectGroupService(sling *sling.Sling) *ProjectGroupService {
	return &ProjectGroupService{
		sling: sling,
		path:  "projectgroups",
	}
}

func (s *ProjectGroupService) Get(id string) (*model.ProjectGroup, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

func (s *ProjectGroupService) GetAll() (*[]model.ProjectGroup, error) {
	var p []model.ProjectGroup
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.ProjectGroups), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.ProjectGroups)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *ProjectGroupService) Add(resource *model.ProjectGroup) (*model.ProjectGroup, error) {
	resp, err := apiAdd(s.sling, resource, new(model.ProjectGroup), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}

func (s *ProjectGroupService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *ProjectGroupService) Update(resource *model.ProjectGroup) (*model.ProjectGroup, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.ProjectGroup), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.ProjectGroup), nil
}
