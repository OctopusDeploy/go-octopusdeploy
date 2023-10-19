package newclient

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
)

type Client interface {
	HttpSession() *HttpSession
	URITemplateCache() *uritemplates.URITemplateCache
	GetSpaceID() string
	// capabilities info could go here if the server supported capability-discovery
}

type client struct {
	httpSession   *HttpSession
	templateCache *uritemplates.URITemplateCache
	spaceID       string
}

func (n *client) HttpSession() *HttpSession {
	return n.httpSession
}

func (n *client) URITemplateCache() *uritemplates.URITemplateCache {
	return n.templateCache
}

func (n *client) GetSpaceID() string {
	return n.spaceID
}

func NewClient(httpSession *HttpSession) Client {
	return &client{
		httpSession:   httpSession,
		templateCache: uritemplates.NewUriTemplateCache(),
	}
}

func NewClientS(httpSession *HttpSession, spaceID string) Client {
	return &client{
		httpSession:   httpSession,
		templateCache: uritemplates.NewUriTemplateCache(),
		spaceID:       spaceID,
	}
}
