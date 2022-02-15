package service

type service struct {
	basePathRelativeToRoot string
	name                   string
	client                 IClient
}

// IService defines the contract for all service that communicate with the
// Octopus API.
type IService interface {
	GetBasePathRelativeToRoot() string
	GetName() string
	GetClient() IClient
}

func NewService(name string, basePathRelativeToRoot string, client IClient) IService {
	return &service{
		name:                   name,
		basePathRelativeToRoot: basePathRelativeToRoot,
		client:                 client,
	}
}

func (s service) GetName() string {
	return s.name
}

func (s service) GetBasePathRelativeToRoot() string {
	return s.basePathRelativeToRoot
}

func (s service) GetClient() IClient {
	return s.client
}
