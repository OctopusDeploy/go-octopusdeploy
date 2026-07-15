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

// getServerTaskApprovalByIdResponse unwraps the nested resource for get-by-id.
type getServerTaskApprovalByIdResponse struct {
	ServerTaskApproval *ServerTaskApproval `json:"ServerTaskApproval,omitempty"`
}

// GetByTaskID returns the approval detail for a server task. Returns (nil, nil)
// when the task has no approval requirement (feature off, not a deployment/runbook,
// or no policy applies).
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
// It returns (nil, nil) if the response contains no server-task approval.
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
	res, err := newclient.Get[getServerTaskApprovalByIdResponse](client.HttpSession(), path)
	if err != nil {
		return nil, err
	}
	return res.ServerTaskApproval, nil
}

// Get returns a paginated collection of server-task approvals matching the query.
func Get(client newclient.Client, spaceID string, query ServerTaskApprovalsQuery) (*resources.Resources[*ServerTaskApproval], error) {
	return newclient.GetByQuery[ServerTaskApproval](client, serverTaskApprovalsTemplate, spaceID, query)
}

// GetAll returns all server-task approvals in the space, following pagination links.
func GetAll(client newclient.Client, spaceID string) ([]*ServerTaskApproval, error) {
	return newclient.GetAll[ServerTaskApproval](client, serverTaskApprovalsTemplate, spaceID)
}

// ListApprovals returns the individual decisions recorded against a server-task approval.
// The server returns a bare array (not a paginated collection).
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

// UpdateApproval modifies an existing decision.
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
