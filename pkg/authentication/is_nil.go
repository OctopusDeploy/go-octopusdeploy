package authentication

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Authentication:
		return v == nil
	default:
		return v == nil
	}
}
