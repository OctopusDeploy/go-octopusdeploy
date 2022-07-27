package channels

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Channel:
		return v == nil
	default:
		return v == nil
	}
}
