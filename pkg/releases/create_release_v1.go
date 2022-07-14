package releases

type ICreateReleaseV1 interface {
}

type CreateReleaseV1 struct {
	SpaceID               string   `json:"spaceId"`
	SpaceIdOrName         string   `json:"spaceIdOrName"`
	ProjectIdOrName       string   `json:"projectIdOrName"`
	PackageVersion        string   `json:"packageVersion,omitempty"`
	GitCommit             string   `json:"gitCommit,omitempty"`
	GitRef                string   `json:"gitRef,omitempty"`
	ReleaseVersion        string   `json:"releaseVersion,omitempty"`
	ChannelName           string   `json:"channelIdOrName,omitempty"`
	Packages              []string `json:"packages,omitempty"`
	ReleaseNotes          string   `json:"releaseNotes,omitempty"`
	IgnoreIfAlreadyExists bool     `json:"ignoreIfAlreadyExists,omitempty"`
	IgnoreChannelRules    bool     `json:"ignoreChannelRules,omitempty"`
	PackagePrerelease     string   `json:"packagePrerelease,omitempty"`
}

type CreateReleaseResponseV1 struct {
	ReleaseID                         string `json:"ReleaseId"`
	ReleaseVersion                    string `json:"ReleaseVersion"`
	AutomaticallyDeployedEnvironments string `json:"AutomaticallyDeployedEnvironments,omitempty"`
}

func NewCreateReleaseV1(spaceIdOrName string, projectIdOrName string) *CreateReleaseV1 {
	return &CreateReleaseV1{
		SpaceID:         spaceIdOrName,
		SpaceIdOrName:   spaceIdOrName,
		ProjectIdOrName: projectIdOrName,
	}
}
