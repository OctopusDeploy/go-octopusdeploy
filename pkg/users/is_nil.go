package users

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *APIKey:
		return v == nil
	case *User:
		return v == nil
	default:
		return v == nil
	}
}
