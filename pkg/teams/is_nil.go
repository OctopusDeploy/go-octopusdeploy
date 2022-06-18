package teams

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Team:
		return v == nil
	default:
		return v == nil
	}
}
