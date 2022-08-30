package releases

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type CreateReleaseCommandV1 struct {
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

func NewCreateReleaseCommandV1(spaceIDOrName string, projectIDOrName string) *CreateReleaseCommandV1 {
	return &CreateReleaseCommandV1{
		SpaceIDOrName:   spaceIDOrName,
		ProjectIDOrName: projectIDOrName,
	}
}

// MarshalJSON adds the redundant 'spaceId' parameter which is required by the server
func (c *CreateReleaseCommandV1) MarshalJSON() ([]byte, error) {
	createReleaseV1 := struct {
		SpaceID string `json:"spaceId"`
		CreateReleaseCommandV1
	}{
		SpaceID:                c.SpaceIDOrName,
		CreateReleaseCommandV1: *c,
	}
	return json.Marshal(createReleaseV1)
}

func CreateReleaseV1(client newclient.Client, command *CreateReleaseCommandV1) (*CreateReleaseResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("CreateV1", "command")
	}
	if client.SpaceID() == "" {
		return nil, internal.CreateInvalidClientStateError("CreateV1")
	}

	// Note: command has a SpaceIDOrName field in it, which carries the space, however, we can't use it
	// as the server's route URL *requires* a space **ID**, not a name. In fact, the client's spaceID should always win.
	command.SpaceIDOrName = client.SpaceID()
	url, err := client.URITemplateCache().Expand(uritemplates.CreateReleaseCommandV1, map[string]any{"spaceId": client.SpaceID()})
	if err != nil {
		return nil, err
	}
	resp, err := services.ApiPost(client.Sling(), command, new(CreateReleaseResponseV1), url)
	if err != nil {
		return nil, err
	}
	return resp.(*CreateReleaseResponseV1), nil
}
