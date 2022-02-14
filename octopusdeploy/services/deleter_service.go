package services

type CanDeleteService[T Resource] struct {
	DeleteByIDer[T]
}

type DeleteByIDer[T Resource] interface {
	DeleteByID(id string) error
	IService
}

// DeleteByID deletes the Resource that matches the input ID.
func (s *CanDeleteService[T]) DeleteByID(id string) error {
	err := ApiDelete[T](s.GetClient(), id, s.GetBasePathRelativeToRoot())
	if err == ErrItemNotFound {
		return err
	}

	return err
}
