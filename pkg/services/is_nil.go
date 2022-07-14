package services

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case resources.IResource:
		return v == nil
	default:
		return v == nil
	}
}
