package resources

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

	Resource
}

type WorkerPoolResources struct {
	Items []*WorkerPoolResource `json:"Items"`
	PagedResults
}

// NewWorkerPoolResource creates and initializes a worker pool Resource.
func NewWorkerPoolResource(name string, workerPoolType WorkerPoolType) *WorkerPoolResource {
	return &WorkerPoolResource{
		Name:           name,
		WorkerPoolType: workerPoolType,
		Resource:       *NewResource(),
	}
}

// GetName returns the name of the worker pool Resource.
func (w *WorkerPoolResource) GetName() string {
	return w.Name
}

// SetName sets the name of the worker pool Resource.
func (w *WorkerPoolResource) SetName(name string) {
	w.Name = name
}

func (w *WorkerPoolResource) GetIsDefault() bool {
	return w.IsDefault
}

// GetWorkerPoolType returns the worker type for this worker pool Resource.
func (w *WorkerPoolResource) GetWorkerPoolType() WorkerPoolType {
	return w.WorkerPoolType
}

// GetWorkerType returns the worker type for this worker pool Resource.
func (w *WorkerPoolResource) GetWorkerType() WorkerType {
	return w.WorkerType
}

// Validate checks the state of the worker pool Resource and returns an error
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
