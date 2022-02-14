package resources

type PackageDescription struct {
	Description   string            `json:"Description,omitempty"`
	ID            string            `json:"Id,omitempty"`
	LatestVersion string            `json:"LatestVersion,omitempty"`
	Links         map[string]string `json:"Links,omitempty"`
	Name          string            `json:"Name,omitempty"`
}

type PackageDescriptions struct {
	Items []*PackageDescription `json:"Items"`
	PagedResults
}
