package runbooks

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
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
//
// Deprecated: use runbooks.Add
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
//
// Deprecated: use runbooks.GetByID
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
		return nil, err
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
//
// Deprecated: use runbooks.Update
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

const template = "/api/{spaceId}/runbooks{/id}{?skip,take,ids,partialName,clone,projectIds}"

// Add returns the runbook that matches the input ID.
func Add(client newclient.Client, runbook *Runbook) (*Runbook, error) {
	return newclient.Add[Runbook](client, template, runbook.SpaceID, runbook)
}

// GetByID returns the runbook that matches the input ID. If one cannot be
// found, it returns nil and an error.
func GetByID(client newclient.Client, spaceID string, ID string) (*Runbook, error) {
	return newclient.GetByID[Runbook](client, template, spaceID, ID)
}

// Update modifies a runbook based on the one provided as input.
func Update(client newclient.Client, runbook *Runbook) (*Runbook, error) {
	return newclient.Update[Runbook](client, template, runbook.SpaceID, runbook.ID, runbook)
}

func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}

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

// GetRunbookSnapshotRunPreview gets a preview of a snapshot run for a given environment.
// This is used by the portal to show which machines would be deployed to, and other information about the deployment,
// before proceeding with it. The CLI uses it to build the selector for picking specific machines to deploy to
func GetRunbookSnapshotRunPreview(client newclient.Client, spaceID string, snapshotID string, environmentID string, includeDisabledSteps bool) (*RunPreview, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetRunbookRunPreview", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetRunbookRunPreview", "spaceID")
	}
	if snapshotID == "" {
		return nil, internal.CreateInvalidParameterError("GetRunbookRunPreview", "snapshotID")
	}
	if environmentID == "" {
		return nil, internal.CreateInvalidParameterError("GetRunbookRunPreview", "environmentID")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.RunbookSnapshotRunPreview, map[string]any{
		"spaceId":              spaceID,
		"snapshotId":           snapshotID,
		"environmentId":        environmentID,
		"includeDisabledSteps": includeDisabledSteps,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[RunPreview](client.HttpSession(), expandedUri)
}

// TODO there is also a tenanted preview, request/response below.
// There's one preview per selected tenant, and the web portal resolves tenant tags down into specific tenants before
// calling the preview endpoint (you can't preview by tag).
// This is how it figures out what the deployment targets are.

// Unresolved: How do we want to resolve this in the CLI?
// Note: like deploy, the executions API doesn't support target machines per-tenant, only a single list.
//
// We could
// - A: resolve all the tags down into tenants, call the multi-preview endpoint, and union all the target envs together?
// - B: always use the untenanted preview endpoint (which is easier, but maybe not perfectly correct)
// [I've gone with B at the moment in the CLI]
//
// POST http://localhost:8050/api/Spaces-1/projects/Projects-561/runbooks/Runbooks-82/runbookRuns/previews
// {"DeploymentPreviews":[{"TenantId":"Tenants-41","EnvironmentId":"Environments-101"},{"TenantId":"Tenants-42","EnvironmentId":"Environments-101"}]}
// Response shape is like this:
/*
[
  {
    // preview for the first tenant/environment combo
  },
  {
     // preview for the second tenant/environment combo
  },
]
*/

// List returns a list of Git runbooks from the server, in a standard Octopus paginated result structure.
// If you don't specify --limit the server will use a default limit (typically 30)
func ListGitRunbooks(client newclient.Client, spaceID string, projectID string, gitRef string, filter string, limit int) (*resources.Resources[*Runbook], error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if gitRef == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("gitRef")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "gitRef": gitRef}
	if filter != "" {
		templateParams["partialName"] = filter
	}
	if limit > 0 {
		templateParams["take"] = limit
	}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbooksByProject, templateParams)
	if err != nil {
		return nil, err
	}

	return newclient.Get[resources.Resources[*Runbook]](client.HttpSession(), expandedUri)
}

// GetGitRunbookByID returns the runbook that matches the input ID and GitRef. If one cannot be
// found, it returns nil and an error.
func GetGitRunbookByID(client newclient.Client, spaceID string, projectID string, gitRef string, ID string) (*Runbook, error) {
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
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookById, templateParams)
	if err != nil {
		return nil, err
	}
	runbook, err := newclient.Get[Runbook](client.HttpSession(), expandedUri)
	if err != nil {
		return nil, err
	}

	return runbook, nil
}

// GetGitRunbookByName searches for a single runbook with name of 'name'.
// If no such runbook can be found, will return nil, nil
func GetGitRunbookByName(client newclient.Client, spaceID string, projectID string, gitRef string, name string) (*Runbook, error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if gitRef == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("gitRef")
	}
	if name == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("name")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "gitRef": gitRef, "partialName": name}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbooksByProject, templateParams)
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

