package credentials

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Resource:
		return v == nil
	default:
		return v == nil
	}
}
