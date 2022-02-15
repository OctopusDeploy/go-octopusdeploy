package service

type AdminService struct {
	AdminClient
	IService
}

type IAdminService interface {
	GetClient() IClient
	IService
}

func NewAdminService(name string, basePathRelativeToRoot string, client *AdminClient) IAdminService {
	return &AdminService{
		IService:    NewService(name, basePathRelativeToRoot),
		AdminClient: *client,
	}
}

func (s AdminService) GetClient() IClient {
	return s.Client
}
