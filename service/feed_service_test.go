package service

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createFeedService(t *testing.T) *feedService {
	service := newFeedService(nil, TestURIFeeds, TestURIBuiltInFeedStats)
	testNewService(t, service, TestURIFeeds, ServiceFeedService)
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

	err := service.DeleteByID(emptyString)
	assert.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)

	err = service.DeleteByID(whitespaceString)
	assert.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)
}

func TestFeedServiceGetByID(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)
}

func TestFeedServiceNew(t *testing.T) {
	ServiceFunction := newFeedService
	client := &sling.Sling{}
	uriTemplate := emptyString
	builtInFeedStats := emptyString
	ServiceName := ServiceFeedService

	testCases := []struct {
		name             string
		f                func(*sling.Sling, string, string) *feedService
		client           *sling.Sling
		uriTemplate      string
		builtInFeedStats string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, builtInFeedStats},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, builtInFeedStats},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, builtInFeedStats},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.builtInFeedStats)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
