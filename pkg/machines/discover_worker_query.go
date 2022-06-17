package machines

type DiscoverWorkerQuery struct {
	Host    string `uri:"host,omitempty" url:"host,omitempty"`
	Port    int    `uri:"port,omitempty" url:"port,omitempty"`
	ProxyID string `uri:"proxyId,omitempty" url:"proxyId,omitempty"`
	Type    string `uri:"type,omitempty" url:"type,omitempty"`
}
