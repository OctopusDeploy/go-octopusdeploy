package environments

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Environment:
		return v == nil
	default:
		return v == nil
	}
}
