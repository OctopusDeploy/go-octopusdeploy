package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type DeploymentProcessService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewDeploymentProcessService(sling *sling.Sling, uriTemplate string) *DeploymentProcessService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &DeploymentProcessService{
		name:  "DeploymentProcessService",
		path:  path,
		sling: sling,
	}
}

func (s *DeploymentProcessService) Get(id string) (*model.DeploymentProcess, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentProcess), nil
}

// GetAll returns all instances of a DeploymentProcess.
func (s *DeploymentProcessService) GetAll() (*[]model.DeploymentProcess, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

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

func (s *DeploymentProcessService) Update(deploymentProcess *model.DeploymentProcess) (*model.DeploymentProcess, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = deploymentProcess.Validate()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", deploymentProcess.ID)
	resp, err := apiUpdate(s.sling, deploymentProcess, new(model.DeploymentProcess), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.DeploymentProcess), nil
}

func (s *DeploymentProcessService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &DeploymentProcessService{}
