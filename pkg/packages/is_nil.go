package packages

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *DeploymentActionPackage:
		return v == nil
	case *Package:
		return v == nil
	default:
		return v == nil
	}
}
