package projects

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Project:
		return v == nil
	default:
		return v == nil
	}
}
