package services

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

type CanGetByIDService[T Resource] struct {
	GetsByIDer[T]
}

type GetsByIDer[T Resource] interface {
	GetByID(id string) (*T, error)
	IService
}

type ResourceQueryer[T Resource] interface {
	Query(queryStruct interface{}, template *uritemplates.UriTemplate) (PagedResults[T], error)
	IService
}

func (s *CanGetByIDService[T]) GetByID(id string) (*T, error) {
	path := fmt.Sprintf("%s/%s", s.GetBasePathRelativeToRoot(), id)
	return ApiGet[T](s.GetClient(), path)
}
