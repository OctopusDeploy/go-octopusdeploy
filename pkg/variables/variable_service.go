package variables

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
)

type errInvalidVariableServiceParameter struct {
	ParameterName string
}

func (e errInvalidVariableServiceParameter) Error() string {
	return fmt.Sprintf("VariableService: invalid parameter, %s", e.ParameterName)
}

type VariableService struct {
	namesPath   string
	previewPath string

	services.Service
}

func NewVariableService(sling *sling.Sling, uriTemplate string, namesPath string, previewPath string) *VariableService {
	return &VariableService{
		namesPath:   namesPath,
		previewPath: previewPath,
		Service:     services.NewService(constants.ServiceVariableService, sling, uriTemplate),
	}
}

// GetAll fetches a collection of variables for a owner ID.
func (s *VariableService) GetAll(ownerID string) (VariableSet, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return VariableSet{}, err
	}

	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	path := internal.TrimTemplate(s.GetPath())
	path = fmt.Sprintf(path+"/variableset-%s", ownerID)

	resp, err := services.ApiGet(s.GetClient(), new(VariableSet), path)
	if err != nil {
		return VariableSet{}, err
	}

	return *resp.(*VariableSet), nil
}

// GetByID fetches a single variable, located by its ID, from Octopus Deploy for a given owner ID.
func (s *VariableService) GetByID(ownerID string, variableID string) (*Variable, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	if internal.IsEmpty(ownerID) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	if internal.IsEmpty(variableID) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "variableID"}
	}

	variables, err := s.GetAll(ownerID)
	if err != nil {
		return nil, err
	}

	for _, variable := range variables.Variables {
		if variable.GetID() == variableID {
			return variable, nil
		}
	}

	return nil, &core.APIError{
		StatusCode:   404,
		ErrorMessage: fmt.Sprintf("Variable ID, %s could not be found with owner ID, %s.", variableID, ownerID),
	}
}

// GetByName fetches variables, located by their name, from Octopus Deploy for a given owner ID. As variable
// names can appear more than once under different scopes, a VariableScope must also be provided, which will
// be used to locate the appropriate variables.
func (s *VariableService) GetByName(ownerID string, name string, scope *VariableScope) ([]*Variable, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	if internal.IsEmpty(ownerID) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	if internal.IsEmpty(name) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "name"}
	}

	variables, err := s.GetAll(ownerID)
	if err != nil {
		return nil, err
	}

	var matchedVariables []*Variable

	for _, variable := range variables.Variables {
		if variable.Name == name {
			matchScope, _, err := s.MatchesScope(variable.Scope, scope)
			if err != nil {
				return nil, err
			}
			if matchScope {
				matchedVariables = append(matchedVariables, variable)
			}
		}
	}

	return matchedVariables, nil
}

// AddSingle adds a single variable to a owner ID. This automates the act of fetching
// the variable set, adding a new item to it, and posting back to Octopus
func (s *VariableService) AddSingle(ownerID string, variable *Variable) (VariableSet, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return VariableSet{}, err
	}

	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	variables, err := s.GetAll(ownerID)
	if err != nil {
		return VariableSet{}, err
	}

	variables.Variables = append(variables.Variables, variable)
	return s.Update(ownerID, variables)
}

// UpdateSingle adds a single variable to a owner ID. This automates the act of fetching
// the variable set, updating the existing item, and posting back to Octopus
func (s *VariableService) UpdateSingle(ownerID string, variable *Variable) (VariableSet, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return VariableSet{}, err
	}

	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	variables, err := s.GetAll(ownerID)
	if err != nil {
		return VariableSet{}, err
	}

	for k, v := range variables.Variables {
		if v.GetID() == variable.ID {
			variables.Variables[k] = variable
			return s.Update(ownerID, variables)
		}
	}

	return VariableSet{}, services.ErrItemNotFound
}

