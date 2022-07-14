package packages

type SelectedPackage struct {
	ActionName           string `json:"ActionName,omitempty"`
	PackageReferenceName string `json:"PackageReferenceName,omitempty"`
	StepName             string `json:"StepName,omitempty"`
	Version              string `json:"Version,omitempty"`
}
