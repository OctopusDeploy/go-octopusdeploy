package machinepolicies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

const template = "/api/{spaceId}/machinepolicies{/id}{?skip,take,ids,partialName}"

// Get returns a collection of machine policies based on the criteria defined
// by its input query parameter. If an error occurs, an empty collection is
// returned along with the associated error.
func Get(client newclient.Client, spaceID string, machinePoliciesQuery MachinePoliciesQuery) (*resources.Resources[*MachinePolicy], error) {
	return newclient.GetByQuery[MachinePolicy](client, template, spaceID, machinePoliciesQuery)
}

// Add creates a new machine policy.
func Add(client newclient.Client, machinePolicy *MachinePolicy) (*MachinePolicy, error) {
	return newclient.Add[MachinePolicy](client, template, machinePolicy.SpaceID, machinePolicy)
}

// GetByID returns the machine policy that matches the input ID. If one cannot
// be found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, id string) (*MachinePolicy, error) {
	return newclient.GetByID[MachinePolicy](client, template, spaceID, id)
}

// Update modifies a machine policy based on the one provided as input.
func Update(client newclient.Client, machinePolicy *MachinePolicy) (*MachinePolicy, error) {
	return newclient.Update[MachinePolicy](client, template, machinePolicy.SpaceID, machinePolicy.ID, machinePolicy)
}

// DeleteByID deletes a machine policy based on the provided ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetAll returns all machine policies. If none can be found or an error
// occurs, it returns an empty collection.
func GetAll(client newclient.Client, spaceID string) (*[]MachinePolicy, error) {
	return newclient.GetAll[MachinePolicy](client, template, spaceID)
}
