package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type MachinePolicyService struct {
	sling *sling.Sling
	path  string
}

func NewMachinePolicyService(sling *sling.Sling) *MachinePolicyService {
	return &MachinePolicyService{
		sling: sling,
		path:  "machinepolicies",
	}
}

// Get returns a single machine with a given MachineID
func (s *MachinePolicyService) Get(machinePolicyID string) (*model.MachinePolicy, error) {
	path := fmt.Sprintf(s.path+"/%s", machinePolicyID)
	resp, err := apiGet(s.sling, new(model.Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.MachinePolicy), nil
}

// GetAll returns all registered machines
func (s *MachinePolicyService) GetAll() (*[]model.MachinePolicy, error) {
	var p []model.MachinePolicy
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.MachinePolicies), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.MachinePolicies)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}
	return &p, nil
}
