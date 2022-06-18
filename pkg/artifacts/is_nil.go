package artifacts

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Artifact:
		return v == nil
	default:
		return v == nil
	}
}
