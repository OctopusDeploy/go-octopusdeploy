package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

type errInvalidvariableServiceParameter struct {
	parameterName string
}

func (e errInvalidvariableServiceParameter) Error() string {
	return fmt.Sprintf("variableService: invalid parameter, %s", e.parameterName)
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
		service:     newService(serviceVariableService, sling, uriTemplate, new(Variable)),
	}
}

// GetAll fetches an entire VariableSet from Octopus Deploy for a given Project ID.
func (s variableService) GetAll(projectID string) (*Variables, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	if isEmpty(projectID) {
		return nil, errInvalidvariableServiceParameter{parameterName: "projectID"}
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
		return nil, errInvalidvariableServiceParameter{parameterName: "projectID"}
	}

	if isEmpty(variableID) {
		return nil, errInvalidvariableServiceParameter{parameterName: "variableID"}
	}

	variables, err := s.GetAll(projectID)
	if err != nil {
		return nil, err
	}

	for _, variable := range variables.Variables {
		if variable.ID == variableID {
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
		return nil, errInvalidvariableServiceParameter{parameterName: "projectID"}
	}

	if isEmpty(name) {
		return nil, errInvalidvariableServiceParameter{parameterName: "name"}
	}

	if scope == nil {
		return nil, errInvalidvariableServiceParameter{parameterName: "scope"}
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
		return nil, errInvalidvariableServiceParameter{parameterName: "projectID"}
	}

	if variable == nil {
		return nil, errInvalidvariableServiceParameter{parameterName: "variable"}
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
		return nil, errInvalidvariableServiceParameter{parameterName: "projectID"}
	}

	if variable == nil {
		return nil, errInvalidvariableServiceParameter{parameterName: "variable"}
	}

	variables, err := s.GetAll(projectID)

	if err != nil {
		return nil, err
	}

	var found bool
	for i, existingVar := range variables.Variables {
		if existingVar.ID == variable.ID {
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
		return nil, errInvalidvariableServiceParameter{parameterName: "projectID"}
	}

	if isEmpty(variableID) {
		return nil, errInvalidvariableServiceParameter{parameterName: "variableID"}
	}

	variables, err := s.GetAll(projectID)
	if err != nil {
		return nil, err
	}

	var found bool
	for i, existingVar := range variables.Variables {
		if existingVar.ID == variableID {
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
		return nil, errInvalidvariableServiceParameter{parameterName: "projectID"}
	}

	if variableSet == nil {
		return nil, errInvalidvariableServiceParameter{parameterName: "variableSet"}
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
		return false, nil, errInvalidvariableServiceParameter{parameterName: "variableScope"}
	}

	//If the scope supplied is nil then match everything
	if definedScope == nil {
		return true, NewVariableScope(), nil
	}

	//Unsupported scopes
	if len(definedScope.Private) > 0 {
		return false, nil, fmt.Errorf("Private is not a supported scope for variable matching")
	}
	if len(definedScope.Project) > 0 {
		return false, nil, fmt.Errorf("Project is not a supported scope for variable matching")
	}
	if len(definedScope.TargetRole) > 0 {
		return false, nil, fmt.Errorf("TargetRole is not a supported scope for variable matching")
	}
	if len(definedScope.Tenant) > 0 {
		return false, nil, fmt.Errorf("Tenant is not a supported scope for variable matching")
	}
	if len(definedScope.User) > 0 {
		return false, nil, fmt.Errorf("User is not a supported scope for variable matching")
	}

	var matchedScopes VariableScope
	var matched bool

	//If there is no scope to filter on return all the results
	if len(definedScope.Environment) > 0 && len(definedScope.Role) > 0 && len(definedScope.Machine) > 0 && len(definedScope.Action) > 0 && len(definedScope.Channel) > 0 && len(definedScope.TenantTag) > 0 {
		return true, &VariableScope{}, nil
	}

	for _, e1 := range definedScope.Environment {
		for _, e2 := range variableScope.Environment {
			if e1 == e2 {
				matched = true
				matchedScopes.Environment = append(matchedScopes.Environment, e1)
			}
		}
	}

	for _, r1 := range definedScope.Role {
		for _, r2 := range variableScope.Role {
			if r1 == r2 {
				matched = true
				matchedScopes.Role = append(matchedScopes.Role, r1)
			}
		}
	}

	for _, m1 := range definedScope.Machine {
		for _, m2 := range variableScope.Machine {
			if m1 == m2 {
				matched = true
				matchedScopes.Machine = append(matchedScopes.Machine, m1)
			}
		}
	}

	for _, a1 := range definedScope.Action {
		for _, a2 := range variableScope.Action {
			if a1 == a2 {
				matched = true
				matchedScopes.Action = append(matchedScopes.Action, a1)
			}
		}
	}

	for _, c1 := range definedScope.Channel {
		for _, c2 := range variableScope.Channel {
			if c1 == c2 {
				matched = true
				matchedScopes.Channel = append(matchedScopes.Channel, c1)
			}
		}
	}

	for _, c1 := range definedScope.TenantTag {
		for _, c2 := range variableScope.TenantTag {
			if c1 == c2 {
				matched = true
				matchedScopes.TenantTag = append(matchedScopes.TenantTag, c1)
			}
		}
	}

	return matched, &matchedScopes, nil
}

var _ IService = &variableService{}
