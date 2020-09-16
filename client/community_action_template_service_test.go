package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewCommunityActionTemplateServiceWithNil(t *testing.T) {
	service := NewCommunityActionTemplateService(nil, "")
	assert.Nil(t, service)
}

func TestCommunityActionTemplateServiceWithEmptyClient(t *testing.T) {
	service := NewCommunityActionTemplateService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestCommunityActionTemplateServiceGetWithEmptyID(t *testing.T) {
	service := NewCommunityActionTemplateService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
