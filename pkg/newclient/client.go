package newclient

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type Client interface {
	Sling() *sling.Sling
	URITemplateCache() *uritemplates.URITemplateCache
	// capabilities info would go here if the server supported capability-discovery
}

type client struct {
	// this should go away when we get rid of sling; it will be replaced with plain old httpClient
	sling *sling.Sling

	// Cache for parsed URI templates
	templateCache *uritemplates.URITemplateCache
}

func (n *client) Sling() *sling.Sling {
	return n.sling
}

func (n *client) URITemplateCache() *uritemplates.URITemplateCache {
	return n.templateCache
}

func NewClient(sling *sling.Sling) Client {
	return &client{
		sling:         sling,
		templateCache: uritemplates.NewUriTemplateCache(),
	}
}
