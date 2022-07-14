package tasks

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Task:
		return v == nil
	default:
		return v == nil
	}
}
