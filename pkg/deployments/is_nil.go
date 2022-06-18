package deployments

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *Deployment:
		return v == nil
	case *DeploymentProcess:
		return v == nil
	case *DeploymentStep:
		return v == nil
	default:
		return v == nil
	}
}
