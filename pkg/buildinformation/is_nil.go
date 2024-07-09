package buildinformation

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *CreateBuildInformationCommand:
		return v == nil
	default:
		return v == nil
	}
}
