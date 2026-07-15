package approvalpolicies

import (
	"math"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

const template = "/api/{spaceId}/approvalpolicies{/id}{?skip,take,partialName}"

// GetByID returns the approval policy that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (*ApprovalPolicy, error) {
	return newclient.GetByID[ApprovalPolicy](client, template, spaceID, ID)
}

// Get returns a paginated collection of approval policies matching the query.
func Get(client newclient.Client, spaceID string, query ApprovalPoliciesQuery) (*resources.Resources[*ApprovalPolicy], error) {
	return newclient.GetByQuery[ApprovalPolicy](client, template, spaceID, query)
}

// GetAll returns all approval policies in the space. The list endpoint requires
// the skip and take query parameters, so this issues a single request with the
// maximum take rather than relying on the generic paginated helper (which omits
// them and would be rejected by the server).
func GetAll(client newclient.Client, spaceID string) ([]*ApprovalPolicy, error) {
	res, err := Get(client, spaceID, ApprovalPoliciesQuery{Skip: 0, Take: math.MaxInt32})
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

// Add creates a new approval policy.
func Add(client newclient.Client, spaceID string, policy *ApprovalPolicy) (*ApprovalPolicy, error) {
	if policy == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("policy")
	}
	return newclient.Add[ApprovalPolicy](client, template, spaceID, policy)
}

// Update modifies an existing approval policy.
func Update(client newclient.Client, spaceID string, policy *ApprovalPolicy) (*ApprovalPolicy, error) {
	if policy == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("policy")
	}
	return newclient.Update[ApprovalPolicy](client, template, spaceID, policy.GetID(), policy)
}

// DeleteByID deletes the approval policy that matches the input ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}
