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