// AddGitRunbook updates the runbook that matches the input ID and GitRef.
func AddGitRunbook(client newclient.Client, runbook *Runbook, gitRef string) (*Runbook, error) {
	if gitRef == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("gitRef")
	}

	if runbook.ProjectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}

	if runbook.SpaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}

	templateParams := map[string]any{"spaceId": runbook.SpaceID, "projectId": runbook.ProjectID, "gitRef": gitRef, "id": runbook.ID}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbooks, templateParams)
	if err != nil {
		return nil, err
	}
	partialRun, creationError := newclient.Add[CreatedRunbook](client, expandedUri, runbook.SpaceID, runbook)

	if creationError != nil {
		return nil, creationError
	}

	return GetGitRunbookByID(client, runbook.SpaceID, partialRun.ProjectID, partialRun.GitRef, partialRun.Slug)
}

// UpdateGitRunbook updates the runbook that matches the input ID and GitRef.
func UpdateGitRunbook(client newclient.Client, runbook *Runbook, gitRef string) (*Runbook, error) {
	if gitRef == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("gitRef")
	}
	templateParams := map[string]any{"spaceId": runbook.SpaceID, "projectId": runbook.ProjectID, "gitRef": gitRef, "id": runbook.ID}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookById, templateParams)
	if err != nil {
		return nil, err
	}
	return newclient.Update[Runbook](client, expandedUri, runbook.SpaceID, runbook.ID, runbook)
}

// DeleteGitRunbook deletes the runbook that matches the input ID and GitRef.
func DeleteGitRunbook(client newclient.Client, spaceID string, projectID string, gitRef string, ID string) error {
	if spaceID == "" {
		return internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if gitRef == "" {
		return internal.CreateRequiredParameterIsEmptyOrNilError("gitRef")
	}
	if ID == "" {
		return internal.CreateRequiredParameterIsEmptyOrNilError("ID")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "gitRef": gitRef, "id": ID}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookById, templateParams)
	if err != nil {
		return err
	}
	err = newclient.Delete(client.HttpSession(), expandedUri)
	if err != nil {
		return err
	}

	return nil
}

// ListEnvironmentsForGitRunbook returns the list of valid environments for a given runbook stored in Git
func ListEnvironmentsForGitRunbook(client newclient.Client, spaceID string, projectID string, runbookID string, gitRef string) ([]*environments.Environment, error) {
	if spaceID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("projectID")
	}
	if runbookID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("runbookID")
	}
	if gitRef == "" {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("gitRef")
	}
	templateParams := map[string]any{"spaceId": spaceID, "projectId": projectID, "runbookId": runbookID, "gitRef": gitRef}
	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookEnvironments, templateParams)
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

func GetGitRunbookProcess(client newclient.Client, spaceID string, projectID string, runbookID string, gitRef string) (*RunbookProcess, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookProcess", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookProcess", "spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookProcess", "projectID")
	}
	if runbookID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookProcess", "runbookID")
	}
	if gitRef == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookProcess", "gitRef")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookProcess, map[string]any{
		"spaceId":   spaceID,
		"projectId": projectID,
		"gitRef":    gitRef,
		"id":        runbookID,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[RunbookProcess](client.HttpSession(), expandedUri)
}

// GetGitRunbookRunPreview gets a preview of a run for a given environment for a runbook stored in Git.
// This is used by the portal to show which machines would be deployed to, and other information about the deployment,
// before proceeding with it. The CLI uses it to build the selector for picking specific machines to deploy to
func GetGitRunbookRunPreview(client newclient.Client, spaceID string, projectID string, runbookID string, gitRef string, environmentID string, includeDisabledSteps bool) (*RunPreview, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "projectID")
	}
	if runbookID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "runbookID")
	}
	if gitRef == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "gitRef")
	}
	if environmentID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "environmentID")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookRunPreview, map[string]any{
		"spaceId":              spaceID,
		"projectId":            projectID,
		"runbookId":            runbookID,
		"gitRef":               gitRef,
		"environment":          environmentID,
		"includeDisabledSteps": includeDisabledSteps,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[RunPreview](client.HttpSession(), expandedUri)
}

func GetGitRunbookSnapshotTemplate(client newclient.Client, spaceID string, projectID string, runbookID string, gitRef string) (*RunbookSnapshotTemplate, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "spaceID")
	}
	if projectID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "projectID")
	}
	if runbookID == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "runbookID")
	}
	if gitRef == "" {
		return nil, internal.CreateInvalidParameterError("GetGitRunbookRunPreview", "gitRef")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.GitRunbookSnapshotTemplate, map[string]any{
		"spaceId":   spaceID,
		"projectId": projectID,
		"runbookId": runbookID,
		"gitRef":    gitRef,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[RunbookSnapshotTemplate](client.HttpSession(), expandedUri)
}
