package scopeduserroles

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// --- new ---
const template = "/api/{spaceId}/scopeduserroles{/id}{?skip,take,ids,partialName,spaces,includeSystem}"

// Add creates a new scoped user role.
func Add(client newclient.Client, scopedUserRole *ScopedUserRole) (*ScopedUserRole, error) {
	return newclient.Add[ScopedUserRole](client, template, scopedUserRole.SpaceID, scopedUserRole)
}

// Get returns a collection of scoped user roles based on the criteria defined by
// its input query parameter.
func Get(client newclient.Client, spaceID string, scopedUserRolesQuery ScopedUserRolesQuery) (*resources.Resources[*ScopedUserRole], error) {
	return newclient.GetByQuery[ScopedUserRole](client, template, spaceID, scopedUserRolesQuery)
}

// GetByID returns the scoped user role that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (*ScopedUserRole, error) {
	return newclient.GetByID[ScopedUserRole](client, template, spaceID, ID)
}

// Update modifies a ScopedUserRole based on the one provided as input.
func Update(client newclient.Client, scopedUserRole *ScopedUserRole) (*ScopedUserRole, error) {
	return newclient.Update[ScopedUserRole](client, template, scopedUserRole.SpaceID, scopedUserRole.ID, scopedUserRole)
}

// DeleteByID deletes a scoped user role that matches the input ID
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}
