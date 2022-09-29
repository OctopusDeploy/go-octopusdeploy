package newclient

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type Client interface {
	HttpSession() *HttpSession
	URITemplateCache() *uritemplates.URITemplateCache
	// capabilities info could go here if the server supported capability-discovery
}

type client struct {
	httpSession   *HttpSession
	templateCache *uritemplates.URITemplateCache
}

func (n *client) HttpSession() *HttpSession {
	return n.httpSession
}

func (n *client) URITemplateCache() *uritemplates.URITemplateCache {
	return n.templateCache
}

func NewClient(httpSession *HttpSession) Client {
	return &client{
		httpSession:   httpSession,
		templateCache: uritemplates.NewUriTemplateCache(),
	}
}
