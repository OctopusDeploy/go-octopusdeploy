package releases

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type CreateReleaseCommandV1 struct {
	// Note: the server requires both SpaceID and SpaceIDOrName, and is capable of looking up names from the JSON
	// payload.
	// It'd be nice to allow SpaceIDOrName, but the current server implementation requires a SpaceID
	// (not a name) in the URL route, so we must force the caller to specify an ID only.
	SpaceID               string   `json:"spaceId"`
	ProjectIDOrName       string   `json:"projectName"`
	PackageVersion        string   `json:"packageVersion,omitempty"`
	GitCommit             string   `json:"gitCommit,omitempty"`
	GitRef                string   `json:"gitRef,omitempty"`
	ReleaseVersion        string   `json:"releaseVersion,omitempty"`
	ChannelIDOrName       string   `json:"channelName,omitempty"`
	Packages              []string `json:"packages,omitempty"`
	ReleaseNotes          string   `json:"releaseNotes,omitempty"`
	IgnoreIfAlreadyExists bool     `json:"ignoreIfAlreadyExists"`
	IgnoreChannelRules    bool     `json:"ignoreChannelRules"`
	PackagePrerelease     string   `json:"packagePrerelease,omitempty"`
}

type CreateReleaseResponseV1 struct {
	ReleaseID      string `json:"ReleaseId"`
	ReleaseVersion string `json:"ReleaseVersion"`
}

func NewCreateReleaseCommandV1(spaceID string, projectIDOrName string) *CreateReleaseCommandV1 {
	return &CreateReleaseCommandV1{
		SpaceID:         spaceID,
		ProjectIDOrName: projectIDOrName,
	}
}

// MarshalJSON adds the redundant 'SpaceIDOrName' parameter which is required by the server
func (c *CreateReleaseCommandV1) MarshalJSON() ([]byte, error) {
	createReleaseV1 := struct {
		SpaceIDOrName string `json:"spaceIdOrName"`
		CreateReleaseCommandV1
	}{
		SpaceIDOrName:          c.SpaceID,
		CreateReleaseCommandV1: *c,
	}
	return json.Marshal(createReleaseV1)
}

func CreateReleaseV1(client newclient.Client, command *CreateReleaseCommandV1) (*CreateReleaseResponseV1, error) {
	if command == nil {
		return nil, internal.CreateInvalidParameterError("CreateReleaseV1", "command")
	}
	if command.SpaceID == "" {
		return nil, internal.CreateInvalidParameterError("CreateReleaseV1", "command.SpaceID")
	}

	// Note: command has a SpaceIDOrName field in it, which carries the space, however, we can't use it
	// as the server's route URL *requires* a space **ID**, not a name. In fact, the client's spaceID should always win.
	path, err := client.URITemplateCache().Expand(uritemplates.CreateReleaseCommandV1, map[string]any{"spaceId": command.SpaceID})
	if err != nil {
		return nil, err
	}
	return newclient.Post[CreateReleaseResponseV1](client.HttpSession(), path, command)
}
