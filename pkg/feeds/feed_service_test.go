package feeds

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createFeedService(t *testing.T) *FeedService {
	service := NewFeedService(nil, constants.TestURIFeeds, constants.TestURIBuiltInFeedStats)
	services.NewServiceTests(t, service, constants.TestURIFeeds, constants.ServiceFeedService)
	return service
}

func TestFeedServiceAdd(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	feed, err := service.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterFeed))
	assert.Nil(t, feed)

	feed, err = service.Add(&FeedResource{})
	require.Error(t, err)
	require.Nil(t, feed)
}

func TestFeedServiceDelete(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	err := service.DeleteByID("")
	assert.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)

	err = service.DeleteByID(" ")
	assert.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)
}

func TestFeedServiceGetByID(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, "id"), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, "id"), err)
	require.Nil(t, resource)
}

func TestFeedServiceNew(t *testing.T) {
	ServiceFunction := NewFeedService
	client := &sling.Sling{}
	uriTemplate := ""
	builtInFeedStats := ""
	ServiceName := constants.ServiceFeedService

	testCases := []struct {
		name             string
		f                func(*sling.Sling, string, string) *FeedService
		client           *sling.Sling
		uriTemplate      string
		builtInFeedStats string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, builtInFeedStats},
		{"EmptyURITemplate", ServiceFunction, client, "", builtInFeedStats},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", builtInFeedStats},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.builtInFeedStats)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
