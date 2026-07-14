package approvalpolicies

// ApprovalPoliciesQuery is the query for the paginated approval-policies list endpoint.
type ApprovalPoliciesQuery struct {
	PartialName string `uri:"partialName,omitempty" url:"partialName,omitempty"`
	Skip        int    `uri:"skip" url:"skip"`
	Take        int    `uri:"take,omitempty" url:"take,omitempty"`
}
