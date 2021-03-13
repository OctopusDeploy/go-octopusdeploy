package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

type errInvalidvariableServiceParameter struct {
	ParameterName string
}

func (e errInvalidvariableServiceParameter) Error() string {
	return fmt.Sprintf("variableService: invalid parameter, %s", e.ParameterName)
}

type variableService struct {
	namesPath   string
	previewPath string

	service
}

func newVariableService(sling *sling.Sling, uriTemplate string, namesPath string, previewPath string) *variableService {
	return &variableService{
		namesPath:   namesPath,
		previewPath: previewPath,
		service:     newService(ServiceVariableService, sling, uriTemplate),
	}
}

// GetAll fetches an entire VariableSet from Octopus Deploy for a given Project ID.
func (s variableService) GetAll(projectID string) (*Variables, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "projectID"}
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/variableset-%s", projectID)

	resp, err := apiGet(s.getClient(), new(Variables), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Variables), nil
}

// GetByID fetches a single variable, located by its ID, from Octopus Deploy for a given Project ID.
func (s variableService) GetByID(projectID string, variableID string) (*Variable, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "projectID"}
	}

	if isEmpty(variableID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "variableID"}
	}

	variables, err := s.GetAll(projectID)
	if err != nil {
		return nil, err
	}

	for _, variable := range variables.Variables {
		if variable.GetID() == variableID {
			return &variable, nil
		}
	}

	return nil, nil
}

// GetByName fetches variables, located by their name, from Octopus Deploy for a given Project ID. As variable
// names can appear more than once under different scopes, a VariableScope must also be provided, which will
// be used to locate the appropriate variables.
func (s variableService) GetByName(projectID string, name string, scope *VariableScope) ([]Variable, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "projectID"}
	}

	if isEmpty(name) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "name"}
	}

	if scope == nil {
		return nil, errInvalidvariableServiceParameter{ParameterName: "scope"}
	}

	variables, err := s.GetAll(projectID)
	if err != nil {
		return nil, err
	}

	var matchedVariables []Variable

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

// AddSingle adds a single variable to a project ID. This automates the act of fetching
// the variable set, adding a new item to it, and posting back to Octopus
func (s variableService) AddSingle(projectID string, variable *Variable) (*Variables, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "projectID"}
	}

	if variable == nil {
		return nil, errInvalidvariableServiceParameter{ParameterName: "variable"}
	}

	variables, err := s.GetAll(projectID)

	if err != nil {
		return nil, err
	}

	variables.Variables = append(variables.Variables, *variable)
	return s.Update(projectID, variables)
}

// UpdateSingle adds a single variable to a project ID. This automates the act of fetching
// the variable set, updating the existing item, and posting back to Octopus
func (s variableService) UpdateSingle(projectID string, variable *Variable) (*Variables, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "projectID"}
	}

	if variable == nil {
		return nil, errInvalidvariableServiceParameter{ParameterName: "variable"}
	}

	variables, err := s.GetAll(projectID)

	if err != nil {
		return nil, err
	}

	var found bool
	for i, existingVar := range variables.Variables {
		if existingVar.GetID() == variable.ID {
			variables.Variables[i] = *variable
			found = true
		}
	}

	if !found {
		return nil, ErrItemNotFound
	}

	return s.Update(projectID, variables)
}

// DeleteSingle removes a single variable from a project ID. This automates the act of fetching
// the variable set, removing the existing item, and posting back to Octopus
func (s variableService) DeleteSingle(projectID string, variableID string) (*Variables, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "projectID"}
	}

	if isEmpty(variableID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "variableID"}
	}

	variables, err := s.GetAll(projectID)
	if err != nil {
		return nil, err
	}

	var found bool
	for i, existingVar := range variables.Variables {
		if existingVar.GetID() == variableID {
			variables.Variables = append(variables.Variables[:i], variables.Variables[i+1:]...)
			found = true
		}
	}

	if !found {
		return nil, ErrItemNotFound
	}

	return s.Update(projectID, variables)
}

// Update takes an entire variable set and posts the entire set back to Octopus Deploy. There are individual
// functions like AddSingle and UpdateSingle that can make this process more of a "typical" CRUD Octopus command.
func (s variableService) Update(projectID string, variableSet *Variables) (*Variables, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{ParameterName: "projectID"}
	}

	if variableSet == nil {
		return nil, errInvalidvariableServiceParameter{ParameterName: "variableSet"}
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/variableset-%s", projectID)

	resp, err := apiUpdate(s.getClient(), variableSet, new(Variables), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Variables), nil
}

// MatchesScope compares two different scopes to see if they match. Generally used for comparing the scope of
// an existing variable against a desired state. Only supports Environment, Role, Machine, Action and Channel
// for scope options. Returns true if definedScope is nil or all elements are empty. Also returns a VariableScope
// of all the scopes that were matched
func (s variableService) MatchesScope(variableScope *VariableScope, definedScope *VariableScope) (bool, *VariableScope, error) {
	err := validateInternalState(s)
	if err != nil {
		return false, nil, err
	}

	if variableScope == nil {
		return false, nil, errInvalidvariableServiceParameter{ParameterName: "variableScope"}
	}

	//If the scope supplied is nil then match everything
	if definedScope == nil {
		return true, NewVariableScope(), nil
	}

	//Unsupported scopes
	if len(definedScope.Private) > 0 {
		return false, nil, fmt.Errorf("Private is not a supported scope for variable matching")
	}
	if len(definedScope.Projects) > 0 {
		return false, nil, fmt.Errorf("Project is not a supported scope for variable matching")
	}
	if len(definedScope.TargetRoles) > 0 {
		return false, nil, fmt.Errorf("TargetRole is not a supported scope for variable matching")
	}
	if len(definedScope.Tenants) > 0 {
		return false, nil, fmt.Errorf("Tenant is not a supported scope for variable matching")
	}
	if len(definedScope.Users) > 0 {
		return false, nil, fmt.Errorf("User is not a supported scope for variable matching")
	}

	var matchedScopes VariableScope
	var matched bool

	//If there is no scope to filter on return all the results
	if len(definedScope.Environments) > 0 && len(definedScope.Roles) > 0 && len(definedScope.Machines) > 0 && len(definedScope.Actions) > 0 && len(definedScope.Channels) > 0 && len(definedScope.TenantTags) > 0 {
		return true, &VariableScope{}, nil
	}

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

var _ IService = &variableService{}
