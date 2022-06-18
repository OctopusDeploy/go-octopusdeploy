package triggers

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *ProjectTrigger:
		return v == nil
	default:
		return v == nil
	}
}
