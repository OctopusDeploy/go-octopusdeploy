package client

import (
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createFeedService(t *testing.T) *feedService {
	service := newFeedService(nil, TestURIFeeds, TestURIBuiltInFeedStats)
	testNewService(t, service, TestURIFeeds, serviceFeedService)
	return service
}

func CreateTestAwsElasticContainerRegistry(t *testing.T, service *feedService) model.IFeed {
	if service == nil {
		service = createFeedService(t)
	}
	require.NotNil(t, service)

	// the feed service validates the input parameters and attempts to make a
	// connection to the Elastic Container Registry (ECR) -- therefore, a valid
	// set of credentials (i.e. access key, secret key) must be provided along
	// with a valid region (i.e. "ap-southeast-2")

	accessKey := "access-key"
	secretKey := model.NewSensitiveValue("secret-key")
	region := "ap-southeast-2"

	feed := model.NewAwsElasticContainerRegistry(getRandomName(), accessKey, secretKey, region)

	resource, err := service.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestGitHubRepositoryFeed(t *testing.T, service *feedService) model.IFeed {
	if service == nil {
		service = createFeedService(t)
	}
	require.NotNil(t, service)

	feed := model.NewGitHubRepositoryFeed(getRandomName(), "http://example.com/")

	resource, err := service.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestHelmFeed(t *testing.T, service *feedService) model.IFeed {
	if service == nil {
		service = createFeedService(t)
	}
	require.NotNil(t, service)

	feed := model.NewHelmFeed(getRandomName(), "http://example.com/")

	resource, err := service.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestMavenFeed(t *testing.T, service *feedService) model.IFeed {
	if service == nil {
		service = createFeedService(t)
	}
	require.NotNil(t, service)

	feed := model.NewMavenFeed(getRandomName(), "http://example.com/")

	resource, err := service.Add(feed)
	require.NoError(t, err)

	return resource
}

func CreateTestNuGetFeed(t *testing.T, service *feedService) model.IFeed {
	if service == nil {
		service = createFeedService(t)
	}
	require.NotNil(t, service)

	feed := model.NewNuGetFeed(getRandomName(), "http://example.com/")

	resource, err := service.Add(feed)
	require.NoError(t, err)

	return resource
}

func DeleteTestFeed(t *testing.T, service *feedService, feed model.IFeed) error {
	if service == nil {
		service = createFeedService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(feed.GetID())
}

func IsEqualFeeds(t *testing.T, expected model.IFeed, actual model.IFeed) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// TODO: compare remaining values
}

func TestFeedServiceAdd(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	feed, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(operationAdd, parameterFeed))
	assert.Nil(t, feed)

	feed, err = service.Add(&model.Feed{})
	require.Error(t, err)
	require.Nil(t, feed)

	// the following code is commented out due to the validation conducted by
	// the feed service

	// feed = CreateTestAwsElasticContainerRegistry(t, service)
	// require.NotNil(t, feed)
	// defer DeleteTestFeed(t, service, feed)

	feed = CreateTestGitHubRepositoryFeed(t, service)
	require.NotNil(t, feed)
	err = DeleteTestFeed(t, service, feed)
	require.NoError(t, err)

	feed = CreateTestHelmFeed(t, service)
	require.NotNil(t, feed)
	err = DeleteTestFeed(t, service, feed)
	require.NoError(t, err)

	feed = CreateTestMavenFeed(t, service)
	require.NotNil(t, feed)
	err = DeleteTestFeed(t, service, feed)
	require.NoError(t, err)

	feed = CreateTestNuGetFeed(t, service)
	require.NotNil(t, feed)
	err = DeleteTestFeed(t, service, feed)
	require.NoError(t, err)
}

func TestFeedServiceAddGetDelete(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	expected := CreateTestNuGetFeed(t, service)

	actual, err := service.GetByID(expected.GetID())
	require.NoError(t, err)
	IsEqualFeeds(t, expected, actual)

	err = DeleteTestFeed(t, service, expected)
	require.NoError(t, err)
}

func TestFeedServiceDelete(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	assert.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	assert.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	assert.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
}

func TestFeedServiceDeleteAll(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		if !strings.Contains(resource.GetID(), "builtin") {
			defer DeleteTestFeed(t, service, resource)
		}
	}
}

func TestFeedServiceGetAll(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	const count int = 32
	expected := map[string]model.IFeed{}
	for i := 0; i < count; i++ {
		feed := CreateTestNuGetFeed(t, service)
		defer DeleteTestFeed(t, service, feed)
		expected[feed.GetID()] = feed
	}

	feeds, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, feeds)
	require.GreaterOrEqual(t, len(feeds), count)

	for _, actual := range feeds {
		_, ok := expected[actual.GetID()]
		if ok {
			IsEqualFeeds(t, expected[actual.GetID()], actual)
		}
	}
}

func TestFeedServiceGetByID(t *testing.T) {
	service := createFeedService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	require.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	require.Nil(t, resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	require.Nil(t, resource)
}

func TestFeedServiceNew(t *testing.T) {
	serviceFunction := newFeedService
	client := &sling.Sling{}
	uriTemplate := emptyString
	builtInFeedStats := emptyString
	serviceName := serviceFeedService

	testCases := []struct {
		name             string
		f                func(*sling.Sling, string, string) *feedService
		client           *sling.Sling
		uriTemplate      string
		builtInFeedStats string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, builtInFeedStats},
		{"EmptyURITemplate", serviceFunction, client, emptyString, builtInFeedStats},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, builtInFeedStats},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.builtInFeedStats)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}
