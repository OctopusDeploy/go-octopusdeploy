package variables

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
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

// GetAll fetches a collection of variables for an owner ID.
func (s *VariableService) GetAll(ownerID string) (VariableSet, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return VariableSet{}, err
	}

	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	path := internal.TrimTemplate(s.GetPath())
	path = fmt.Sprintf(path+"/variableset-%s", ownerID)

	resp, err := api.ApiGet(s.GetClient(), new(VariableSet), path)
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
		if strings.EqualFold(variable.Name, name) {
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

// ----- new -------

// GetVariableSet returns the variable set with matching ID for a given space ID.
// This might be a project level variable set (ID might be 'variableset-Projects-314')
// or it might be a release snapshot variable set (ID might be 'variableset-Projects-314-s-2-XM74V')
func GetVariableSet(client newclient.Client, spaceID string, ID string) (*VariableSet, error) {
	if client == nil {
		return nil, internal.CreateInvalidParameterError("GetVariableSet", "client")
	}
	if spaceID == "" {
		return nil, internal.CreateInvalidParameterError("GetVariableSet", "project")
	}
	if ID == "" {
		return nil, internal.CreateInvalidParameterError("GetVariableSet", "ID")
	}

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.Variables, map[string]any{
		"spaceId": spaceID,
		"id":      ID,
	})
	if err != nil {
		return nil, err
	}
	return newclient.Get[VariableSet](client.HttpSession(), expandedUri)
}

// These methods below will be replacing the old defined in th service in this file. Soon.

// GetByID fetches a single variable, located by its ID, from Octopus Deploy for a given space ID and owner ID.
func GetByID(client newclient.Client, spaceID string, ownerID string, variableID string) (*Variable, error) {
	if internal.IsEmpty(ownerID) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	if internal.IsEmpty(variableID) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "variableID"}
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	variables, err := GetAll(client, spaceID, ownerID)
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
func GetByName(client newclient.Client, spaceID string, ownerID string, name string, scope *VariableScope) ([]*Variable, error) {
	if internal.IsEmpty(ownerID) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	if internal.IsEmpty(name) {
		return nil, errInvalidVariableServiceParameter{ParameterName: "name"}
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return nil, err
	}

	variables, err := GetAll(client, spaceID, ownerID)
	if err != nil {
		return nil, err
	}

	var matchedVariables []*Variable

	for _, variable := range variables.Variables {
		if strings.EqualFold(variable.Name, name) {
			matchScope, _, err := MatchesScope(variable.Scope, scope)
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

// GetAll fetches a collection of variables for an owner ID.
func GetAll(client newclient.Client, spaceID string, ownerID string) (VariableSet, error) {
	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return VariableSet{}, err
	}

	id := fmt.Sprintf("variableset-%s", ownerID)

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.Variables, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})
	if err != nil {
		return VariableSet{}, err
	}

	response, err := newclient.Get[VariableSet](client.HttpSession(), expandedUri)
	if err != nil {
		return VariableSet{}, err
	}

	return *response, nil
}

// AddSingle adds a single variable to an owner ID. This automates the act of fetching
// the variable set, adding a new item to it, and posting back to Octopus
func AddSingle(client newclient.Client, spaceID string, ownerID string, variable *Variable) (VariableSet, error) {
	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return VariableSet{}, err
	}

	variables, err := GetAll(client, spaceID, ownerID)
	if err != nil {
		return VariableSet{}, err
	}

	variables.Variables = append(variables.Variables, variable)
	return Update(client, spaceID, ownerID, variables)
}

// UpdateSingle adds a single variable to an owner ID. This automates the act of fetching
// the variable set, updating the existing item, and posting back to Octopus
func UpdateSingle(client newclient.Client, spaceID string, ownerID string, variable *Variable) (VariableSet, error) {
	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	variables, err := GetAll(client, spaceID, ownerID)
	if err != nil {
		return VariableSet{}, err
	}

	for k, v := range variables.Variables {
		if v.GetID() == variable.ID {
			variables.Variables[k] = variable
			return Update(client, spaceID, ownerID, variables)
		}
	}

	return VariableSet{}, services.ErrItemNotFound
}

// Update takes an entire variable set and posts the entire set back to Octopus Deploy. There are individual
// functions like AddSingle and UpdateSingle that can make this process more of a "typical" CRUD Octopus command.
func Update(client newclient.Client, spaceID string, ownerID string, variableSet VariableSet) (VariableSet, error) {
	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return VariableSet{}, err
	}

	id := fmt.Sprintf("variableset-%s", ownerID)

	expandedUri, err := client.URITemplateCache().Expand(uritemplates.Variables, map[string]any{
		"spaceId": spaceID,
		"id":      id,
	})
	if err != nil {
		return VariableSet{}, err
	}

	if _, err := newclient.Put[VariableSet](client.HttpSession(), expandedUri, variableSet); err != nil {
		return VariableSet{}, err
	}

	// 2021-04-22 (John Bristowe): we need to retrieve the variable set (again)
	// via HTTP GET (below) due to a bug for HTTP POST and HTTP PUT which will
	// provide a null scope value set in their responses

	return GetAll(client, spaceID, ownerID)
}

// MatchesScope compares two different scopes to see if they match. Generally used for comparing the scope of
// an existing variable against a desired state. Only supports Environment, Role, Machine, Action and Channel
// for scope options. Returns true if definedScope is nil or all elements are empty. Also returns a VariableScope
// of all the scopes that were matched
func MatchesScope(variableScope VariableScope, definedScope *VariableScope) (bool, *VariableScope, error) {
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

// DeleteSingle removes a single variable from an owner ID. This automates the act of fetching
// the variable set, removing the existing item, and posting back to Octopus
func DeleteSingle(client newclient.Client, spaceID string, ownerID string, variableID string) (VariableSet, error) {
	if internal.IsEmpty(ownerID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "ownerID"}
	}

	if internal.IsEmpty(variableID) {
		return VariableSet{}, errInvalidVariableServiceParameter{ParameterName: "variableID"}
	}

	spaceID, err := internal.GetSpaceID(spaceID, client.GetSpaceID())
	if err != nil {
		return VariableSet{}, err
	}

	variableSet, err := GetAll(client, spaceID, ownerID)
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

	return Update(client, spaceID, ownerID, variableSet)
}
