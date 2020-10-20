package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type DynamicWorkerPool struct {
	WorkerPoolType string `json:"WorkerPoolType" validate:"required,eq=DynamicWorkerPool"`
	WorkerType     string `json:"WorkerType" validate:"required,oneof=Ubuntu1804 UbuntuDefault Windows2016 Windows2019 WindowsDefault"`

	WorkerPoolResource
}

type DynamicWorkerPools struct {
	Items []*DynamicWorkerPool `json:"Items"`
	PagedResults
}

// NewDynamicWorkerPool creates and initializes a dynamic worker pool.
func NewDynamicWorkerPool(name string, workerType string) *DynamicWorkerPool {
	return &DynamicWorkerPool{
		WorkerPoolType:     "DynamicWorkerPool",
		WorkerType:         workerType,
		WorkerPoolResource: *newWorkerPoolResource(name),
	}
}

// GetWorkerPoolType returns the worker pool type for this worker pool.
func (d *DynamicWorkerPool) GetWorkerPoolType() string {
	return d.WorkerPoolType
}

// GetWorkerType returns the worker type for this worker pool.
func (d *DynamicWorkerPool) GetWorkerType() string {
	return d.WorkerType
}

// Validate checks the state of the dynamic worker pool and returns an error if
// invalid.
func (d *DynamicWorkerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(d)
}

var _ IDynamicWorkerPool = &DynamicWorkerPool{}
