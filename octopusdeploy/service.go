package octopusdeploy

const (
	emptyString      string = ""
	whitespaceString string = " "
)

// Service defines the contract for all services that communicate with the
// Octopus API.
type Service interface {
	getName() string
	getClient() Client
}

type service struct {
	Name     string
	itemType IResource
	Service
}

type canDeleteService struct {
	service
}

type AdminService interface {
	Service
}

type adminService struct {
	client      AdminClient
	service
	AdminService
}

type SpaceScopedService interface {
	Service
}

type spaceScopedService struct {
	client      SpaceScopedClient
	service
	SpaceScopedService
}

func newService(name string) service {
	return service{
		Name:     name,
	}
}

func newAdminService(name string, client AdminClient) adminService {
	return adminService{
		service:     newService(name),
		client:      client,
	}
}

func newSpaceScopedService(name string, client SpaceScopedClient) spaceScopedService {
	return spaceScopedService{
		service:     newService(name),
		client:      client,
	}
}

func (s adminService) getClient() AdminClient {
	return s.client
}

func (s spaceScopedService) getClient() SpaceScopedClient {
	return s.client
}

func (s service) getName() string {
	return s.Name
}

func (s service) deleteByID(id string) error {
	return s.getClient().apiDelete(id)
}

// DeleteByID deletes the resource that matches the input ID.
func (s *canDeleteService) DeleteByID(id string) error {
	err := s.deleteByID(id)
	if err == ErrItemNotFound {
		return err
	}

	return err
}

var _ Service = &service{}
