package newclient

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"net/url"
)

type Client interface {
	HttpSession() *HttpSession

	Sling() *sling.Sling // TODO sling will be removed in favour of HttpSession()
	URITemplateCache() *uritemplates.URITemplateCache
	// capabilities info would go here if the server supported capability-discovery
}

type client struct {
	httpSession *HttpSession

	sling *sling.Sling // TODO sling will be removed at some point

	// Cache for parsed URI templates
	templateCache *uritemplates.URITemplateCache
}

func (n *client) HttpSession() *HttpSession {
	return n.httpSession
}

func (n *client) Sling() *sling.Sling {
	return n.sling
}

func (n *client) URITemplateCache() *uritemplates.URITemplateCache {
	return n.templateCache
}

func NewClient(httpSession *HttpSession, baseURL *url.URL, sling *sling.Sling) Client {
	return &client{
		httpSession:   httpSession,
		sling:         sling,
		templateCache: uritemplates.NewUriTemplateCache(),
	}
}
