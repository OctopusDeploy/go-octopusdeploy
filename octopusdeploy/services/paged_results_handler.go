package services

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
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
	client                 octopusdeploy.IClient
	IPagedResultsHandler[T]
}

func NewPagedResultsHandler[T resources.IResource](client octopusdeploy.IClient, pageSize int, basePathRelativeToRoot string) IPagedResultsHandler[T] {
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
	//skipTakeQuery := &SkipTakeQuery{
	//	Skip: t.currentPage * t.pageSize,
	//	Take: t.pageSize,
	//}

	//TODO: include skip/take params in the basePathRelativeToRoot
	resp, err := octopusdeploy.ApiGetMany[T](t.client, t.basePathRelativeToRoot)
	if err != nil {
		return nil, err
	}

	t.currentPage = t.currentPage + 1
	t.totalResults = &resp.TotalResults

	return resp.Items, nil
}