package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type DeploymentProcessService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

func NewDeploymentProcessService(sling *sling.Sling) *DeploymentProcessService {
	if sling == nil {
		return nil
	}

	return &DeploymentProcessService{
		sling: sling,
		path:  "deploymentprocesses",
	}
}

func (s *DeploymentProcessService) Get(id string) (*model.DeploymentProcess, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	if isEmpty(id) {
		return nil, errors.New("DeploymentProcessService: invalid parameter, id")
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
		return fmt.Errorf("DeploymentProcessService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("DeploymentProcessService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &DeploymentProcessService{}
