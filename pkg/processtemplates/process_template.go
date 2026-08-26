package processtemplates

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
)

// Icon represents the icon shown against a process template.
type Icon struct {
	// ID is the Font Awesome icon identifier.
	ID string `json:"Id"`
	// Color is the icon background colour, as a hex string.
	Color string `json:"Color"`
}

// ProcessTemplate represents a Platform Hub process template, including its steps and parameters.
type ProcessTemplate struct {
	ID          string                        `json:"Id,omitempty"`
	Name        string                        `json:"Name"`
	GitRef      string                        `json:"GitRef"`
	Slug        string                        `json:"Slug"`
	Description string                        `json:"Description,omitempty"`
	Icon        *Icon                         `json:"Icon,omitempty"`
	Steps       []*deployments.DeploymentStep `json:"Steps,omitempty"`
	Parameters  []*Parameter                  `json:"Parameters,omitempty"`
}

// Parameter represents a parameter declared by a process template.
type Parameter struct {
	Name            string            `json:"Name"`
	Label           string            `json:"Label,omitempty"`
	HelpText        string            `json:"HelpText,omitempty"`
	IsOptional      bool              `json:"IsOptional"`
	DisplaySettings map[string]string `json:"DisplaySettings,omitempty"`
	Values          []*ParameterValue `json:"Values,omitempty"`
}

// ParameterValue represents a scoped default value for a process template parameter.
type ParameterValue struct {
	Value core.PropertyValue      `json:"Value"`
	Scope variables.VariableScope `json:"Scope"`
}

// Summary represents a process template without its steps or parameters. Listings should
// prefer this shape: the full template route returns every step body.
type Summary struct {
	ID            string     `json:"Id,omitempty"`
	Name          string     `json:"Name"`
	GitRef        string     `json:"GitRef"`
	Slug          string     `json:"Slug"`
	Description   string     `json:"Description,omitempty"`
	Icon          *Icon      `json:"Icon,omitempty"`
	Version       string     `json:"Version,omitempty"`
	PublishedDate *time.Time `json:"PublishedDate,omitempty"`
	HasError      bool       `json:"HasError"`
}
