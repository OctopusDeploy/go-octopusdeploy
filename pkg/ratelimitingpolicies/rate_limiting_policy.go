package ratelimitingpolicies

type RateLimitingPolicy struct {
	ID              string                      `json:"Id"`
	IsBuiltIn       bool                        `json:"IsBuiltIn"`
	Name            string                      `json:"Name"`
	IsEnabled       bool                        `json:"IsEnabled"`
	ScopeType       RateLimitingPolicyScopeType `json:"ScopeType"`
	RequestsPerHour int                         `json:"RequestsPerHour"`
	BurstLimit      int                         `json:"BurstLimit"`
	AuditMode       bool                        `json:"AuditMode"`
}

type GetRateLimitingPolicyByIdRequest struct {
	ID string `uri:"id"`
}

type ListRateLimitingPoliciesRequest struct {
	Skip int `uri:"skip,omitempty"`
	Take int `uri:"take,omitempty"`
}

type ListRateLimitingPoliciesResponse struct {
	ItemType       string               `json:"ItemType"`
	TotalResults   int                  `json:"TotalResults"`
	ItemsPerPage   int                  `json:"ItemsPerPage"`
	Items          []RateLimitingPolicy `json:"Items"`
	NumberOfPages  int                  `json:"NumberOfPages"`
	LastPageNumber int                  `json:"LastPageNumber"`
}

type ModifyRateLimitingPolicyCommand struct {
	ID              string                      `uri:"id" json:"-"`
	Name            string                      `json:"Name"`
	IsEnabled       bool                        `json:"IsEnabled"`
	ScopeType       RateLimitingPolicyScopeType `json:"ScopeType"`
	RequestsPerHour int                         `json:"RequestsPerHour"`
	BurstLimit      int                         `json:"BurstLimit"`
	AuditMode       bool                        `json:"AuditMode"`
}
