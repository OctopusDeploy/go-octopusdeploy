package deploymentfreezes

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"time"
)

type DeploymentFreezes resources.Resources[DeploymentFreeze]

type DeploymentFreeze struct {
	Name                    string              `json:"Name" validate:"required"`
	Start                   *time.Time          `json:"Start,required"`
	End                     *time.Time          `json:"End,required"`
	ProjectEnvironmentScope map[string][]string `json:"ProjectEnvironmentScope,omitempty"`

	resources.Resource
}

func (d *DeploymentFreeze) GetName() string {
	return d.Name
}

func (d *DeploymentFreeze) SetName(name string) {
	d.Name = name
}
