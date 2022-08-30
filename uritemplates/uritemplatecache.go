package uritemplates

import (
	"errors"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type URITemplateCache struct {
	cache map[string]*UriTemplate
}

func NewUriTemplateCache() *URITemplateCache {
	return &URITemplateCache{
		cache: map[string]*UriTemplate{},
	}
}

func (c *URITemplateCache) Intern(uriTemplate string) (*UriTemplate, error) {
	cachedTemplate, ok := c.cache[uriTemplate]
	if ok {
		return cachedTemplate, nil
	}

	template, err := Parse(uriTemplate)
	if err != nil {
		return nil, err
	}
	c.cache[uriTemplate] = template
	return template, nil
}

func (c *URITemplateCache) Expand(uriTemplate string, value any) (string, error) {
	template, err := c.Intern(uriTemplate)
	if err != nil {
		return "", err
	}
	return template.Expand(value)
}

func (c *URITemplateCache) ExpandLinked(resource *resources.Resource, linkKey string, value any) (string, error) {
	rawTemplate, ok := resource.Links[linkKey]
	if !ok {
		return "", errors.New("ExpandLinked could not find linkKey")
	}
	template, err := c.Intern(rawTemplate)
	if err != nil {
		return "", err
	}
	return template.Expand(value)
}
