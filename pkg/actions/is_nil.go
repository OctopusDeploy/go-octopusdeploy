package actions

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *CommunityActionTemplate:
		return v == nil
	default:
		return v == nil
	}
}
