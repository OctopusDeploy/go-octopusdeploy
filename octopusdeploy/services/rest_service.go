package services

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

type GetsByIDer[T octopusdeploy.Resource] interface {
	GetByID(id string) (*T, error)
	NamedServicer
}

type ResourceAdder[T octopusdeploy.Resource] interface {
	Add(resource *T) (*T, error)
	NamedServicer
}

type ResourceUpdater[T octopusdeploy.Resource] interface {
	Update(resource *T) (*T, error)
	NamedServicer
}

type ResourceQueryer[T octopusdeploy.Resource] interface {
	Query(queryStruct interface{}, template *uritemplates.UriTemplate) (octopusdeploy.PagedResults[T], error)
	NamedServicer
}

//type RESTService[T octopusdeploy.Resource] struct {
//	Service
//	GetsByIDer[T]
//	ResourceQueryer[T]
//	ResourceAdder[T]
//	ResourceUpdater[T]
//}

type DeleteByIDer[T octopusdeploy.Resource] interface {
	DeleteByID(id string) error
	NamedServicer
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
	NamedServicer
}

type AdminService struct {
	*octopusdeploy.AdminClient
	Service
	Adminer
}

type SpaceScoper interface {
	NamedServicer
}

type SpaceScopedService struct {
	*octopusdeploy.SpaceScopedClient
	Service
	SpaceScoper
}

func NewAdminService(name string, client *octopusdeploy.AdminClient) AdminService {
	return AdminService{
		Service: *NewService(name),
		AdminClient: client,
	}
}

func NewSpaceScopedService(name string, client *octopusdeploy.SpaceScopedClient) SpaceScopedService {
	return SpaceScopedService{
		Service: *NewService(name),
		SpaceScopedClient: client,
	}
}

func (s AdminService) GetClient() *octopusdeploy.Client {
	return &s.Client
}

func (s SpaceScopedService) GetClient() *octopusdeploy.Client {
	return &s.Client
}

func (s *CanAddService[T]) Add(resource *T) (*T, error) {
	if resource == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(octopusdeploy.OperationAdd, octopusdeploy.ParameterResource)
	}

	response, err := octopusdeploy.ApiAdd[T](s.GetClient(), resource)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *CanUpdateService[T]) Update(resource *T) (*T, error) {
	if resource == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(octopusdeploy.OperationUpdate, octopusdeploy.ParameterResource)
	}

	response, err := octopusdeploy.ApiUpdate[T](s.GetClient(), resource)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteByID deletes the Resource that matches the input ID.
func (s *CanDeleteService[T]) DeleteByID(id string) error {
	err := octopusdeploy.ApiDelete[T](s.GetClient(), id)
	if err == octopusdeploy.ErrItemNotFound {
		return err
	}

	return err
}