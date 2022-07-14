package resources

type PagedResults struct {
	ItemsPerPage   int    `json:"ItemsPerPage"`
	ItemType       string `json:"ItemType"`
	LastPageNumber int    `json:"LastPageNumber"`
	Links          Links  `json:"Links"`
	NumberOfPages  int    `json:"NumberOfPages"`
	TotalResults   int    `json:"TotalResults"`
}
