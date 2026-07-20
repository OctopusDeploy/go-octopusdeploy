package ratelimitingpolicies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
)

const rateLimitingPoliciesTemplate = "/api/ratelimitingpolicies{/id}{?skip,take}"

// GetByID returns the rate limiting policy that matches the given ID.
func GetByID(client newclient.Client, request GetRateLimitingPolicyByIdRequest) (*RateLimitingPolicy, error) {
	if request.Id == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("Id")
	}

	path, pathError := client.URITemplateCache().Expand(rateLimitingPoliciesTemplate, request)
	if pathError != nil { return nil, pathError }

	result, resultError := newclient.Get[RateLimitingPolicy](client.HttpSession(), path)
	if resultError != nil { return nil, resultError }

	return result, nil
}

// List returns a paginated collection of rate limiting policies.
func List(client newclient.Client, request ListRateLimitingPoliciesRequest) (*ListRateLimitingPoliciesResponse, error) {
	path, pathError := client.URITemplateCache().Expand(rateLimitingPoliciesTemplate, request)
	if pathError != nil { return nil, pathError }

	result, resultError := newclient.Get[ListRateLimitingPoliciesResponse](client.HttpSession(), path)
	if resultError != nil { return nil, resultError }

	return result, nil
}

// Modify changes the rate limiting policy that matches the given ID.
func Modify(client newclient.Client, command ModifyRateLimitingPolicyCommand) (*RateLimitingPolicy, error) {
	if command.Id == "" {
		return nil, internal.CreateRequiredParameterIsEmptyError("Id")
	}

	path, pathError := client.URITemplateCache().Expand(rateLimitingPoliciesTemplate, command)
	if pathError != nil { return nil, pathError }

	result, resultError := newclient.Put[RateLimitingPolicy](client.HttpSession(), path, command)
	if resultError != nil { return nil, resultError }

	return result, nil
}
