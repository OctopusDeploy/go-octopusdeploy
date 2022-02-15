package service

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"

type CanDeleteService[T resources.IResource] struct {
	IService
}

type DeleteByIDer[T resources.IResource] interface {
	DeleteByID(id string) error
}

// DeleteByID deletes the Resource that matches the input ID.
func (s CanDeleteService[T]) DeleteByID(id string) error {
	err := ApiDelete[T](s.GetClient(), id, s.GetBasePathRelativeToRoot())
	if err == ErrItemNotFound {
		return err
	}

	return err
}
