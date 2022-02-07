package services

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

const (
	emptyString      string = ""
	whitespaceString string = " "
)

// Service defines the contract for all services that communicate with the
// Octopus API.
type NamedServicer interface {
	GetName() string
	GetClient() octopusdeploy.Client
}

type Service struct {
	name string
	NamedServicer
}

func NewService(name string) *Service {
	return &Service{
		name: name,
	}
}

func (s *Service) GetName() string {
	return s.name
}

var _ NamedServicer = &Service{}
