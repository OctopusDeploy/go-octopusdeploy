package model

type PagedResults struct {
	ItemType       string `json:"ItemType"`
	TotalResults   int    `json:"TotalResults"`
	NumberOfPages  int    `json:"NumberOfPages"`
	LastPageNumber int    `json:"LastPageNumber"`
	ItemsPerPage   int    `json:"ItemsPerPage"`
	IsStale        bool   `json:"IsStale"`
	Links          Links  `json:"Links"`
}
