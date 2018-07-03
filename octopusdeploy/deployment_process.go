package octopusdeploy

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

type DeploymentProcessService struct {
	sling *sling.Sling
}

func NewDeploymentProcessService(sling *sling.Sling) *DeploymentProcessService {
	return &DeploymentProcessService{
		sling: sling,
	}
}

type DeploymentProcesses struct {
	Items []DeploymentProcessResource `json:"Items"`
	PagedResults
}

type DeploymentProcessResource struct {

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn string `json:"LastModifiedOn,omitempty"`

	// last snapshot Id
	LastSnapshotID string `json:"LastSnapshotId,omitempty"`

	// links
	Links Links `json:"Links,omitempty"`

	// project Id
	ProjectID string `json:"ProjectId,omitempty"`

	// steps
	Steps []DeploymentStepResource `json:"Steps"`

	// version
	// Required: true
	Version *int32 `json:"Version"`
}

func (d *DeploymentProcessService) Get(deploymentProcessID string) (*DeploymentProcessResource, error) {
	var deploymentProcess DeploymentProcessResource
	octopusDeployError := new(APIError)
	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcessID)

	resp, err := d.sling.New().Get(path).Receive(&deploymentProcess, &octopusDeployError)

	if err != nil {
		return nil, fmt.Errorf("cannot get deploymentprocess id %s from server. failure from http client %v", deploymentProcessID, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrItemNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot get project id %s from server. response from server %s", deploymentProcessID, resp.Status)
	}

	return &deploymentProcess, err
}

func (d *DeploymentProcessService) GetAll() (*[]DeploymentProcessResource, error) {
	var listOfDeploymentProcess []DeploymentProcessResource
	path := fmt.Sprintf("deploymentprocesses")

	for {
		var deploymentProcesses DeploymentProcesses
		var octopusDeployError APIError

		resp, err := d.sling.New().Get(path).Receive(&deploymentProcesses, &octopusDeployError)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		if octopusDeployError.Errors != nil {
			return nil, fmt.Errorf("cannot get all deployment processes. response from octopusdeploy %s: ", octopusDeployError.Errors)
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("cannot get all projects. response from server %s", resp.Status)
		}

		for _, deploymentProcess := range deploymentProcesses.Items {
			listOfDeploymentProcess = append(listOfDeploymentProcess, deploymentProcess)
		}

		if deploymentProcesses.PagedResults.Links.PageNext != "" {
			path = deploymentProcesses.PagedResults.Links.PageNext
		} else {
			break
		}
	}

	return &listOfDeploymentProcess, nil // no more pages to go through
}

func (s *DeploymentProcessService) Update(deploymentProcess *DeploymentProcessResource) (*DeploymentProcessResource, error) {
	var updated DeploymentProcessResource
	var octopusDeployError APIError

	path := fmt.Sprintf("deploymentprocesses/%s", deploymentProcess.ID)
	resp, err := s.sling.New().Put(path).BodyJSON(deploymentProcess).Receive(&updated, &octopusDeployError)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if octopusDeployError.Errors != nil {
		return nil, fmt.Errorf("cannot update deployment process. response from octopusdeploy %s: ", octopusDeployError.Errors)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot update deployment process at url %s. response from server %s", resp.Request.URL, resp.Status)
	}

	return &updated, nil
}
