package resources

type Links struct {
	Self        string `json:"Self"`
	Template    string `json:"Template"`
	PageAll     string `json:"Page.All"`
	PageCurrent string `json:"Page.Current"`
	PageLast    string `json:"Page.Last"`
	PageNext    string `json:"Page.Next"`
}
