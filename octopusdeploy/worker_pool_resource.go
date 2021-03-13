package octopusdeploy

import (
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

	resource
}

type WorkerPoolResources struct {
	Items []*WorkerPoolResource `json:"Items"`
	PagedResults
}

// newWorkerPoolResource creates and initializes a worker pool resource.
func newWorkerPoolResource(name string) *WorkerPoolResource {
	return &WorkerPoolResource{
		Name:     name,
		resource: *newResource(),
	}
}

// GetName returns the name of the worker pool resource.
func (w *WorkerPoolResource) GetName() string {
	return w.Name
}

// SetName sets the name of the worker pool resource.
func (w *WorkerPoolResource) SetName(name string) {
	w.Name = name
}

func (w *WorkerPoolResource) GetIsDefault() bool {
	return w.IsDefault
}

// GetWorkerPoolType returns the worker type for this worker pool resource.
func (w *WorkerPoolResource) GetWorkerPoolType() WorkerPoolType {
	return w.WorkerPoolType
}

// GetWorkerType returns the worker type for this worker pool resource.
func (w *WorkerPoolResource) GetWorkerType() WorkerType {
	return w.WorkerType
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
