package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

type IResources[T any] interface {
	GetNextPage(client *sling.Sling) (*Resources[T], error)
	GetAllPages(client *sling.Sling) ([]T, error)
}

type Resources[T any] struct {
	Items []T `json:"Items"`
	PagedResults
}

// GetNextPage retrives the next page from the links collection. If no next page
// exists it will return nill
func (r *Resources[T]) GetNextPage(client *sling.Sling) (*Resources[T], error) {
	if r.Links.PageNext == "" {
		return nil, nil
	}
	response, err := api.ApiGet(client, new(Resources[T]), r.Links.PageNext)
	if err != nil {
		return nil, err
	}
	return response.(*Resources[T]), nil
}

// GetAllPages will retrive all remaining next pages in the link collection
// and return the result as list of concatenated Items; Including the items
// from the base Resource.
func (r *Resources[T]) GetAllPages(client *sling.Sling) ([]T, error) {
	items := make([]T, 0)
	res := r
	var err error
	for res != nil {
		items = append(items, res.Items...)
		res, err = res.GetNextPage(client)
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}

var _ IResources[any] = &Resources[any]{}
