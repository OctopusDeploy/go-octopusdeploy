package service

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
)

type CanUpdateService[T resources.IResource] struct {
	ResourceUpdater[T]
}

type ResourceUpdater[T resources.IResource] interface {
	Update(resource *T) (*T, error)
	IService
}

func (s CanUpdateService[T]) Update(resource *T) (*T, error) {
	if resource == nil {
		return nil, internal.CreateInvalidParameterError(OperationUpdate, octopusdeploy.ParameterResource)
	}

	response, err := ApiUpdate[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}
