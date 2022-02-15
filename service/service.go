package service

type service struct {
	basePathRelativeToRoot string
	name                   string
	IService
}

// IService defines the contract for all service that communicate with the
// Octopus API.
type IService interface {
	GetBasePathRelativeToRoot() string
	GetName() string
	GetClient() IClient
}

func NewService(name string, basePathRelativeToRoot string) *service {
	return &service{
		name:                   name,
		basePathRelativeToRoot: basePathRelativeToRoot,
	}
}

func (s *service) GetName() string {
	return s.name
}

func (s *service) GetBasePathRelativeToRoot() string {
	return s.basePathRelativeToRoot
}

var _ IService = &service{}
