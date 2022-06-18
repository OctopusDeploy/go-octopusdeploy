package workerpools

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type WorkerPoolResource struct {
	CanAddWorkers  bool           `json:"CanAddWorkers"`
	Description    string         `json:"Description,omitempty"`
	IsDefault      bool           `json:"IsDefault"`
	Name           string         `json:"Name" validate:"required,notblank"`
	SpaceID        string         `json:"SpaceId,omitempty" validate:"omitempty,notblank"`
	SortOrder      int            `json:"SortOrder"`
	WorkerPoolType WorkerPoolType `json:"WorkerPoolType"`
	WorkerType     WorkerType     `json:"WorkerType,omitempty"`

	resources.Resource
}

type WorkerPoolResources struct {
	Items []*WorkerPoolResource `json:"Items"`
	resources.PagedResults
}

// NewWorkerPoolResource creates and initializes a worker pool resource.
func NewWorkerPoolResource(name string, workerPoolType WorkerPoolType) *WorkerPoolResource {
	return &WorkerPoolResource{
		Name:           name,
		WorkerPoolType: workerPoolType,
		Resource:       *resources.NewResource(),
	}
}

func (w *WorkerPoolResource) GetCanAddWorkers() bool {
	return w.CanAddWorkers
}

func (w *WorkerPoolResource) GetDescription() string {
	return w.Description
}

func (w *WorkerPoolResource) GetIsDefault() bool {
	return w.IsDefault
}

// GetName returns the name of the worker pool resource.
func (w *WorkerPoolResource) GetName() string {
	return w.Name
}

func (w *WorkerPoolResource) GetSpaceID() string {
	return w.SpaceID
}

func (w *WorkerPoolResource) GetSortOrder() int {
	return w.SortOrder
}

// GetWorkerPoolType returns the worker type for this worker pool resource.
func (w *WorkerPoolResource) GetWorkerPoolType() WorkerPoolType {
	return w.WorkerPoolType
}

// GetWorkerType returns the worker type for this worker pool resource.
func (w *WorkerPoolResource) GetWorkerType() WorkerType {
	return w.WorkerType
}

func (w *WorkerPoolResource) SetCanAddWorkers(canAddWorkers bool) {
	w.CanAddWorkers = canAddWorkers
}

func (w *WorkerPoolResource) SetDescription(description string) {
	w.Description = description
}

func (w *WorkerPoolResource) SetIsDefault(isDefault bool) {
	w.IsDefault = isDefault
}

// SetName sets the name of the worker pool resource.
func (w *WorkerPoolResource) SetName(name string) {
	w.Name = name
}

func (w *WorkerPoolResource) SetSpaceID(spaceID string) {
	w.SpaceID = spaceID
}

func (w *WorkerPoolResource) SetSortOrder(sortOrder int) {
	w.SortOrder = sortOrder
}

func (w *WorkerPoolResource) SetWorkerPoolType(workerPoolType WorkerPoolType) {
	w.WorkerPoolType = workerPoolType
}

func (w *WorkerPoolResource) SetWorkerType(workerType WorkerType) {
	w.WorkerType = workerType
}

// Validate checks the state of the worker pool resource and returns an error
// if invalid.
func (w *WorkerPoolResource) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(w)
}

var _ IDynamicWorkerPool = &WorkerPoolResource{}
