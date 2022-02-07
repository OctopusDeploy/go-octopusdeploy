package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createFeedService(t *testing.T) *feedService {
	service := newFeedService(nil, TestURIFeeds, TestURIBuiltInFeedStats)
	services.testNewService(t, service, TestURIFeeds, ServiceFeedService)
	return service
}

func TestFeedServiceAdd(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	feed, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterFeed))
	assert.Nil(t, feed)

	feed, err = service.Add(&FeedResource{})
	require.Error(t, err)
	require.Nil(t, feed)
}

func TestFeedServiceDelete(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(services.emptyString)
	assert.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)

	err = service.DeleteByID(services.whitespaceString)
	assert.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)
}

func TestFeedServiceGetByID(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(services.emptyString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)
}

func TestFeedServiceNew(t *testing.T) {
	ServiceFunction := newFeedService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	builtInFeedStats := services.emptyString
	ServiceName := ServiceFeedService

	testCases := []struct {
		name             string
		f                func(*sling.Sling, string, string) *feedService
		client           *sling.Sling
		uriTemplate      string
		builtInFeedStats string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, builtInFeedStats},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, builtInFeedStats},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, builtInFeedStats},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.builtInFeedStats)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
