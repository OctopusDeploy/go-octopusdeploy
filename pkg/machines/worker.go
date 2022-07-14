package machines

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Worker struct {
	SpaceID       string   `json:"SpaceId,omitempty"`
	WorkerPoolIDs []string `json:"WorkerPoolIds,omitempty"`

	machine
}

// Workers defines a collection of workers with built-in support for paged
// results.
type Workers struct {
	Items []*Worker `json:"Items"`
	resources.PagedResults
}

func NewWorker(name string, endpoint IEndpoint) *Worker {
	worker := &Worker{
		WorkerPoolIDs: []string{},
		machine:       *newMachine(name, endpoint),
	}

	return worker
}

// GetSpaceID returns the space ID that is associated with this worker.
func (w *Worker) GetSpaceID() string {
	return w.SpaceID
}

// SetSpaceID sets the space ID that is associated with this worker.
func (w *Worker) SetSpaceID(spaceID string) {
	w.SpaceID = spaceID
}

// MarshalJSON returns a worker as its JSON encoding.
func (w *Worker) MarshalJSON() ([]byte, error) {
	worker := struct {
		SpaceID       string   `json:"SpaceId,omitempty"`
		WorkerPoolIDs []string `json:"WorkerPoolIds,omitempty"`
		machine
	}{
		SpaceID:       w.SpaceID,
		WorkerPoolIDs: w.WorkerPoolIDs,
		machine:       w.machine,
	}

	return json.Marshal(worker)
}

// UnmarshalJSON sets this worker to its representation in JSON.
func (w *Worker) UnmarshalJSON(b []byte) error {
	var fields struct {
		SpaceID       string   `json:"SpaceId,omitempty"`
		WorkerPoolIDs []string `json:"WorkerPoolIds,omitempty"`
	}
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	// validate incoming JSON
	validate := validator.New()
	err = validate.Struct(fields)
	if err != nil {
		return err
	}

	w.SpaceID = fields.SpaceID
	w.WorkerPoolIDs = fields.WorkerPoolIDs

	err = json.Unmarshal(b, &w.machine)
	if err != nil {
		return err
	}

	return nil
}

var _ resources.IHasSpace = &Worker{}
