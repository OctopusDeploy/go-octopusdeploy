package machines

// IRunsOnAWorker defines the interface for workers.
type IRunsOnAWorker interface {
	GetDefaultWorkerPoolID() string
	SetDefaultWorkerPoolID(string)
}
