package releases

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Release:
		return v == nil
	case *ReleaseQuery:
		return v == nil
	case *ReleasesQuery:
		return v == nil
	default:
		return v == nil
	}
}
