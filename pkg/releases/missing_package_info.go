package releases

type MissingPackages struct {
	Packages []MissingPackageInfo `json:"Packages"`
}

type MissingPackageInfo struct {
	ID      string `json:"Id,omitempty"`
	Version string `json:"Version,omitempty"`
}
