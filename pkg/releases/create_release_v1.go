package releases

import "encoding/json"

type ICreateReleaseV1 interface {
}

type CreateReleaseV1 struct {
	ChannelNameOrID       string   `json:"channelIdOrName,omitempty"`
	IgnoreChannelRules    bool     `json:"ignoreChannelRules,omitempty"`
	IgnoreIfAlreadyExists bool     `json:"ignoreIfAlreadyExists,omitempty"`
	GitCommit             string   `json:"gitCommit,omitempty"`
	GitRef                string   `json:"gitRef,omitempty"`
	SpaceNameOrID         string   `json:"spaceIdOrName"`
	PackagePrerelease     string   `json:"packagePrerelease,omitempty"`
	Packages              []string `json:"packages,omitempty"`
	PackageVersion        string   `json:"packageVersion,omitempty"`
	ProjectNameOrID       string   `json:"projectName"`
	ReleaseVersion        string   `json:"releaseVersion,omitempty"`
	ReleaseNotes          string   `json:"releaseNotes,omitempty"`
}

type CreateReleaseResponseV1 struct {
	ReleaseID                         string `json:"ReleaseId"`
	ReleaseVersion                    string `json:"ReleaseVersion"`
	AutomaticallyDeployedEnvironments string `json:"AutomaticallyDeployedEnvironments,omitempty"`
}

func NewCreateReleaseV1(spaceNameOrID string, projectNameOrID string) *CreateReleaseV1 {
	return &CreateReleaseV1{
		SpaceNameOrID:   spaceNameOrID,
		ProjectNameOrID: projectNameOrID,
	}
}

// MarshalJSON returns create-release resource (V1) as its JSON encoding.
func (c *CreateReleaseV1) MarshalJSON() ([]byte, error) {
	createReleaseV1 := struct {
		SpaceID string `json:"spaceId"`
		CreateReleaseV1
	}{
		SpaceID: c.SpaceNameOrID,
		CreateReleaseV1: CreateReleaseV1{
			ChannelNameOrID:       c.ChannelNameOrID,
			IgnoreChannelRules:    c.IgnoreChannelRules,
			IgnoreIfAlreadyExists: c.IgnoreIfAlreadyExists,
			GitCommit:             c.GitCommit,
			GitRef:                c.GitRef,
			SpaceNameOrID:         c.SpaceNameOrID,
			PackagePrerelease:     c.PackagePrerelease,
			Packages:              c.Packages,
			PackageVersion:        c.PackageVersion,
			ProjectNameOrID:       c.ProjectNameOrID,
			ReleaseVersion:        c.ReleaseVersion,
			ReleaseNotes:          c.ReleaseNotes,
		},
	}

	return json.Marshal(createReleaseV1)
}
