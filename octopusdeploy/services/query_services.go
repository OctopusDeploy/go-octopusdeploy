package services

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

type CanGetByIDService[T octopusdeploy.Resource] struct {
	GetsByIDer[T]
}

type GetsByIDer[T octopusdeploy.Resource] interface {
	GetByID(id string) (*T, error)
	IService
}

type ResourceQueryer[T octopusdeploy.Resource] interface {
	Query(queryStruct interface{}, template *uritemplates.UriTemplate) (octopusdeploy.PagedResults[T], error)
	IService
}

func (s *CanGetByIDService[T]) GetByID(id string) (*T, error) {
	path := fmt.Sprintf("%s/%s", s.GetBasePathRelativeToRoot(), id)
	return octopusdeploy.ApiGet[T](s.GetClient(), path)
}
