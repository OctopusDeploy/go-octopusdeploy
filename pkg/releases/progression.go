package releases

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type Progression struct {
	ChannelEnvironments map[string][]resources.ReferenceDataItem `json:"ChannelEnvironments,omitempty"`
	Environments        []*resources.ReferenceDataItem           `json:"Environments"`
	Releases            []*ReleaseProgression                    `json:"Releases"`

	resources.Resource
}
