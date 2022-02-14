package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

type CanAddService[T octopusdeploy.IResource] struct {
	ResourceAdder[T]
}

type ResourceAdder[T octopusdeploy.IResource] interface {
	Add(resource *T) (*T, error)
	IService
}

func (s *CanAddService[T]) Add(resource *T) (*T, error) {
	if resource == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(octopusdeploy.OperationAdd, octopusdeploy.ParameterResource)
	}

	response, err := octopusdeploy.ApiAdd[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}
