package users

import (
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAPIKeyService(t *testing.T) {
	ServiceFunction := NewAPIKeyService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceAPIKeyService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *ApiKeyService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := createAPIKeyService(t)
	resource, err := service.GetByUserID("")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByUserID", "userID"))
	assert.Nil(t, resource)

	resource, err = service.GetByUserID(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByUserID", "userID"))
	assert.Nil(t, resource)
}

func createAPIKeyService(t *testing.T) *ApiKeyService {
	service := NewAPIKeyService(nil, constants.TestURIAPIKeys)
	services.NewServiceTests(t, service, constants.TestURIAPIKeys, constants.ServiceAPIKeyService)
	return service
}
