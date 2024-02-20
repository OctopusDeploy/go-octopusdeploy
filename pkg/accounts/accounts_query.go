package accounts

// AccountsQuery represents parameters to query the Accounts service.
type AccountsQuery struct {
	AccountType AccountType `uri:"accountType,omitempty"`
	IDs         []string    `uri:"ids,omitempty"`
	PartialName string      `uri:"partialName,omitempty"`
	Skip        int         `uri:"skip,omitempty"`
	Take        int         `uri:"take,omitempty"`
}
