package octopusdeploy

// Releases defines a collection of Release instance with built-in support for paged results from the API.
type Releases struct {
	Items []*Release `json:"Items"`
	PagedResults
}
