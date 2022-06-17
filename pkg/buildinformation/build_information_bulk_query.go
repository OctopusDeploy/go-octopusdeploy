package buildinformation

type BuildInformationBulkQuery struct {
	IDs []string `uri:"ids,omitempty" url:"ids,omitempty"`
}
