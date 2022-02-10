package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

const (
	emptyString      string = ""
	whitespaceString string = " "
)

// IService defines the contract for all services that communicate with the
// Octopus API.
type IService interface {
	GetBasePathRelativeToRoot() string
	GetName() string
	GetClient() *octopusdeploy.Client
}

type service struct {
	basePathRelativeToRoot string
	name string
	IService
}

func NewService(name string, basePathRelativeToRoot string) *service {
	return &service{
		name: name,
		basePathRelativeToRoot: basePathRelativeToRoot,
	}
}

func (s *service) GetName() string {
	return s.name
}

var _ IService = &service{}