// DeleteSingle removes a single variable from a owner ID. This automates the act of fetching
// the variable set, removing the existing item, and posting back to Octopus
func (s *VariableService) DeleteSingle(ownerID string, variableID string) (VariableSet, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return VariableSet{}, err
	}

	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	if internal.IsEmpty(variableID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "variableID"}
	}

	variableSet, err := s.GetAll(ownerID)
	if err != nil {
		return VariableSet{}, err
	}

	var found bool
	for k, v := range variableSet.Variables {
		if v.GetID() == variableID {
			variableSet.Variables = append(variableSet.Variables[:k], variableSet.Variables[k+1:]...)
			found = true
		}
	}

	if !found {
		return VariableSet{}, &core.APIError{
			StatusCode:   404,
			ErrorMessage: fmt.Sprintf("Variable ID, %s could not be found with owner ID, %s.", variableID, ownerID),
		}
	}

	return s.Update(ownerID, variableSet)
}

// Update takes an entire variable set and posts the entire set back to Octopus Deploy. There are individual
// functions like AddSingle and UpdateSingle that can make this process more of a "typical" CRUD Octopus command.
func (s *VariableService) Update(ownerID string, variableSet VariableSet) (VariableSet, error) {
	err := services.ValidateInternalState(s)
	if err != nil {
		return VariableSet{}, err
	}

	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	path := internal.TrimTemplate(s.GetPath())
	path = fmt.Sprintf(path+"/variableset-%s", ownerID)

	if _, err := services.ApiUpdate(s.GetClient(), variableSet, new(VariableSet), path); err != nil {
		return VariableSet{}, err
	}

	// 2021-04-22 (John Bristowe): we need to retrieve the variable set (again)
	// via HTTP GET (below) due to a bug for HTTP POST and HTTP PUT which will
	// provide a null scope value set in their responses

	return s.GetAll(ownerID)
}

// MatchesScope compares two different scopes to see if they match. Generally used for comparing the scope of
// an existing variable against a desired state. Only supports Environment, Role, Machine, Action and Channel
// for scope options. Returns true if definedScope is nil or all elements are empty. Also returns a VariableScope
// of all the scopes that were matched
func (s *VariableService) MatchesScope(variableScope VariableScope, definedScope *VariableScope) (bool, *VariableScope, error) {
	err := services.ValidateInternalState(s)
	if err != nil {
		return false, nil, err
	}

	if definedScope == nil {
		return true, &VariableScope{}, nil
	}

	if definedScope.IsEmpty() {
		return true, &VariableScope{}, nil
	}

	var matchedScopes VariableScope
	var matched bool

	for _, e1 := range definedScope.Environments {
		for _, e2 := range variableScope.Environments {
			if e1 == e2 {
				matched = true
				matchedScopes.Environments = append(matchedScopes.Environments, e1)
			}
		}
	}

	for _, r1 := range definedScope.Roles {
		for _, r2 := range variableScope.Roles {
			if r1 == r2 {
				matched = true
				matchedScopes.Roles = append(matchedScopes.Roles, r1)
			}
		}
	}

	for _, m1 := range definedScope.Machines {
		for _, m2 := range variableScope.Machines {
			if m1 == m2 {
				matched = true
				matchedScopes.Machines = append(matchedScopes.Machines, m1)
			}
		}
	}

	for _, a1 := range definedScope.Actions {
		for _, a2 := range variableScope.Actions {
			if a1 == a2 {
				matched = true
				matchedScopes.Actions = append(matchedScopes.Actions, a1)
			}
		}
	}

	for _, c1 := range definedScope.Channels {
		for _, c2 := range variableScope.Channels {
			if c1 == c2 {
				matched = true
				matchedScopes.Channels = append(matchedScopes.Channels, c1)
			}
		}
	}

	for _, c1 := range definedScope.TenantTags {
		for _, c2 := range variableScope.TenantTags {
			if c1 == c2 {
				matched = true
				matchedScopes.TenantTags = append(matchedScopes.TenantTags, c1)
			}
		}
	}

	return matched, &matchedScopes, nil
}

var _ services.IService = &VariableService{}
