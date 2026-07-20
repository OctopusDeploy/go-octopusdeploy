package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/ratelimitingpolicies"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListRateLimitingPolicies(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	response, err := ratelimitingpolicies.List(
		client,
		ratelimitingpolicies.ListRateLimitingPoliciesRequest{})

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.Items, 3)
	assert.Equal(t, 30, response.ItemsPerPage)
	assert.Equal(t, 3, response.TotalResults)
	assert.Equal(t, 1, response.NumberOfPages)
	assert.Equal(t, 0, response.LastPageNumber)
	assert.Equal(t, "RateLimitingPolicy", response.ItemType)
}

func TestListRateLimitingPoliciesSkipTake(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	response1, err1 := ratelimitingpolicies.List(
		client,
		ratelimitingpolicies.ListRateLimitingPoliciesRequest{
			Skip: 0,
			Take: 1,
		})

	assert.NoError(t, err1)
	assert.NotNil(t, response1)
	assert.Len(t, response1.Items, 1)
	assert.Equal(t, 1, response1.ItemsPerPage)
	assert.Equal(t, 3, response1.TotalResults)
	assert.Equal(t, 3, response1.NumberOfPages)
	assert.Equal(t, 2, response1.LastPageNumber)
	assert.Equal(t, "RateLimitingPolicy-1", response1.Items[0].Id)

	response2, err2 := ratelimitingpolicies.List(
		client,
		ratelimitingpolicies.ListRateLimitingPoliciesRequest{
			Skip: 1,
			Take: 1,
		})

	assert.NoError(t, err2)
	assert.NotNil(t, response2)
	assert.Len(t, response2.Items, 1)
	assert.Equal(t, 1, response2.ItemsPerPage)
	assert.Equal(t, 3, response2.TotalResults)
	assert.Equal(t, 3, response2.NumberOfPages)
	assert.Equal(t, 2, response2.LastPageNumber)
	assert.Equal(t, "RateLimitingPolicy-2", response2.Items[0].Id)
}

func TestGetRateLimitingPolicyByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	listResponse, listError := ratelimitingpolicies.List(
		client,
		ratelimitingpolicies.ListRateLimitingPoliciesRequest{})

	assert.NoError(t, listError)
	assert.NotNil(t, listResponse)
	assert.Len(t, listResponse.Items, 3)

	for _, policy := range listResponse.Items {
		getResponse, getError := ratelimitingpolicies.GetByID(
			client,
			ratelimitingpolicies.GetRateLimitingPolicyByIdRequest{
				Id: policy.Id,
			})

		assert.NoError(t, getError)
		assert.NotNil(t, getResponse)
		assert.Equal(t, policy.Id, getResponse.Id)
	}
}

func TestModifyRateLimitingPolicy(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	listResponse, listError := ratelimitingpolicies.List(
		client,
		ratelimitingpolicies.ListRateLimitingPoliciesRequest{
			Take: 1,
		})

	assert.NoError(t, listError)
	assert.NotNil(t, listResponse)
	assert.Len(t, listResponse.Items, 1)
	policy := listResponse.Items[0]

	testModify := func(isEnabled bool, requestsPerHour int, burstLimit int, auditMode bool) {
		modifyResponse, modifyError := ratelimitingpolicies.Modify(
			client,
			ratelimitingpolicies.ModifyRateLimitingPolicyCommand{
				Id:              policy.Id,
				Name:            policy.Name,
				ScopeType:       policy.ScopeType,

				IsEnabled:       isEnabled,
				RequestsPerHour: requestsPerHour,
				BurstLimit:      burstLimit,
				AuditMode:       auditMode,
			})
		assert.NoError(t, modifyError)
		assert.NotNil(t, modifyResponse)
		assert.Equal(t, policy.Id, modifyResponse.Id)
		assert.Equal(t, policy.Name, modifyResponse.Name)
		assert.Equal(t, policy.ScopeType, modifyResponse.ScopeType)
		assert.Equal(t, isEnabled, modifyResponse.IsEnabled)
		assert.Equal(t, requestsPerHour, modifyResponse.RequestsPerHour)
		assert.Equal(t, burstLimit, modifyResponse.BurstLimit)
		assert.Equal(t, auditMode, modifyResponse.AuditMode)

		// Ensure the modify actually set those properties with a follow-up GET
		getResponse, getError := ratelimitingpolicies.GetByID(
			client,
			ratelimitingpolicies.GetRateLimitingPolicyByIdRequest{
				Id: policy.Id,
			})

		assert.NoError(t, getError)
		assert.NotNil(t, getResponse)
		assert.Equal(t, policy.Id, getResponse.Id)
		assert.Equal(t, policy.Name, getResponse.Name)
		assert.Equal(t, policy.ScopeType, getResponse.ScopeType)
		assert.Equal(t, isEnabled, getResponse.IsEnabled)
		assert.Equal(t, requestsPerHour, getResponse.RequestsPerHour)
		assert.Equal(t, burstLimit, getResponse.BurstLimit)
		assert.Equal(t, auditMode, getResponse.AuditMode)
	}

	// Modify twice in case the properties were already set to what we modified them to
	testModify(false, 12, 34, true)
	testModify(true, 56, 78, false)
}

func TestModifyRateLimitingPolicyError(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	listResponse, listError := ratelimitingpolicies.List(
		client,
		ratelimitingpolicies.ListRateLimitingPoliciesRequest{
			Take: 1,
		})

	assert.NoError(t, listError)
	assert.NotNil(t, listResponse)
	assert.Len(t, listResponse.Items, 1)

	policy := listResponse.Items[0]
	assert.Equal(t, true, policy.IsBuiltIn)

	modifyResponse, modifyError := ratelimitingpolicies.Modify(
		client,
		ratelimitingpolicies.ModifyRateLimitingPolicyCommand{
			Id:              policy.Id,
			Name:            "New name", // Not allowed to change built-in policy names
			ScopeType:       policy.ScopeType,

			IsEnabled:       policy.IsEnabled,
			RequestsPerHour: policy.RequestsPerHour,
			BurstLimit:      policy.BurstLimit,
			AuditMode:       policy.AuditMode,
		})
	assert.Nil(t, modifyResponse)
	assert.Error(t, modifyError)
	assert.Equal(
		t,
		"The name of a built-in rate limiting policy cannot be changed.",
		modifyError.(*core.APIError).Errors[0])
}
