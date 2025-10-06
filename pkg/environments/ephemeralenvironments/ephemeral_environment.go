package ephemeralenvironments

type EphemeralEnvironment struct {
	ID                  string `json:"Id"`
	Name                string `json:"Name"`
	SpaceID             string `json:"SpaceId"`
	Slug                string `json:"Slug"`
	Description         string `json:"Description"`
	Type                string `json:"Type"`
	SortOrder           int    `json:"SortOrder"`
	UseGuidedFailure    bool   `json:"UseGuidedFailure"`
	ParentEnvironmentId string `json:"ParentEnvironmentId"`
}

type EphemeralEnvironmentV2Response struct {
	Environments *EphemeralEnvironmentResponse `json:Environments`
}

type EphemeralEnvironmentResponse struct {
	ItemType       string                  `json:"ItemType"`
	TotalResults   int                     `json:"TotalResults"`
	ItemsPerPage   int                     `json:"ItemsPerPage"`
	NumberOfPages  int                     `json:"NumberOfPages"`
	LastPageNumber int                     `json:"LastPageNumber"`
	Items          []*EphemeralEnvironment `json:"Items"`
}
