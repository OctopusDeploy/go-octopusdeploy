package runbookprocess

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const template = "/api/{spaceId}/runbookProcesses{/id}{?skip,take,ids}"

// GetByID returns the runbook process that matches the input ID. If one cannot
// be found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*RunbookProcess, error) {
	return newclient.GetByID[RunbookProcess](client, template, spaceID, ID)
}

// Update modifies a runbook process based on the one provided as input.
func Update(client newclient.Client, runbook *RunbookProcess) (*RunbookProcess, error) {
	return newclient.Update[RunbookProcess](client, template, runbook.SpaceID, runbook.ID, runbook)
}
