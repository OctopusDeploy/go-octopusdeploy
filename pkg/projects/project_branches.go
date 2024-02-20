package projects

type CreateBranchRequest struct {
	BaseGitRef    string `json:"BaseGitRef"`
	NewBranchName string `json:"NewBranchName"`
}
