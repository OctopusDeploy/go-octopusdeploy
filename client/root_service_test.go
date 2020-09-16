package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewRootServiceWithNil(t *testing.T) {
	service := NewRootService(nil, "")
	assert.Nil(t, service)
}

func TestRootServiceWithEmptyClient(t *testing.T) {
	service := NewRootService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}
