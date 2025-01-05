package workers

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

const template = "/api/{spaceId}/workers{/id}{?skip,take,name,ids,partialName,isDisabled,healthStatuses,commStyles,shellNames}"

// Add creates a new worker.
func Add(client newclient.Client, worker *machines.Worker) (*machines.Worker, error) {
	return newclient.Add[machines.Worker](client, template, worker.SpaceID, worker)
}

// Get returns a collection of workers based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func Get(client newclient.Client, spaceID string, workersQuery machines.WorkersQuery) (*resources.Resources[*machines.Worker], error) {
	return newclient.GetByQuery[machines.Worker](client, template, spaceID, workersQuery)
}

// GetAll returns all workers. If an error occurs, it returns nil.
func GetAll(client newclient.Client, spaceID string) ([]*machines.Worker, error) {
	return newclient.GetAll[machines.Worker](client, template, spaceID)
}

// GetByID returns the worker that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*machines.Worker, error) {
	return newclient.GetByID[machines.Worker](client, template, spaceID, ID)
}

// Update updates an existing machine in Octopus Deploy
func Update(client newclient.Client, worker *machines.Worker) (*machines.Worker, error) {
	return newclient.Update[machines.Worker](client, template, worker.SpaceID, worker.ID, worker)
}

// DeleteByID deletes the worker based on the ID provided.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}
