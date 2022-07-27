package variables

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type ScriptModule struct {
	Description   string `json:"Description,omitempty"`
	Name          string `json:"Name" validate:"required"`
	ScriptBody    string `json:"scriptBody" validate:"required"`
	SpaceID       string `json:"SpaceId,omitempty"`
	Syntax        string `json:"syntax" validate:"required"`
	VariableSetID string `json:"VariableSetId,omitempty"`

	resources.Resource
}

// NewScriptModule creates and initializes a script module.
func NewScriptModule(name string) *ScriptModule {
	return &ScriptModule{
		Name: name,
	}
}

// MarshalJSON returns a script module as its JSON encoding.
func (s *ScriptModule) MarshalJSON() ([]byte, error) {
	scriptModule := struct {
		ContentType   string
		Description   string `json:"Description,omitempty"`
		Name          string `json:"Name" validate:"required"`
		ScriptBody    string `json:"scriptBody" validate:"required"`
		SpaceID       string `json:"SpaceId,omitempty"`
		Syntax        string `json:"syntax" validate:"required"`
		VariableSetID string `json:"VariableSetId,omitempty"`

		resources.Resource
	}{
		ContentType:   "ScriptModule",
		Description:   s.Description,
		Name:          s.Name,
		ScriptBody:    s.ScriptBody,
		SpaceID:       s.SpaceID,
		Syntax:        s.Syntax,
		VariableSetID: s.VariableSetID,

		Resource: s.Resource,
	}

	return json.Marshal(scriptModule)
}

// Validate checks the state of the script module and returns an error if
// invalid.
func (s *ScriptModule) Validate() error {
	return validator.New().Struct(s)
}
