package services

type CanAddService[T IResource] struct {
	ResourceAdder[T]
}

type ResourceAdder[T IResource] interface {
	Add(resource *T) (*T, error)
	IService
}

func (s *CanAddService[T]) Add(resource *T) (*T, error) {
	if resource == nil {
		return nil, CreateInvalidParameterError(OperationAdd, ParameterResource)
	}

	response, err := ApiAdd[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}
