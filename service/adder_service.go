package service

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
)

type CanAddService[T resources.IResource] struct {
	IService
}

type ResourceAdder[T resources.IResource] interface {
	Add(resource T) (*T, error)
}

func (s CanAddService[T]) Add(resource T) (*T, error) {

	response, err := ApiAdd[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}
