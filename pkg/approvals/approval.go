package approvals

// ApprovalStatus is an individual approver's decision.
type ApprovalStatus string

const (
	ApprovalStatusApproved ApprovalStatus = "Approved"
	ApprovalStatusRejected ApprovalStatus = "Rejected"
)

// Approval is a single person's decision recorded against a ServerTaskApproval.
type Approval struct {
	Id                   string         `json:"Id,omitempty"`
	SpaceId              string         `json:"SpaceId,omitempty"`
	Name                 string         `json:"Name,omitempty"`
	ServerTaskApprovalId string         `json:"ServerTaskApprovalId"`
	UserId               string         `json:"UserId,omitempty"`
	Status               ApprovalStatus `json:"Status"`
	Notes                string         `json:"Notes"`
}
