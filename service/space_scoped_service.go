package service

type spaceScopedService struct {
	IService
}

type ISpaceScopedService interface {
	IService
}

func NewSpaceScopedService(name string, basePathRelativeToRoot string, client ISpaceScopedClient) ISpaceScopedService {
	return &spaceScopedService{
		IService: NewService(name, basePathRelativeToRoot, client),
	}
}
