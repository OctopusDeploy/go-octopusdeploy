package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type StaticWorkerPool struct {
	WorkerPoolType WorkerPoolType `json:"WorkerPoolType"`

	WorkerPool
}

type StaticWorkerPools struct {
	Items []*StaticWorkerPool `json:"Items"`
	PagedResults
}

// NewStaticWorkerPool creates and initializes a static worker pool.
func NewStaticWorkerPool(name string) *StaticWorkerPool {
	return &StaticWorkerPool{
		WorkerPoolType: WorkerPoolTypeStatic,
		WorkerPool:     *newWorkerPool(name),
	}
}

func (s *StaticWorkerPool) GetIsDefault() bool {
	return s.IsDefault
}

// GetWorkerPoolType returns the worker type for this worker pool.
func (s *StaticWorkerPool) GetWorkerPoolType() WorkerPoolType {
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
