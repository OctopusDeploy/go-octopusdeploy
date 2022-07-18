package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type DeploymentProcessTemplate struct {
	DeploymentProcessId            string                            `json:"DeploymentProcessId,omitempty"`
	LastReleaseVersion             string                            `json:"LastReleaseVersion,omitempty"`
	NextVersionIncrement           string                            `json:"NextVersionIncrement,omitempty"`
	VersioningPackageStepName      *string                           `json:"VersioningPackageStepName,omitempty"`
	VersioningPackageReferenceName *string                           `json:"VersioningPackageReferenceName,omitempty"`
	Packages                       []releases.ReleaseTemplatePackage `json:"Packages,omitempty"`

	resources.Resource
}
