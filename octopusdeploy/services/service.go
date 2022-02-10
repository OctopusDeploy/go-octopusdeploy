package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

const (
	emptyString      string = ""
	whitespaceString string = " "
)

// IService defines the contract for all services that communicate with the
// Octopus API.
type IService interface {
	GetName() string
	GetClient() *octopusdeploy.Client
}

type service struct {
	name string
	IService
}

func NewService(name string) *service {
	return &service{
		name: name,
	}
}

func (s *service) GetName() string {
	return s.name
}

var _ IService = &service{}
