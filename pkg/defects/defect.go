package defects

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type DefectStatus string

const (
	DefectStatusResolved   = "Resolved"
	DefectStatusUnresolved = "Unresolved"
)

type Defect struct {
	Description string       `json:"Description" validate:"required,notblank"`
	Status      DefectStatus `json:"Status,omitempty"`

	resources.Resource
}

type ResolveReleaseDefectCommand struct {
	ReleaseID string `json:"ReleaseId"`
}

func NewResolveReleaseDefectCommand(releaseID string) (*ResolveReleaseDefectCommand, error) {
	if internal.IsEmpty(releaseID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("releaseid")
	}

	command := ResolveReleaseDefectCommand{
		ReleaseID: releaseID,
	}

	return &command, nil
}

type CreateReleaseDefectCommand struct {
	ReleaseID   string `json:"ReleaseId"`
	Description string `json:"Description"`
}

func NewCreateReleaseDefectCommand(releaseID string, description string) (*CreateReleaseDefectCommand, error) {
	if internal.IsEmpty(releaseID) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("releaseid")
	}

	if internal.IsEmpty(description) {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError("description")
	}

	command := CreateReleaseDefectCommand{
		ReleaseID:   releaseID,
		Description: description,
	}

	return &command, nil
}
