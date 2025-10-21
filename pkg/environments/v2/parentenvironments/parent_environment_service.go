package parentenvironments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/spaces/{spaceId}/parentEnvironments{/id}"

// Add creates a new parent environment.
func Add(client newclient.Client, parentEnvironment *ParentEnvironment) (*ParentEnvironment, error) {
	return newclient.Add[ParentEnvironment](client, template, parentEnvironment.SpaceID, parentEnvironment)
}

// DeleteByID deletes the parent environment based on the ID provided as input.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

// GetByID returns the parent environment that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*ParentEnvironment, error) {
	return newclient.GetByID[ParentEnvironment](client, template, spaceID, ID)
}

// Update updates an existing parent environment.
func Update(client newclient.Client, parentEnvironment *ParentEnvironment) (*ParentEnvironment, error) {
	return newclient.Update[ParentEnvironment](client, template, parentEnvironment.SpaceID, parentEnvironment.ID, parentEnvironment)
}
