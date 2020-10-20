package octopusdeploy

type Tag struct {
	ID               string `json:"Id"`
	Name             string `json:"Name"`
	Color            string `json:"Color"`
	CanonicalTagName string `json:"CanonicalTagName"`
	Description      string `json:"Description"`
	SortOrder        int    `json:"SortOrder"`
}
