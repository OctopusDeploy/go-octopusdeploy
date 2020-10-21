package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// WorkerPoolResource is the embedded struct used for all worker pools.
type WorkerPoolResource struct {
	CanAddWorkers bool   `json:"CanAddWorkers"`
	Description   string `json:"Description,omitempty"`
	IsDefault     bool   `json:"IsDefault"`
	Name          string `json:"Name" validate:"required,notblank"`
	SpaceID       string `json:"SpaceId,omitempty" validate:"omitempty,notblank"`
	SortOrder     int    `json:"SortOrder"`

	resource
}

// newWorkerPoolResource creates and initializes a worker pool resource.
func newWorkerPoolResource(name string) *WorkerPoolResource {
	return &WorkerPoolResource{
		CanAddWorkers: false,
		Name:          name,
		SortOrder:     0,
		resource:      *newResource(),
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

var _ IHasName = &WorkerPoolResource{}
