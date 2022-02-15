package service

type AdminService struct {
	adminClient *AdminClient
	service
	Adminer
}

func NewAdminService(name string, basePathRelativeToRoot string, client *AdminClient) AdminService {
	return AdminService{
		service:     *NewService(name, basePathRelativeToRoot),
		adminClient: client,
	}
}

func (s AdminService) GetClient() IClient {
	return s.adminClient
}
