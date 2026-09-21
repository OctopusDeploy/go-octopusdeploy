package approvals

// ChangeRequestApprovalState is the state of the change carried by a server-task approval.
type ChangeRequestApprovalState string

const (
	ChangeRequestApprovalStatePreApproval  ChangeRequestApprovalState = "PreApproval"
	ChangeRequestApprovalStateApproved     ChangeRequestApprovalState = "Approved"
	ChangeRequestApprovalStatePostApproval ChangeRequestApprovalState = "PostApproval"
)

// ChangeRequest is the (optionally ITSM-backed) change associated with a
// server-task approval. Its Id is the key that groups server task approvals:
// deployments are gated per change request, and several server tasks can share
// one — for example the tenants of a release when the governing approval rule
// uses the PerRelease tenant approval strategy. Approval votes are tallied across
// every server task approval sharing an Id. There is no dedicated change request
// endpoint on the API; use Get or GetAll and group by this Id.
type ChangeRequest struct {
	Id                         string                     `json:"Id"`
	Number                     string                     `json:"Number,omitempty"`
	Description                string                     `json:"Description"`
	Active                     bool                       `json:"Active"`
	Type                       string                     `json:"Type,omitempty"`
	ChangeRequestApprovalState ChangeRequestApprovalState `json:"ChangeRequestApprovalState"`
	Self                       string                     `json:"Self,omitempty"`
}

// ServerTaskApproval is the approval requirement attached to a server task
// (list / get-by-id representation).
type ServerTaskApproval struct {
	Id                 string        `json:"Id"`
	ServerTaskId       string        `json:"ServerTaskId"`
	SpaceId            string        `json:"SpaceId"`
	ApprovalProviderId string        `json:"ApprovalProviderId"`
	Name               string        `json:"Name"`
	ChangeRequest      ChangeRequest `json:"ChangeRequest"`
}

// NamedReferenceItem is a lightweight id/display-name reference used for approving users/teams.
type NamedReferenceItem struct {
	Id               string `json:"Id"`
	DisplayName      string `json:"DisplayName"`
	DisplayIdAndName bool   `json:"DisplayIdAndName"`
}

// ServerTaskApprovalDetail is the richer representation returned by the by-task-id endpoint.
type ServerTaskApprovalDetail struct {
	Id                       string               `json:"Id"`
	SpaceId                  string               `json:"SpaceId"`
	Approvals                []Approval           `json:"Approvals"`
	ApprovalState            string               `json:"ApprovalState"`
	ApprovingUsers           []NamedReferenceItem `json:"ApprovingUsers,omitempty"`
	ApprovingTeams           []NamedReferenceItem `json:"ApprovingTeams,omitempty"`
	ApprovalsCount           int                  `json:"ApprovalsCount"`
	MinimumApproversRequired int                  `json:"MinimumApproversRequired"`
}
