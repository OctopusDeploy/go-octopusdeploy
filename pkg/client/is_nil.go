package client

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *RootResource:
		return v == nil
	default:
		return v == nil
	}
}
