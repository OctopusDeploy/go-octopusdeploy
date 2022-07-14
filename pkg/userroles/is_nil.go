package userroles

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *ScopedUserRole:
		return v == nil
	case *UserRole:
		return v == nil
	default:
		return v == nil
	}
}
