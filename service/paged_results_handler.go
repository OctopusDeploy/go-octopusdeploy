package service

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/google/go-querystring/query"
)

type IPagedResultsHandler[T resources.IResource] interface {
	HasMorePages() bool
	GetPage(pageNumber int) (items []T, e error)
	NextPage() (items []T, e error)
}

type pagedResultsHandler[T resources.IResource] struct {
	currentPage            int
	pageSize               int
	basePathRelativeToRoot string
	totalResults           *int
	client                 IClient
}

func NewPagedResultsHandler[T resources.IResource](client IClient, pageSize int, basePathRelativeToRoot string) IPagedResultsHandler[T] {
	t := &pagedResultsHandler[T]{
		pageSize:               pageSize,
		currentPage:            0,
		basePathRelativeToRoot: basePathRelativeToRoot,
		client:                 client,
	}
	return t
}

func (t pagedResultsHandler[T]) HasMorePages() bool {
	if t.totalResults == nil {
		return true
	}
	return t.currentPage < *t.totalResults/t.pageSize
}

func (t pagedResultsHandler[T]) GetPage(pageNumber int) (items []T, e error) {
	t.currentPage = pageNumber
	return t.NextPage()
}

func (t pagedResultsHandler[T]) NextPage() (items []T, e error) {
	skipTakeQuery := &SkipTakeQuery{
		Skip: t.currentPage * t.pageSize,
		Take: t.pageSize,
	}
	urlValues, err := query.Values(skipTakeQuery)
	if err != nil {
		return nil, err
	}

	basePathRelativeToRootWithSkipTake := fmt.Sprintf("%s%s", t.basePathRelativeToRoot, urlValues.Encode())
	resp, err := ApiGetMany[T](t.client, basePathRelativeToRootWithSkipTake)
	if err != nil {
		return nil, err
	}

	t.currentPage = t.currentPage + 1
	t.totalResults = &resp.TotalResults

	return resp.Items, nil
}
