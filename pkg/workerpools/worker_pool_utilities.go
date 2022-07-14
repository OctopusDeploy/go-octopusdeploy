package workerpools

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
)

func ToWorkerPool(workerPoolResource *WorkerPoolResource) (IWorkerPool, error) {
	if IsNil(workerPoolResource) {
		return nil, internal.CreateInvalidParameterError("ToWorkerPool", "workerPoolResource")
	}

	if err := workerPoolResource.Validate(); err != nil {
		return nil, err
	}

	var workerPool IWorkerPool

	switch workerPoolResource.GetWorkerPoolType() {
	case WorkerPoolTypeDynamic:
		workerPool = NewDynamicWorkerPool(workerPoolResource.GetName(), workerPoolResource.GetWorkerType())
	case WorkerPoolTypeStatic:
		workerPool = NewStaticWorkerPool(workerPoolResource.GetName())
	}

	workerPool.SetCanAddWorkers((workerPoolResource.GetCanAddWorkers()))
	workerPool.SetDescription((workerPoolResource.GetDescription()))
	workerPool.SetID(workerPoolResource.GetID())
	workerPool.SetIsDefault((workerPoolResource.GetIsDefault()))
	workerPool.SetLinks(workerPoolResource.GetLinks())
	workerPool.SetModifiedBy(workerPoolResource.GetModifiedBy())
	workerPool.SetModifiedOn(workerPoolResource.GetModifiedOn())
	workerPool.SetName((workerPoolResource.GetName()))
	workerPool.SetSpaceID((workerPoolResource.GetSpaceID()))
	workerPool.SetSortOrder((workerPoolResource.GetSortOrder()))
	workerPool.SetWorkerPoolType((workerPoolResource.GetWorkerPoolType()))

	return workerPool, nil
}

func ToWorkerPools(workerPoolResources *WorkerPoolResources) *WorkerPools {
	return &WorkerPools{
		Items:        ToWorkerPoolArray(workerPoolResources.Items),
		PagedResults: workerPoolResources.PagedResults,
	}
}

func ToWorkerPoolResource(workerPool IWorkerPool) (*WorkerPoolResource, error) {
	if IsNil(workerPool) {
		return nil, internal.CreateInvalidParameterError("ToWorkerPoolResource", "workerPool")
	}

	// conversion unnecessary if input worker pool is *WorkerPoolResource
	if v, ok := workerPool.(*WorkerPoolResource); ok {
		return v, nil
	}

	workerPoolResource := NewWorkerPoolResource(workerPool.GetName(), workerPool.GetWorkerPoolType())

	switch workerPoolResource.GetWorkerPoolType() {
	case WorkerPoolTypeDynamic:
		dynamicWorkerPool := workerPool.(*DynamicWorkerPool)
		workerPoolResource.WorkerType = dynamicWorkerPool.GetWorkerType()
	case WorkerPoolTypeStatic:
		// nothing to copy
	}

	workerPoolResource.SetCanAddWorkers((workerPool.GetCanAddWorkers()))
	workerPoolResource.SetDescription((workerPool.GetDescription()))
	workerPoolResource.SetID(workerPool.GetID())
	workerPoolResource.SetIsDefault((workerPool.GetIsDefault()))
	workerPoolResource.SetLinks(workerPool.GetLinks())
	workerPoolResource.SetModifiedBy(workerPool.GetModifiedBy())
	workerPoolResource.SetModifiedOn(workerPool.GetModifiedOn())
	workerPoolResource.SetName((workerPool.GetName()))
	workerPoolResource.SetSpaceID((workerPool.GetSpaceID()))
	workerPoolResource.SetSortOrder((workerPool.GetSortOrder()))
	workerPoolResource.SetWorkerPoolType((workerPool.GetWorkerPoolType()))

	return workerPoolResource, nil
}

func ToWorkerPoolArray(workerPoolResources []*WorkerPoolResource) []IWorkerPool {
	items := []IWorkerPool{}
	for _, workerPoolResource := range workerPoolResources {
		workerPool, err := ToWorkerPool(workerPoolResource)
		if err != nil {
			return nil
		}
		items = append(items, workerPool)
	}
	return items
}
