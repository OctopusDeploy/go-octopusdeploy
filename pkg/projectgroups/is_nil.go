package projectgroups

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *ProjectGroup:
		return v == nil
	default:
		return v == nil
	}
}
