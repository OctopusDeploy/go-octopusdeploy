package metadata

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"

type DisplayInfo struct {
	ConnectivityCheck     *core.ConnectivityCheck `json:"ConnectivityCheck,omitempty"`
	Description           string                  `json:"Description,omitempty"`
	Label                 string                  `json:"Label,omitempty"`
	ListAPI               *ListAPIMetadata        `json:"ListApi,omitempty"`
	Options               *OptionsMetadata        `json:"Options,omitempty"`
	PropertyApplicability *PropertyApplicability  `json:"PropertyApplicability,omitempty"`
	ReadOnly              bool                    `json:"ReadOnly"`
	Required              bool                    `json:"Required"`
	ShowCopyToClipboard   bool                    `json:"ShowCopyToClipboard"`
}

func NewDisplayInfo() *DisplayInfo {
	return &DisplayInfo{}
}
