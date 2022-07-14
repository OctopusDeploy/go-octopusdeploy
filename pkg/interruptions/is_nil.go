package interruptions

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Interruption:
		return v == nil
	default:
		return v == nil
	}
}
