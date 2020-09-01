package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type DeploymentProcessService struct {
	sling *sling.Sling
	path  string
}

func NewDeploymentProcessService(sling *sling.Sling) *DeploymentProcessService {
	return &DeploymentProcessService{
		sling: sling,
		path:  "deploymentprocesses",
	}
}

func (s *DeploymentProcessService) Get(id string) (*model.DeploymentProcess, error) {
	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentProcess), nil
}

func (s *DeploymentProcessService) GetAll() (*[]model.DeploymentProcess, error) {
	var p []model.DeploymentProcess
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.DeploymentProcesses), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.DeploymentProcesses)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *DeploymentProcessService) Update(resource *model.DeploymentProcess) (*model.DeploymentProcess, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentProcess), nil
}
