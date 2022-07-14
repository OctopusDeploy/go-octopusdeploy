package packages

type PackageNote struct {
	FeedID    string              `json:"FeedId,omitempty"`
	Notes     *PackageNotesResult `json:"Notes,omitempty"`
	PackageID string              `json:"PackageId,omitempty"`
	Version   string              `json:"Version,omitempty"`
}

func NewPackageNote() *PackageNote {
	return &PackageNote{}
}
