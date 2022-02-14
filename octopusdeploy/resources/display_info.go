package resources

type DisplayInfo struct {
	ConnectivityCheck     *ConnectivityCheck                   `json:"ConnectivityCheck,omitempty"`
	Description           string                               `json:"Description,omitempty"`
	Label                 string                               `json:"Label,omitempty"`
	ListAPI               *ListAPIMetadata       `json:"ListApi,omitempty"`
	Options               *OptionsMetadata       `json:"Options,omitempty"`
	PropertyApplicability *PropertyApplicability `json:"PropertyApplicability,omitempty"`
	ReadOnly              bool                   `json:"ReadOnly,omitempty"`
	Required              bool                                 `json:"Required,omitempty"`
	ShowCopyToClipboard   bool                                 `json:"ShowCopyToClipboard,omitempty"`
}

func NewDisplayInfo() *DisplayInfo {
	return &DisplayInfo{}
}
