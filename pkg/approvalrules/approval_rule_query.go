package approvalrules

// ApprovalRulesQuery is the query for the paginated approval-rules list endpoint.
type ApprovalRulesQuery struct {
	PartialName string `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int    `uri:"skip" url:"skip"`
	Take        int    `uri:"take,omitempty" url:"take,omitempty"`
}
