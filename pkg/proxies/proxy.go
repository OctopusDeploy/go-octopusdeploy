package proxies

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type Proxy struct {
	Name      string               `json:"Name" validate:"required,notblank"`
	SpaceID   string               `json:"SpaceId,omitempty"`
	Host      string               `json:"Host,omitempty"`
	Port      int                  `json:"Port,omitempty"`
	Username  string               `json:"Username,omitempty"`
	Password  *core.SensitiveValue `json:"Password,omitempty"`
	ProxyType string               `json:"ProxyType,omitempty" validate:"oneof=HTTP"`

	resources.Resource
}

func NewProxy(name string, hostname string, username string, password *core.SensitiveValue) *Proxy {
	return &Proxy{
		Name:      name,
		Host:      hostname,
		Port:      80,
		ProxyType: "HTTP",
		Username:  username,
		Password:  password,
		Resource:  *resources.NewResource(),
	}
}

func (p *Proxy) GetName() string {
	return p.Name
}

func (p *Proxy) SetName(name string) {
	p.Name = name
}

func (p *Proxy) GetSpaceID() string {
	return p.SpaceID
}

func (p *Proxy) SetSpaceID(spaceID string) {
	p.SpaceID = spaceID
}

func (p *Proxy) GetHost() string {
	return p.Host
}

func (p *Proxy) SetHost(host string) {
	p.Host = host
}

func (p *Proxy) GetPort() int {
	return p.Port
}

func (p *Proxy) SetPort(port int) {
	p.Port = port
}

func (p *Proxy) GetUsername() string {
	return p.Username
}

func (p *Proxy) SetUsername(username string) {
	p.Username = username
}

func (p *Proxy) GetPassword() *core.SensitiveValue {
	return p.Password
}

func (p *Proxy) SetPassword(password *core.SensitiveValue) {
	p.Password = password
}
