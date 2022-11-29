package workerpools

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// IWorkerPool defines the interface for worker pools.
type IWorkerPool interface {
	GetCanAddWorkers() bool
	GetDescription() string
	GetIsDefault() bool
	GetName() string
	GetSpaceID() string
	GetSortOrder() int
	GetWorkerPoolType() WorkerPoolType
	SetCanAddWorkers(bool)
	SetDescription(string)
	SetIsDefault(bool)
	SetName(string)
	SetSpaceID(string)
	SetSortOrder(int)
	SetWorkerPoolType(WorkerPoolType)

	resources.IHasName
	resources.IResource
}

type WorkerPoolListResult struct {
	ID             string         `json:"Id,omitempty"`
	Name           string         `json:"Name" validate:"required,notblank"`
	WorkerPoolType WorkerPoolType `json:"WorkerPoolType"`
	IsDefault      bool           `json:"IsDefault"`
	CanAddWorkers  bool           `json:"CanAddWorkers"`
	Slug           string         `json:"Slug"`
}

// workerPool is the embedded struct used for all worker pools.
type workerPool struct {
	CanAddWorkers  bool           `json:"CanAddWorkers"`
	Description    string         `json:"Description,omitempty"`
	IsDefault      bool           `json:"IsDefault"`
	Name           string         `json:"Name" validate:"required,notblank"`
	SpaceID        string         `json:"SpaceId,omitempty" validate:"omitempty,notblank"`
	SortOrder      int            `json:"SortOrder"`
	WorkerPoolType WorkerPoolType `json:"WorkerPoolType"`

	resources.Resource
}

type WorkerPools struct {
	Items []IWorkerPool `json:"Items"`
	resources.PagedResults
}

// newWorkerPool creates and initializes a worker pool resource.
func newWorkerPool(name string, workerPoolType WorkerPoolType) *workerPool {
	return &workerPool{
		CanAddWorkers:  false,
		Name:           name,
		SortOrder:      0,
		WorkerPoolType: workerPoolType,
		Resource:       *resources.NewResource(),
	}
}

func (w *workerPool) GetCanAddWorkers() bool {
	return w.CanAddWorkers
}

func (w *workerPool) GetDescription() string {
	return w.Description
}

func (w *workerPool) GetIsDefault() bool {
	return w.IsDefault
}

// GetName returns the name of the worker pool resource.
func (w *workerPool) GetName() string {
	return w.Name
}

func (w *workerPool) GetSpaceID() string {
	return w.SpaceID
}

func (w *workerPool) GetSortOrder() int {
	return w.SortOrder
}

// GetWorkerPoolType returns the worker type for this worker pool.
func (w *workerPool) GetWorkerPoolType() WorkerPoolType {
	return w.WorkerPoolType
}

func (w *workerPool) SetCanAddWorkers(canAddWorkers bool) {
	w.CanAddWorkers = canAddWorkers
}

func (w *workerPool) SetDescription(description string) {
	w.Description = description
}

func (w *workerPool) SetIsDefault(isDefault bool) {
	w.IsDefault = isDefault
}

// SetName sets the name of the worker pool.
func (w *workerPool) SetName(name string) {
	w.Name = name
}

func (w *workerPool) SetSpaceID(spaceID string) {
	w.SpaceID = spaceID
}

func (w *workerPool) SetSortOrder(sortOrder int) {
	w.SortOrder = sortOrder
}

func (w *workerPool) SetWorkerPoolType(workerPoolType WorkerPoolType) {
	w.WorkerPoolType = workerPoolType
}

// Validate checks the state of the worker pool and returns an error if
// invalid.
func (w *workerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(w)
}

var _ resources.IHasName = &workerPool{}
var _ IWorkerPool = &workerPool{}
