package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
)

func ToFeed(feedResource *FeedResource) (IFeed, error) {
	if IsNil(feedResource) {
		return nil, internal.CreateInvalidParameterError("ToFeed", "feedResource")
	}

	if err := feedResource.Validate(); err != nil {
		return nil, err
	}

	var feed IFeed

	switch feedResource.GetFeedType() {
	case FeedTypeAwsElasticContainerRegistry:
		awsElasticContainerRegistry, err := NewAwsElasticContainerRegistry(feedResource.GetName(), feedResource.AccessKey, feedResource.SecretKey, feedResource.Region)
		if err != nil {
			return nil, err
		}
		feed = awsElasticContainerRegistry
	case FeedTypeBuiltIn:
		builtInFeed, err := NewBuiltInFeed(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		builtInFeed.DeleteUnreleasedPackagesAfterDays = feedResource.DeleteUnreleasedPackagesAfterDays
		builtInFeed.DownloadAttempts = feedResource.DownloadAttempts
		builtInFeed.DownloadRetryBackoffSeconds = feedResource.DownloadRetryBackoffSeconds
		builtInFeed.IsBuiltInRepoSyncEnabled = feedResource.IsBuiltInRepoSyncEnabled
		feed = builtInFeed
	case FeedTypeDocker:
		dockerContainerRegistry, err := NewDockerContainerRegistry(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		dockerContainerRegistry.APIVersion = feedResource.APIVersion
		dockerContainerRegistry.FeedURI = feedResource.FeedURI
		dockerContainerRegistry.RegistryPath = feedResource.RegistryPath
		feed = dockerContainerRegistry
	case FeedTypeGitHub:
		gitHubRepositoryFeed, err := NewGitHubRepositoryFeed(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		gitHubRepositoryFeed.DownloadAttempts = feedResource.DownloadAttempts
		gitHubRepositoryFeed.DownloadRetryBackoffSeconds = feedResource.DownloadRetryBackoffSeconds
		gitHubRepositoryFeed.FeedURI = feedResource.FeedURI
		feed = gitHubRepositoryFeed
	case FeedTypeHelm:
		helmFeed, err := NewHelmFeed(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		helmFeed.FeedURI = feedResource.FeedURI
		feed = helmFeed
	case FeedTypeMaven:
		mavenFeed, err := NewMavenFeed(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		mavenFeed.DownloadAttempts = feedResource.DownloadAttempts
		mavenFeed.DownloadRetryBackoffSeconds = feedResource.DownloadRetryBackoffSeconds
		mavenFeed.FeedURI = feedResource.FeedURI
		feed = mavenFeed
	case FeedTypeNuGet:
		nuGetFeed, err := NewNuGetFeed(feedResource.GetName(), feedResource.FeedURI)
		if err != nil {
			return nil, err
		}
		nuGetFeed.DownloadAttempts = feedResource.DownloadAttempts
		nuGetFeed.DownloadRetryBackoffSeconds = feedResource.DownloadRetryBackoffSeconds
		nuGetFeed.EnhancedMode = feedResource.EnhancedMode
		nuGetFeed.FeedURI = feedResource.FeedURI
		feed = nuGetFeed
	case FeedTypeOctopusProject:
		octopusProjectFeed, err := NewOctopusProjectFeed(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		feed = octopusProjectFeed
	}

	feed.SetID(feedResource.GetID())
	feed.SetLinks(feedResource.GetLinks())
	feed.SetModifiedBy(feedResource.GetModifiedBy())
	feed.SetModifiedOn(feedResource.GetModifiedOn())
	feed.SetPackageAcquisitionLocationOptions(feedResource.GetPackageAcquisitionLocationOptions())
	feed.SetPassword(feedResource.GetPassword())
	feed.SetSpaceID(feedResource.GetSpaceID())
	feed.SetUsername(feedResource.GetUsername())

	return feed, nil
}

func ToFeeds(feedResources *FeedResources) *Feeds {
	return &Feeds{
		Items:        ToFeedArray(feedResources.Items),
		PagedResults: feedResources.PagedResults,
	}
}

func ToFeedResource(feed IFeed) (*FeedResource, error) {
	if IsNil(feed) {
		return nil, internal.CreateInvalidParameterError("ToFeedResource", "feed")
	}

	// conversion unnecessary if input feed is *FeedResource
	if v, ok := feed.(*FeedResource); ok {
		return v, nil
	}

	feedResource := NewFeedResource(feed.GetName(), feed.GetFeedType())

	switch feedResource.GetFeedType() {
	case FeedTypeAwsElasticContainerRegistry:
		awsElasticContainerRegistry := feed.(*AwsElasticContainerRegistry)
		feedResource.AccessKey = awsElasticContainerRegistry.AccessKey
		feedResource.Region = awsElasticContainerRegistry.Region
		feedResource.SecretKey = awsElasticContainerRegistry.SecretKey
	case FeedTypeBuiltIn:
		builtInFeed := feed.(*BuiltInFeed)
		feedResource.DeleteUnreleasedPackagesAfterDays = builtInFeed.DeleteUnreleasedPackagesAfterDays
		feedResource.DownloadAttempts = builtInFeed.DownloadAttempts
		feedResource.DownloadRetryBackoffSeconds = builtInFeed.DownloadRetryBackoffSeconds
		feedResource.IsBuiltInRepoSyncEnabled = builtInFeed.IsBuiltInRepoSyncEnabled
	case FeedTypeDocker:
		dockerContainerRegistry := feed.(*DockerContainerRegistry)
		feedResource.APIVersion = dockerContainerRegistry.APIVersion
		feedResource.FeedURI = dockerContainerRegistry.FeedURI
		feedResource.RegistryPath = dockerContainerRegistry.RegistryPath
	case FeedTypeGitHub:
		gitHubRepositoryFeed := feed.(*GitHubRepositoryFeed)
		feedResource.DownloadAttempts = gitHubRepositoryFeed.DownloadAttempts
		feedResource.DownloadRetryBackoffSeconds = gitHubRepositoryFeed.DownloadRetryBackoffSeconds
		feedResource.FeedURI = gitHubRepositoryFeed.FeedURI
	case FeedTypeHelm:
		helmFeed := feed.(*HelmFeed)
		feedResource.FeedURI = helmFeed.FeedURI
	case FeedTypeMaven:
		mavenFeed := feed.(*MavenFeed)
		feedResource.DownloadAttempts = mavenFeed.DownloadAttempts
		feedResource.DownloadRetryBackoffSeconds = mavenFeed.DownloadRetryBackoffSeconds
		feedResource.FeedURI = mavenFeed.FeedURI
	case FeedTypeNuGet:
		nuGetFeed := feed.(*NuGetFeed)
		feedResource.DownloadAttempts = nuGetFeed.DownloadAttempts
		feedResource.DownloadRetryBackoffSeconds = nuGetFeed.DownloadRetryBackoffSeconds
		feedResource.EnhancedMode = nuGetFeed.EnhancedMode
		feedResource.FeedURI = nuGetFeed.FeedURI
	case FeedTypeOctopusProject:
		// nothing to copy
	}

	feedResource.SetID(feed.GetID())
	feedResource.SetLinks(feed.GetLinks())
	feedResource.SetModifiedBy(feed.GetModifiedBy())
	feedResource.SetModifiedOn(feed.GetModifiedOn())
	feedResource.SetPackageAcquisitionLocationOptions(feed.GetPackageAcquisitionLocationOptions())
	feedResource.SetPassword(feed.GetPassword())
	feedResource.SetSpaceID(feed.GetSpaceID())
	feedResource.SetUsername(feed.GetUsername())

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
