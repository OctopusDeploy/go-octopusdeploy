package service

type adminService struct {
	IService
}

type IAdminService interface {
	IService
}

func NewAdminService(name string, basePathRelativeToRoot string, client IAdminClient) IAdminService {
	return &adminService{
		IService: NewService(name, basePathRelativeToRoot, client),
	}
}
