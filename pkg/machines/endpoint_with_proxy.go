package machines

type IEndpointWithProxy interface {
	GetProxyID() string
	SetProxyID(string)
}
