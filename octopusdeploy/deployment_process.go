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
	Items []DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	ID             string                   `json:"Id"`
	ProjectID      string                   `json:"ProjectId"`
	Steps          []DeploymentStepResource `json:"Steps"`
	Version        int                      `json:"Version"`
	LastSnapshotID string                   `json:"LastSnapshotId"`
	LastModifiedOn string                   `json:"LastModifiedOn"` // date time
	LastModifiedBy string                   `json:"LastModifiedBy"`
	Links          Links                    `json:"Links"`
}

func (d *DeploymentProcessService) Get(deploymentProcessId string) (DeploymentProcess, error) {
	deploymentProcess := new(DeploymentProcess)
	octopusDeployError := new(OctopusDeployError)
	path := fmt.Sprintf("api/deploymentprocesses/%s", deploymentProcessId)

	resp, err := d.sling.New().Get(path).Receive(deploymentProcess, octopusDeployError)

	if err != nil {
		return *deploymentProcess, fmt.Errorf("cannot get deploymentprocess id %s from server. failure from http client %v", deploymentProcessId, err)
	}

	if resp.StatusCode != http.StatusOK {
		return *deploymentProcess, fmt.Errorf("cannot get deploymentprocess id %s from server. response from server %s", deploymentProcessId, resp.Status)
	}

	return *deploymentProcess, err
}

func (d *DeploymentProcessService) GetAll() ([]DeploymentProcess, error) {
	var listOfDeloymentProcess []DeploymentProcess
	path := fmt.Sprintf("api/deploymentprocesses")

	for {
		deploymentProcesses := new(DeploymentProcesses)
		octopusDeployError := new(OctopusDeployError)

		resp, err := d.sling.New().Get(path).Receive(deploymentProcesses, octopusDeployError)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Response: %v", resp.Status)
		fmt.Printf("Total Results: %d", deploymentProcesses.NumberOfPages)

		for _, deploymentProcess := range deploymentProcesses.Items {
			listOfDeloymentProcess = append(listOfDeloymentProcess, deploymentProcess)
		}

		if deploymentProcesses.PagedResults.Links.PageNext != "" {
			fmt.Printf("More pages to go! Next link: %s", deploymentProcesses.PagedResults.Links.PageNext)
			path = deploymentProcesses.PagedResults.Links.PageNext
		} else {
			break
		}
	}

	return listOfDeloymentProcess, nil // no more pages to go through
}
