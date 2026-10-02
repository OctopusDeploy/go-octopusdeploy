package approvals

// ApprovalStatus is an individual approver's decision.
type ApprovalStatus string

const (
	ApprovalStatusApproved ApprovalStatus = "Approved"
	ApprovalStatusRejected ApprovalStatus = "Rejected"
)

// Approval is a single person's decision recorded against a ServerTaskApproval.
//
// When returned from a list, ServerTaskApprovalId identifies the server task
// approval the vote was cast against, which is not necessarily the one that was
// queried: votes are tallied across every server task approval sharing a change
// request. The Approver fields are populated by the server and ignored on write.
type Approval struct {
	Id                   string         `json:"Id,omitempty"`
	SpaceId              string         `json:"SpaceId,omitempty"`
	Name                 string         `json:"Name,omitempty"`
	ServerTaskApprovalId string         `json:"ServerTaskApprovalId"`
	UserId               string         `json:"UserId,omitempty"`
	ApproverUsername     string         `json:"ApproverUsername,omitempty"`
	ApproverDisplayName  string         `json:"ApproverDisplayName,omitempty"`
	ApproverEmailAddress string         `json:"ApproverEmailAddress,omitempty"`
	Status               ApprovalStatus `json:"Status"`
	Notes                string         `json:"Notes"`
}
