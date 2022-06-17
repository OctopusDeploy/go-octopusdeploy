package certificates

type CertificatesQuery struct {
	Archived    string   `uri:"archived,omitempty" url:"archived,omitempty"`
	FirstResult string   `uri:"firstResult,omitempty" url:"firstResult,omitempty"`
	IDs         []string `uri:"ids,omitempty" url:"ids,omitempty"`
	OrderBy     string   `uri:"orderBy,omitempty" url:"orderBy,omitempty"`
	PartialName string   `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Search      string   `uri:"search,omitempty" url:"search,omitempty"`
	Skip        int      `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int      `uri:"take,omitempty" url:"take,omitempty"`
	Tenant      string   `uri:"tenant,omitempty" url:"tenant,omitempty"`
}
