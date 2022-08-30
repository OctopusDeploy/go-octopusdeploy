package newclient

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
)

type Client interface {
	Sling() *sling.Sling
	SpaceID() string
	URITemplateCache() *uritemplates.URITemplateCache
}

type client struct {
	// this should go away when we get rid of sling; it will be replaced with plain old httpClient
	sling *sling.Sling

	// If the client is space-scoped, this will be set, and it must be valid or the client would fail to init; else empty
	spaceID string

	// Cache for parsed URI templates
	templateCache *uritemplates.URITemplateCache
}

func (n *client) Sling() *sling.Sling {
	return n.sling
}

func (n *client) SpaceID() string {
	return n.spaceID
}

func (n *client) URITemplateCache() *uritemplates.URITemplateCache {
	return n.templateCache
}

func NewClient(sling *sling.Sling, SpaceID string) Client {
	return &client{
		sling:         sling,
		spaceID:       SpaceID,
		templateCache: uritemplates.NewUriTemplateCache(),
	}
}
