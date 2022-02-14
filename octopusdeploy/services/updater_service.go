package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

type CanUpdateService[T octopusdeploy.Resource] struct {
	ResourceUpdater[T]
}

type ResourceUpdater[T octopusdeploy.Resource] interface {
	Update(resource *T) (*T, error)
	IService
}

func (s *CanUpdateService[T]) Update(resource *T) (*T, error) {
	if resource == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(octopusdeploy.OperationUpdate, octopusdeploy.ParameterResource)
	}

	response, err := octopusdeploy.ApiUpdate[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}
