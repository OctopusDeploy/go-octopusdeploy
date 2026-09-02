package approvalrules

import (
	"math"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

const template = "/api/{spaceId}/approvalrules{/id}{?skip,take,partialName}"

// GetByID returns the approval rule that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (*ApprovalRule, error) {
	return newclient.GetByID[ApprovalRule](client, template, spaceID, ID)
}

// Get returns a paginated collection of approval rules matching the query.
func Get(client newclient.Client, spaceID string, query ApprovalRulesQuery) (*resources.Resources[*ApprovalRule], error) {
	return newclient.GetByQuery[ApprovalRule](client, template, spaceID, query)
}

// GetAll returns all approval rules in the space. The list endpoint requires the
// skip and take query parameters, so this issues a single request with the
// maximum take rather than relying on the generic paginated helper (which omits
// them and would be rejected by the server).
func GetAll(client newclient.Client, spaceID string) ([]*ApprovalRule, error) {
	res, err := Get(client, spaceID, ApprovalRulesQuery{Skip: 0, Take: math.MaxInt32})
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

// Add creates a new approval rule.
func Add(client newclient.Client, spaceID string, rule *ApprovalRule) (*ApprovalRule, error) {
	if rule == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("rule")
	}
	return newclient.Add[ApprovalRule](client, template, spaceID, rule)
}

// Update modifies an existing approval rule. The whole resource is replaced, so
// send a fully populated rule — see the ApprovalRule documentation.
func Update(client newclient.Client, spaceID string, rule *ApprovalRule) (*ApprovalRule, error) {
	if rule == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("rule")
	}
	return newclient.Update[ApprovalRule](client, template, spaceID, rule.GetID(), rule)
}

// DeleteByID deletes the approval rule that matches the input ID.
func DeleteByID(client newclient.Client, spaceID string, ID string) error {
	return newclient.DeleteByID(client, template, spaceID, ID)
}
