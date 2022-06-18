package channels

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Channel:
		return v == nil
	case *Channels:
		return v == nil
	case *DeploymentActionPackage:
		return v == nil
	default:
		return v == nil
	}
}
