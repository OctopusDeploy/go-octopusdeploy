package octopusdeploy

type Progression struct {
	ChannelEnvironments map[string][]ReferenceDataItem `json:"ChannelEnvironments,omitempty"`
	Environments        []*ReferenceDataItem           `json:"Environments"`
	Releases            []*ReleaseProgression          `json:"Releases"`

	resource
}
