package client

type ServiceInterface interface {
	validateInternalState() error
}
