package service

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
)

type CanAddService[T resources.IResource] struct {
	ResourceAdder[T]
}

type ResourceAdder[T resources.IResource] interface {
	Add(resource *T) (*T, error)
	IService
}

func (s CanAddService[T]) Add(resource *T) (*T, error) {
	if resource == nil {
		return nil, internal.CreateInvalidParameterError(OperationAdd, octopusdeploy.ParameterResource)
	}

	response, err := ApiAdd[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}
