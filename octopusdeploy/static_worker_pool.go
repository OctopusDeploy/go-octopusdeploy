package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type StaticWorkerPool struct {
	WorkerPoolType string `json:"WorkerPoolType" validate:"required,eq=StaticWorkerPool"`

	WorkerPoolResource
}

type StaticWorkerPools struct {
	Items []*StaticWorkerPool `json:"Items"`
	PagedResults
}

// NewStaticWorkerPool creates and initializes a static worker pool.
func NewStaticWorkerPool(name string) *StaticWorkerPool {
	return &StaticWorkerPool{
		WorkerPoolType:     "StaticWorkerPool",
		WorkerPoolResource: *newWorkerPoolResource(name),
	}
}

// GetWorkerPoolType returns the worker type for this worker pool.
func (s *StaticWorkerPool) GetWorkerPoolType() string {
	return s.WorkerPoolType
}

// Validate checks the state of the static worker pool and returns an error if
// invalid.
func (s *StaticWorkerPool) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(s)
}

var _ IWorkerPool = &StaticWorkerPool{}
