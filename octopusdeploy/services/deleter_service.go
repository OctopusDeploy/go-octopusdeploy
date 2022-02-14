package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

type CanDeleteService[T octopusdeploy.Resource] struct {
	DeleteByIDer[T]
}

type DeleteByIDer[T octopusdeploy.Resource] interface {
	DeleteByID(id string) error
	IService
}

// DeleteByID deletes the Resource that matches the input ID.
func (s *CanDeleteService[T]) DeleteByID(id string) error {
	err := octopusdeploy.ApiDelete[T](s.GetClient(), id, s.GetBasePathRelativeToRoot())
	if err == octopusdeploy.ErrItemNotFound {
		return err
	}

	return err
}
