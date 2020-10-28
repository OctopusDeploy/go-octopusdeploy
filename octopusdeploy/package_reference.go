package octopusdeploy

type PackageReference struct {
	ID                  string            `json:"Id,omitempty"`
	Name                string            `json:"Name,omitempty"`
	PackageID           string            `json:"PackageId,omitempty"`
	FeedID              string            `json:"FeedId"`
	AcquisitionLocation string            `json:"AcquisitionLocation"` // This can be an expression
	Properties          map[string]string `json:"Properties"`
}
