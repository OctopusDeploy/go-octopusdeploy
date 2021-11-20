package octopusdeploy

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type ScriptModule struct {
	Description   string `json:"Description,omitempty"`
	Name          string `json:"Name" validate:"required"`
	SpaceID       string `json:"SpaceId,omitempty"`
	VariableSetID string `json:"VariableSetId,omitempty"`

	resource
}

type ScriptModules struct {
	Items []*ScriptModule `json:"Items"`
	PagedResults
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
		ContentType   string `json:"ContentType" validate:"required"`
		Description   string `json:"Description,omitempty"`
		Name          string `json:"Name" validate:"required"`
		SpaceID       string `json:"SpaceId,omitempty"`
		VariableSetID string `json:"VariableSetId,omitempty"`

		resource
	}{
		ContentType:   "ScriptModule",
		Description:   s.Description,
		Name:          s.Name,
		SpaceID:       s.SpaceID,
		VariableSetID: s.VariableSetID,

		resource: s.resource,
	}

	return json.Marshal(scriptModule)
}

// Validate checks the state of the script module and returns an error if
// invalid.
func (s *ScriptModule) Validate() error {
	return validator.New().Struct(s)
}
