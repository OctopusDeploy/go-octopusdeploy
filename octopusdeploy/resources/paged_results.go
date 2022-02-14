package resources

type PagedResults[T Resource] struct {
	ItemType       string `json:"ItemType"`
	TotalResults   int    `json:"TotalResults"`
	ItemsPerPage   int    `json:"ItemsPerPage"`
	NumberOfPages  int    `json:"NumberOfPages"`
	LastPageNumber int    `json:"LastPageNumber"`
	Items          []T    `json:"Items"`
}
