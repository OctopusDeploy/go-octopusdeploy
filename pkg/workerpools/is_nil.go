package workerpools

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *WorkerPoolResource:
		return v == nil
	case IWorkerPool:
		return v == nil
	default:
		return v == nil
	}
}
