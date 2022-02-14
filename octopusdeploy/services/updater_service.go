package services

type CanUpdateService[T Resource] struct {
	ResourceUpdater[T]
}

type ResourceUpdater[T Resource] interface {
	Update(resource *T) (*T, error)
	IService
}

func (s *CanUpdateService[T]) Update(resource *T) (*T, error) {
	if resource == nil {
		return nil, CreateInvalidParameterError(OperationUpdate, ParameterResource)
	}

	response, err := ApiUpdate[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}
