package workerpools

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type IDynamicWorkerPool interface {
	GetWorkerType() string

	IWorkerPool
}

type DynamicWorkerPool struct {
	WorkerType string `json:"WorkerType"`

	workerPool
}

// NewDynamicWorkerPool creates and initializes a dynamic worker pool.
func NewDynamicWorkerPool(name string, workerType string) *DynamicWorkerPool {
	return &DynamicWorkerPool{
		WorkerType: workerType,
		workerPool: *newWorkerPool(name, WorkerPoolTypeDynamic),
	}
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
