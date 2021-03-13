package octopusdeploy

type PackageReference struct {
	AcquisitionLocation string            `json:"AcquisitionLocation"` // This can be an expression
	FeedID              string            `json:"FeedId"`
	ID                  string            `json:"Id,omitempty"`
	Name                string            `json:"Name,omitempty"`
	PackageID           string            `json:"PackageId,omitempty"`
	Properties          map[string]string `json:"Properties"`
}
