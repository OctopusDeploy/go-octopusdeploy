package packages

type PackageNotesResult struct {
	DisplayMessage string `json:"DisplayMessage,omitempty"`
	FailureReason  string `json:"FailureReason,omitempty"`
	Notes          string `json:"Notes,omitempty"`
	Succeeded      bool   `json:"Succeeded"`
}

func NewPackageNotesResult() *PackageNotesResult {
	return &PackageNotesResult{}
}
