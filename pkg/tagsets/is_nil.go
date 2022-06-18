package tagsets

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *TagSet:
		return v == nil
	default:
		return v == nil
	}
}
