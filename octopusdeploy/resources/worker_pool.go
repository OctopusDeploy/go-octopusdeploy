package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// WorkerPool is the embedded struct used for all worker pools.
type WorkerPool struct {
	CanAddWorkers  bool           `json:"CanAddWorkers"`
	Description    string         `json:"Description,omitempty"`
	IsDefault      bool           `json:"IsDefault"`
	Name           string         `json:"Name" validate:"required,notblank"`
	SpaceID        string         `json:"SpaceId,omitempty" validate:"omitempty,notblank"`
	SortOrder      int            `json:"SortOrder"`
	WorkerPoolType WorkerPoolType `json:"WorkerPoolType"`

	Resource
}

// newWorkerPool creates and initializes a worker pool Resource.
func newWorkerPool(name string, workerPoolType WorkerPoolType) *WorkerPool {
	return &WorkerPool{
		CanAddWorkers:  false,
		Name:           name,
		SortOrder:      0,
		WorkerPoolType: workerPoolType,
		Resource:       *NewResource(),
	}
}

// GetIsDefaults returns the default status of this worker pool.
func (w *WorkerPool) GetIsDefault() bool {
	return w.IsDefault
}

// GetName returns the name of the worker pool.
func (w *WorkerPool) GetName() string {
	return w.Name
}

// GetWorkerPoolType returns the worker type for this worker pool.
func (w *WorkerPool) GetWorkerPoolType() WorkerPoolType {
	return w.WorkerPoolType
}

// SetName sets the name of the worker pool.
func (w *WorkerPool) SetName(name string) {
	w.Name = name
}

// Validate checks the state of the worker pool and returns an error if
// invalid.
func (w *WorkerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(w)
}

var _ IHasName = &WorkerPool{}
var _ IWorkerPool = &WorkerPool{}
