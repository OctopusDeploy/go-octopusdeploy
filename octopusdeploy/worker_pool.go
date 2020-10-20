package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type WorkerPool struct {
	CanAddWorkers  bool   `json:"CanAddWorkers"`
	Description    string `json:"Description,omitempty"`
	IsDefault      bool   `json:"IsDefault"`
	Name           string `json:"Name" validate:"required,notblank"`
	SpaceID        string `json:"SpaceId,omitempty" validate:"omitempty,notblank"`
	SortOrder      int    `json:"SortOrder"`
	WorkerPoolType string `json:"WorkerPoolType" validate:"required,oneof=DynamicWorkerPool StaticWorkerPool"`
	WorkerType     string `json:"WorkerType" validate:"required,oneof=Ubuntu1804 UbuntuDefault Windows2016 Windows2019 WindowsDefault"`

	Resource
}

type WorkerPools struct {
	Items []*WorkerPool `json:"Items"`
	PagedResults
}

// newWorkerPool creates and initializes a worker pool.
func newWorkerPool(name string) *WorkerPool {
	return &WorkerPool{
		Name:     name,
		Resource: *newResource(),
	}
}

// GetName returns the name of the worker pool.
func (w *WorkerPool) GetName() string {
	return w.Name
}

// SetName sets the name of the worker pool.
func (w *WorkerPool) SetName(name string) {
	w.Name = name
}

// GetWorkerPoolType returns the worker type for this worker pool.
func (w *WorkerPool) GetWorkerPoolType() string {
	return w.WorkerPoolType
}

// GetWorkerType returns the worker type for this worker pool.
func (w *WorkerPool) GetWorkerType() string {
	return w.WorkerType
}

// Validate checks the state of the worker pool and returns an error
// if invalid.
func (w *WorkerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(w)
}

var _ IDynamicWorkerPool = &WorkerPool{}
