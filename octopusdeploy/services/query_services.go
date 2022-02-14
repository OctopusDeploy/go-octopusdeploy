package services

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

type CanGetByIDService[T resources.IResource] struct {
	GetsByIDer[T]
}

type GetsByIDer[T resources.IResource] interface {
	GetByID(id string) (*T, error)
	IService
}

type ResourceQueryer[T resources.IResource] interface {
	Query(queryStruct interface{}, template *uritemplates.UriTemplate) (IPagedResultsHandler[T], error)
	IService
}

func (s *CanGetByIDService[T]) GetByID(id string) (*T, error) {
	path := fmt.Sprintf("%s/%s", s.GetBasePathRelativeToRoot(), id)
	return ApiGet[T](s.GetClient(), path)
}
