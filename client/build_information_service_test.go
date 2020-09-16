package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewBuildInformationServiceWithNil(t *testing.T) {
	service := NewBuildInformationService(nil, "")
	assert.Nil(t, service)
}

func TestBuildInformationServiceWithEmptyClient(t *testing.T) {
	service := NewBuildInformationService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}
