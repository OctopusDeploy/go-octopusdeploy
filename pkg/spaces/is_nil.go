package spaces

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Space:
		return v == nil
	default:
		return v == nil
	}
}
