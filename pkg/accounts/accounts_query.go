package accounts

// AccountsQuery represents parameters to query the Accounts service.
type AccountsQuery struct {
	AccountType AccountType `uri:"accountType,omitempty" url:"accountType,omitempty"`
	IDs         []string    `uri:"ids,omitempty" url:"ids,omitempty"`
	PartialName string      `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int         `uri:"skip,omitempty" url:"skip,omitempty"`
	Take        int         `uri:"take,omitempty" url:"take,omitempty"`
}
