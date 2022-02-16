package service

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/google/go-querystring/query"
)

type CanGetByIDService[T resources.IResource] struct {
	IService
}

type GetsByIDer[T resources.IResource] interface {
	GetByID(id string) (*T, error)
}

type ResourceQueryer[Q any, T resources.IResource] interface {
	Query(queryStruct Q, pageSize *int) (IPagedResultsHandler[T], error)
}

type CanQueryService[Q any, T resources.IResource] struct {
	IService
}

func (s CanQueryService[Q, T]) Query(queryStruct Q, pageSize *int) (IPagedResultsHandler[T], error) {
	sizeOfPage := 30
	if pageSize != nil {
		sizeOfPage = *pageSize
	}
	urlValues, err := query.Values(s)
	if err != nil {
		return nil, err
	}
	basePathRelativeToRootWithQuery := fmt.Sprintf("%s?%s", s.GetBasePathRelativeToRoot(), urlValues.Encode())
	pageResultHandler := NewPagedResultsHandler[T](s.GetClient(), sizeOfPage, basePathRelativeToRootWithQuery)
	return pageResultHandler, nil
}

func (s CanGetByIDService[T]) GetByID(id string) (*T, error) {
	path := fmt.Sprintf("%s/%s", s.GetBasePathRelativeToRoot(), id)
	return ApiGet[T](s.GetClient(), path)
}
