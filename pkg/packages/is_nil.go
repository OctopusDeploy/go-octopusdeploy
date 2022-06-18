package packages

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Package:
		return v == nil
	default:
		return v == nil
	}
}
