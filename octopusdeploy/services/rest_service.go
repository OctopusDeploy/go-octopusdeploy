package services

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

type GetsByIDer[T octopusdeploy.Resource] interface {
	GetByID(id string) (*T, error)
	IService
}

type ResourceAdder[T octopusdeploy.Resource] interface {
	Add(resource *T) (*T, error)
	IService
}

type ResourceUpdater[T octopusdeploy.Resource] interface {
	Update(resource *T) (*T, error)
	IService
}

type ResourceQueryer[T octopusdeploy.Resource] interface {
	Query(queryStruct interface{}, template *uritemplates.UriTemplate) (octopusdeploy.PagedResults[T], error)
	IService
}

//type RESTService[T octopusdeploy.Resource] struct {
//	service
//	GetsByIDer[T]
//	ResourceQueryer[T]
//	ResourceAdder[T]
//	ResourceUpdater[T]
//}

type DeleteByIDer[T octopusdeploy.Resource] interface {
	DeleteByID(id string) error
	IService
}

type CanGetByIDService[T octopusdeploy.Resource] struct {
	GetsByIDer[T]
}

type CanAddService[T octopusdeploy.Resource] struct {
	ResourceAdder[T]
}

type CanUpdateService[T octopusdeploy.Resource] struct {
	ResourceUpdater[T]
}

type CanDeleteService[T octopusdeploy.Resource] struct {
	DeleteByIDer[T]
}

type Adminer interface {
	IService
}

type AdminService struct {
	*octopusdeploy.AdminClient
	service
	Adminer
}

type SpaceScoper interface {
	IService
}

type SpaceScopedService struct {
	*octopusdeploy.SpaceScopedClient
	service
	SpaceScoper
}

func NewAdminService(name string, basePathRelativeToRoot string, client *octopusdeploy.AdminClient) AdminService {
	return AdminService{
		service:     *NewService(name, basePathRelativeToRoot),
		AdminClient: client,
	}
}

func NewSpaceScopedService(name string, basePathRelativeToRoot string, client *octopusdeploy.SpaceScopedClient) SpaceScopedService {
	return SpaceScopedService{
		service:           *NewService(name, basePathRelativeToRoot),
		SpaceScopedClient: client,
	}
}

func (s AdminService) GetClient() *octopusdeploy.Client {
	return &s.Client
}

func (s SpaceScopedService) GetClient() *octopusdeploy.Client {
	return &s.Client
}

func (s *CanGetByIDService[T]) ApiGetByID(id string) (*T, error) {
	path := fmt.Sprintf("%s/%s", s.GetBasePathRelativeToRoot(), id)
	return octopusdeploy.ApiGet[T](s.GetClient(), path)
}

func (s *CanAddService[T]) Add(resource *T) (*T, error) {
	if resource == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(octopusdeploy.OperationAdd, octopusdeploy.ParameterResource)
	}

	response, err := octopusdeploy.ApiAdd[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *CanUpdateService[T]) Update(resource *T) (*T, error) {
	if resource == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(octopusdeploy.OperationUpdate, octopusdeploy.ParameterResource)
	}

	response, err := octopusdeploy.ApiUpdate[T](s.GetClient(), resource, s.GetBasePathRelativeToRoot())
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteByID deletes the Resource that matches the input ID.
func (s *CanDeleteService[T]) DeleteByID(id string) error {
	err := octopusdeploy.ApiDelete[T](s.GetClient(), id, s.GetBasePathRelativeToRoot())
	if err == octopusdeploy.ErrItemNotFound {
		return err
	}

	return err
}
