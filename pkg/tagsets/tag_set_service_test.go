package tagsets

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTagSetService(t *testing.T) {
	ServiceFunction := NewTagSetService
	client := &sling.Sling{}
	uriTemplate := ""
	sortOrderPath := ""
	ServiceName := constants.ServiceTagSetService

	testCases := []struct {
		name          string
		f             func(*sling.Sling, string, string) *TagSetService
		client        *sling.Sling
		uriTemplate   string
		sortOrderPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, sortOrderPath},
		{"EmptyURITemplate", ServiceFunction, client, "", sortOrderPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", sortOrderPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.sortOrderPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestTagSetServiceGetWithEmptyID(t *testing.T) {
	service := NewTagSetService(nil, "", constants.TestURITagSetSortOrder)

	resource, err := service.GetByID("")
	require.Error(t, err)
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
