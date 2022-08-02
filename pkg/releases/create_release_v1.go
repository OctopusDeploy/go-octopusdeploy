package releases

import "encoding/json"

type ICreateReleaseV1 interface {
}

type CreateReleaseV1 struct {
	// server requires "spaceId" as well, but this redundant data is mapped in MarshalJSON
	SpaceIDOrName         string   `json:"spaceIdOrName"`
	ProjectIDOrName       string   `json:"projectName"`
	PackageVersion        string   `json:"packageVersion,omitempty"`
	GitCommit             string   `json:"gitCommit,omitempty"`
	GitRef                string   `json:"gitRef,omitempty"`
	ReleaseVersion        string   `json:"releaseVersion,omitempty"`
	ChannelIDOrName       string   `json:"channelName,omitempty"`
	Packages              []string `json:"packages,omitempty"`
	ReleaseNotes          string   `json:"releaseNotes,omitempty"`
	IgnoreIfAlreadyExists bool     `json:"ignoreIfAlreadyExists,omitempty"`
	IgnoreChannelRules    bool     `json:"ignoreChannelRules,omitempty"`
	PackagePrerelease     string   `json:"packagePrerelease,omitempty"`
}

type CreateReleaseResponseV1 struct {
	ReleaseID      string `json:"ReleaseId"`
	ReleaseVersion string `json:"ReleaseVersion"`
	// q: the server has this as IDictionary<DeploymentEnvironmentName, IEnumerable<TenantName>> which would
	// translate to map[string][]string in go. Can JSON serialize that?
	AutomaticallyDeployedEnvironments string `json:"AutomaticallyDeployedEnvironments,omitempty"`
}

func NewCreateReleaseV1(spaceIDOrName string, projectIDOrName string) *CreateReleaseV1 {
	return &CreateReleaseV1{
		SpaceIDOrName:   spaceIDOrName,
		ProjectIDOrName: projectIDOrName,
	}
}

// MarshalJSON returns create-release resource (V1) as its JSON encoding.
func (c *CreateReleaseV1) MarshalJSON() ([]byte, error) {
	createReleaseV1 := struct {
		SpaceID string `json:"spaceId"`
		CreateReleaseV1
	}{
		SpaceID: c.SpaceIDOrName,
		CreateReleaseV1: CreateReleaseV1{
			ChannelIDOrName:       c.ChannelIDOrName,
			IgnoreChannelRules:    c.IgnoreChannelRules,
			IgnoreIfAlreadyExists: c.IgnoreIfAlreadyExists,
			GitCommit:             c.GitCommit,
			GitRef:                c.GitRef,
			SpaceIDOrName:         c.SpaceIDOrName,
			PackagePrerelease:     c.PackagePrerelease,
			Packages:              c.Packages,
			PackageVersion:        c.PackageVersion,
			ProjectIDOrName:       c.ProjectIDOrName,
			ReleaseVersion:        c.ReleaseVersion,
			ReleaseNotes:          c.ReleaseNotes,
		},
	}

	return json.Marshal(createReleaseV1)
}
