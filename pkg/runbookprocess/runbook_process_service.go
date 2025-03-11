package runbookprocess

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

const template = "/api/{spaceId}/runbookProcesses{/id}{?skip,take,ids}"

// GetByID returns the runbook process that matches the input ID. If one cannot
// be found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*RunbookProcess, error) {
	return newclient.GetByID[RunbookProcess](client, template, spaceID, ID)
}

func GetGitRunbookProcessByID(client newclient.Client, spaceID string, projectID string, gitRef string, ID string) (*RunbookProcess, error) {

	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if gitRef == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("gitRef")
	}
	if ID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("ID")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "gitRef": gitRef, "id": ID}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookProcess, templateParams)
	if err != nil {
		return nil, err
	}
	runbook, err := newclient.Get[RunbookProcess](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return runbook, nil
}

// Update modifies a runbook process based on the one provided as input.
func Update(client newclient.Client, runbook *RunbookProcess) (*RunbookProcess, error) {
	return newclient.Update[RunbookProcess](client, template, runbook.SpaceID, runbook.ID, runbook)
}

func UpdateGitRunbook(client newclient.Client, runbookProcess *RunbookProcess, gitRef string) (*RunbookProcess, error) {

	if runbookProcess.SpaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceId")
	}

	templateParams := map[string]any{"spaceId": runbookProcess.SpaceID, "projectId": runbookProcess.ProjectID, "gitRef": gitRef, "id": runbookProcess.ID}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookProcess, templateParams)
	if err != nil {
		return nil, err
	}
	return newclient.Update[RunbookProcess](client, expandedUri, runbookProcess.SpaceID, runbookProcess.ID, runbookProcess)
}
