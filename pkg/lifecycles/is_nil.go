package lifecycles

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Lifecycle:
		return v == nil
	default:
		return v == nil
	}
}
