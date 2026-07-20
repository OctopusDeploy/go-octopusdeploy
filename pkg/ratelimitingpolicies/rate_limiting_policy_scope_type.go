package ratelimitingpolicies

type RateLimitingPolicyScopeType int

const (
	Unauthenticated RateLimitingPolicyScopeType = iota
	AuthenticatedHuman
	AuthenticatedAgent
)
