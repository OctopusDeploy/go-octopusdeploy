package resources

type WorkerPoolType string

const (
	WorkerPoolTypeDynamic WorkerPoolType = "DynamicWorkerPool"
	WorkerPoolTypeStatic  WorkerPoolType = "StaticWorkerPool"
)
