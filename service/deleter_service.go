package service

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

type CanDeleteService[T resource.IResource] struct {
	DeleteByIDer[T]
}

type DeleteByIDer[T resource.IResource] interface {
	DeleteByID(id string) error
	IService
}

// DeleteByID deletes the Resource that matches the input ID.
func (s CanDeleteService[T]) DeleteByID(id string) error {
	err := ApiDelete[T](s.GetClient(), id, s.GetBasePathRelativeToRoot())
	if err == ErrItemNotFound {
		return err
	}

	return err
}
