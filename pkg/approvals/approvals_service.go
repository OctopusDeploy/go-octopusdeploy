package approvals

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

const serverTaskApprovalsTemplate = "/api/{spaceId}/servertaskapprovals{/id}{?skip,take,providerId,state}"
const serverTaskApprovalByTaskTemplate = "/api/{spaceId}/tasks/{taskId}/servertaskapproval"
const approvalsTemplate = "/api/{spaceId}/servertaskapprovals/{serverTaskApprovalId}/approvals{/id}"

// ServerTaskApprovalsQuery is the query for the paginated server-task approvals list endpoint.
type ServerTaskApprovalsQuery struct {
	ProviderId string `uri:"providerId,omitempty" url:"providerId,omitempty"`
	State      string `uri:"state,omitempty" url:"state,omitempty"`
	Skip       int    `uri:"skip" url:"skip"`
	Take       int    `uri:"take,omitempty" url:"take,omitempty"`
}

// getServerTaskApprovalByTaskIdResponse is the nullable envelope for the by-task endpoint.
type getServerTaskApprovalByTaskIdResponse struct {
	Resource *ServerTaskApprovalDetail `json:"Resource,omitempty"`
}

// GetByTaskID returns the approval detail for a server task. Returns (nil, nil)
// when the task has no approval requirement (approvals disabled, not a
// deployment/runbook, or no approval rule applies).
//
// Approvals and ApprovalsCount are tallied across the whole change request, not
// just this server task: they include votes cast against sibling server task
// approvals that share the same ChangeRequest.Id, deduplicated to one vote per
// user with a rejection taking precedence.
func GetByTaskID(client newclient.Client, spaceID string, serverTaskID string) (*ServerTaskApprovalDetail, error) {
	if serverTaskID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("serverTaskID")
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}
	path, err := client.URITemplateCache().Expand(serverTaskApprovalByTaskTemplate, map[string]any{
		"spaceId": spaceID,
		"taskId":  serverTaskID,
	})
	if err != nil {
		return nil, err
	}
	res, err := newclient.Get[getServerTaskApprovalByTaskIdResponse](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}
	return res.Resource, nil
}

// GetByID returns the server-task approval that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (*ServerTaskApproval, error) {
	if ID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("ID")
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}
	path, err := client.URITemplateCache().Expand(serverTaskApprovalsTemplate, map[string]any{
		"spaceId": spaceID,
		"id":      ID,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[ServerTaskApproval](client.HttpSession(), path)
}

// Get returns a paginated collection of server-task approvals matching the query.
func Get(client newclient.Client, spaceID string, query ServerTaskApprovalsQuery) (*resources.Resources[*ServerTaskApproval], error) {
	return newclient.GetByQuery[ServerTaskApproval](client, serverTaskApprovalsTemplate, spaceID, query)
}

// GetAll returns all server-task approvals in the space, following pagination links.
func GetAll(client newclient.Client, spaceID string) ([]*ServerTaskApproval, error) {
	return newclient.GetAll[ServerTaskApproval](client, serverTaskApprovalsTemplate, spaceID)
}

// ListApprovals returns the individual decisions recorded for the change request
// that the given server-task approval belongs to. The server returns a bare array
// (not a paginated collection).
//
// The result is scoped to the change request rather than to serverTaskApprovalID:
// entries may carry a different ServerTaskApprovalId when sibling server tasks
// share the same ChangeRequest.Id. Votes are deduplicated to one per user, with a
// rejection taking precedence over an approval.
func ListApprovals(client newclient.Client, spaceID string, serverTaskApprovalID string) ([]*Approval, error) {
	if serverTaskApprovalID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("serverTaskApprovalID")
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}
	path, err := client.URITemplateCache().Expand(approvalsTemplate, map[string]any{
		"spaceId":              spaceID,
		"serverTaskApprovalId": serverTaskApprovalID,
	})
	if err != nil {
		return nil, err
	}
	res, err := newclient.Get[[]*Approval](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}
	return *res, nil
}

// GetApprovalByID returns a single decision by ID.
func GetApprovalByID(client newclient.Client, spaceID string, serverTaskApprovalID string, ID string) (*Approval, error) {
	if serverTaskApprovalID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("serverTaskApprovalID")
	}
	if ID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("ID")
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}
	path, err := client.URITemplateCache().Expand(approvalsTemplate, map[string]any{
		"spaceId":              spaceID,
		"serverTaskApprovalId": serverTaskApprovalID,
		"id":                   ID,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[Approval](client.HttpSession(), path)
}

// AddApproval records a new decision (approve/reject) against a server-task approval.
//
// A user may vote only once per change request, so this fails when the user has
// already voted against any server task approval sharing the change request. It
// also fails when the change is no longer active or has left the PreApproval
// state, or when the approval rule snapshot that governed the change request is
// no longer available.
func AddApproval(client newclient.Client, spaceID string, serverTaskApprovalID string, approval *Approval) (*Approval, error) {
	if approval == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("approval")
	}
	if serverTaskApprovalID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("serverTaskApprovalID")
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}
	path, err := client.URITemplateCache().Expand(approvalsTemplate, map[string]any{
		"spaceId":              spaceID,
		"serverTaskApprovalId": serverTaskApprovalID,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Post[Approval](client.HttpSession(), path, approval)
}

// UpdateApproval modifies an existing decision. It fails once the change has
// completed, that is when the change is inactive or has left the PreApproval state.
func UpdateApproval(client newclient.Client, spaceID string, serverTaskApprovalID string, approval *Approval) (*Approval, error) {
	if approval == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("approval")
	}
	if serverTaskApprovalID == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("serverTaskApprovalID")
	}
	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}
	path, err := client.URITemplateCache().Expand(approvalsTemplate, map[string]any{
		"spaceId":              spaceID,
		"serverTaskApprovalId": serverTaskApprovalID,
		"id":                   approval.Id,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Put[Approval](client.HttpSession(), path, approval)
}
