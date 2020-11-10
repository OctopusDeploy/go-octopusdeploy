package octopusdeploy

import "github.com/jinzhu/copier"

func ToFeed(feedResource *FeedResource) (IFeed, error) {
	if isNil(feedResource) {
		return nil, createInvalidParameterError("ToFeed", "feedResource")
	}

	var feed IFeed
	var err error
	switch feedResource.GetFeedType() {
	case FeedTypeAwsElasticContainerRegistry:
		feed = NewAwsElasticContainerRegistry(feedResource.GetName(), feedResource.AccessKey, feedResource.SecretKey, feedResource.Region)
	case FeedTypeBuiltIn:
		feed = NewBuiltInFeed(feedResource.GetName(), feedResource.FeedURI)
	case FeedTypeDocker:
		feed = NewDockerContainerRegistry(feedResource.GetName())
	case FeedTypeGitHub:
		feed = NewGitHubRepositoryFeed(feedResource.GetName())
	case FeedTypeHelm:
		feed = NewHelmFeed(feedResource.GetName())
	case FeedTypeMaven:
		feed = NewMavenFeed(feedResource.GetName())
	case FeedTypeNuGet:
		feed = NewNuGetFeed(feedResource.GetName(), feedResource.FeedURI)
	case FeedTypeOctopusProject:
		feed = NewOctopusProjectFeed(feedResource.GetName(), feedResource.FeedURI)
	}

	err = copier.Copy(feed, feedResource)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

func ToFeeds(feedResources *FeedResources) *Feeds {
	return &Feeds{
		Items:        ToFeedArray(feedResources.Items),
		PagedResults: feedResources.PagedResults,
	}
}

func ToFeedResource(feed IFeed) (*FeedResource, error) {
	if isNil(feed) {
		return nil, createInvalidParameterError("ToFeedResource", ParameterFeed)
	}

	feedResource := NewFeedResource(feed.GetName(), feed.GetFeedType())

	err := copier.Copy(&feedResource, feed)
	if err != nil {
		return nil, err
	}

	return feedResource, nil
}

func ToFeedArray(feedResources []*FeedResource) []IFeed {
	items := []IFeed{}
	for _, feedResource := range feedResources {
		feed, err := ToFeed(feedResource)
		if err != nil {
			return nil
		}
		items = append(items, feed)
	}
	return items
}
