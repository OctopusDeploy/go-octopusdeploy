package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"strings"
)

type RunbookService struct {
	services.CanDeleteService
}

func NewRunbookService(sling *sling.Sling, uriTemplate string) *RunbookService {
	return &RunbookService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceRunbookService, sling, uriTemplate),
		},
	}
}

// Add returns the runbook that matches the input ID.
func (s *RunbookService) Add(runbook *Runbook) (*Runbook, error) {
	if IsNil(runbook) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterRunbook)
	}

	if err := runbook.Validate(); err != nil {
		return nil, internal.CreateValidationFailureError(constants.OperationAdd, err)
	}

	path, err := services.GetAddPath(s, runbook)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), runbook, new(Runbook), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Runbook), nil
}

// GetAll returns all runbooks. If none can be found or an error occurs, it
// returns an empty collection.
func (s *RunbookService) GetAll() ([]*Runbook, error) {
	items := []*Runbook{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the runbook that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *RunbookService) GetByID(id string) (*Runbook, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Runbook), path)
	if err != nil {
		return nil, internal.CreateResourceNotFoundError("runbook", "ID", id)
	}

	return resp.(*Runbook), nil
}

func (s *RunbookService) GetRunbookSnapshotTemplate(runbook *Runbook) (*RunbookSnapshotTemplate, error) {
	resp, err := api.ApiGet(s.GetClient(), new(RunbookSnapshotTemplate), runbook.Links["RunbookSnapshotTemplate"])
	if err != nil {
		return nil, err
	}

	return resp.(*RunbookSnapshotTemplate), nil
}

// Update modifies a runbook based on the one provided as input.
func (s *RunbookService) Update(runbook *Runbook) (*Runbook, error) {
	if runbook == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterRunbook)
	}

	path, err := services.GetUpdatePath(s, runbook)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), runbook, new(Runbook), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Runbook), nil
}

// ---------------------------

// List returns a list of runbooks from the server, in a standard Octopus paginated result structure.
// If you don't specify --limit the server will use a default limit (typically 30)
func List(client newclient.Client, spaceID string, projectID string, filter string, limit int) (*resources.Resources[*Runbook], error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID}
	if filter != "" {
		templateParams["partialName"] = filter
	}
	if limit > 0 {
		templateParams["take"] = limit
	}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbooksByProject, templateParams)
	if err != nil {
		return nil, err
	}

	return newclient.Get[resources.Resources[*Runbook]](client.HttpSession(), expandedUri)
}

// GetByName searches for a single runbook with name of 'name'.
// If no such runbook can be found, will return nil, nil
func GetByName(client newclient.Client, spaceID string, projectID string, name string) (*Runbook, error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if name == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "partialName": name}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbooksByProject, templateParams)
	if err != nil {
		return nil, err
	}

	searchResults, err := newclient.Get[resources.Resources[*Runbook]](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}
	for _, item := range searchResults.Items {
		if strings.EqualFold(name, item.Name) {
			return item, nil
		}
	}
	return nil, nil
}

// ListSnapshots returns a list of runbook snapshots from the server, in a standard Octopus paginated result structure.
// If you don't specify --limit the server will use a default limit (typically 30)
func ListSnapshots(client newclient.Client, spaceID string, projectID string, runbookID string, limit int) (*resources.Resources[*RunbookSnapshot], error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if runbookID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("runbookID")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "runbookId": runbookID}
	if limit > 0 {
		templateParams["take"] = limit
	}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbookSnapshotsByRunbook, templateParams)
	if err != nil {
		return nil, err
	}

	return newclient.Get[resources.Resources[*RunbookSnapshot]](client.HttpSession(), expandedUri)
}

// GetSnapshot loads a single runbook snapshot.
// You can supply either a name "Snapshot FWKMLUX" or an ID "RunbookSnapshots-41" for snapshotIDorName
func GetSnapshot(client newclient.Client, spaceID string, projectID string, snapshotIDorName string) (*RunbookSnapshot, error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if snapshotIDorName == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("snapshotIDorName")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "name": snapshotIDorName}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbookSnapshotsByProject, templateParams)
	if err != nil {
		return nil, err
	}

	return newclient.Get[RunbookSnapshot](client.HttpSession(), expandedUri)
}

// ListEnvironments returns the list of valid environments for a given runbook
func ListEnvironments(client newclient.Client, spaceID string, projectID string, runbookID string) ([]*environments.Environment, error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if runbookID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("runbookID")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "runbookId": runbookID}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbookEnvironments, templateParams)
	if err != nil {
		return nil, err
	}

	// our generic Get method must return pointers, so we need to dereference the pointer-to-slice before returning it
	tmp, err := newclient.Get[[]*environments.Environment](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}
	return *tmp, nil
}

// GetProcess fetches a runbook process. This may either be the project level process (template),
// or a snapshot, depending on the value of ID
func GetProcess(client newclient.Client, spaceID string, projectID string, ID string) (*RunbookProcess, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetProcess", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetProcess", "spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateInvalidParameterError("GetProcess", "projectID")
	}
	if ID == "" {
		return nil, internal.CreateInvalidParameterError("GetProcess", "ID")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbookProcess, map[string]any{
		"spaceId":   spaceID,
		"projectId": projectID,
		"id":        ID,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[RunbookProcess](client.HttpSession(), expandedUri)
}
