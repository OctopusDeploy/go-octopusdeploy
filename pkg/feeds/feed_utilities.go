package feeds

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
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
		awsElasticContainerRegistry, err := NewAwsElasticContainerRegistry(feedResource.GetName(), feedResource.AccessKey, feedResource.SecretKey, feedResource.Region, feedResource.ElasticContainerRegistryOidcAuthentication)
		if err != nil {
			return nil, err
		}
		feed = awsElasticContainerRegistry
	case FeedTypeAzureContainerRegistry:
		azureContainerRegistry, err := NewAzureContainerRegistry(feedResource.GetName(), feedResource.GetUsername(), feedResource.GetPassword(), feedResource.AzureContainerRegistryOidcAuthentication)
		if err != nil {
			return nil, err
		}
		azureContainerRegistry.APIVersion = feedResource.APIVersion
		azureContainerRegistry.FeedURI = feedResource.FeedURI
		azureContainerRegistry.RegistryPath = feedResource.RegistryPath
		feed = azureContainerRegistry
	case FeedTypeBuiltIn:
		builtInFeed, err := NewBuiltInFeed(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		builtInFeed.DeletePackagesAssociatedWithReleases = feedResource.DeletePackagesAssociatedWithReleases
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
	case FeedTypeGoogleContainerRegistry:
		googleContainerRegistry, err := NewGoogleContainerRegistry(feedResource.GetName(), feedResource.GetUsername(), feedResource.GetPassword(), feedResource.GoogleContainerRegistryOidcAuthentication)
		if err != nil {
			return nil, err
		}
		googleContainerRegistry.APIVersion = feedResource.APIVersion
		googleContainerRegistry.FeedURI = feedResource.FeedURI
		googleContainerRegistry.RegistryPath = feedResource.RegistryPath
		feed = googleContainerRegistry
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
	case FeedTypeArtifactoryGeneric:
		artifactoryGenericFeed, err := NewArtifactoryGenericFeed(feedResource.GetName())
		artifactoryGenericFeed.LayoutRegex = feedResource.LayoutRegex
		artifactoryGenericFeed.Repository = feedResource.Repository
		artifactoryGenericFeed.FeedURI = feedResource.FeedURI
		if err != nil {
			return nil, err
		}
		feed = artifactoryGenericFeed
	case FeedTypeS3:
		s3Feed, err := NewS3Feed(feedResource.GetName(), feedResource.AccessKey, feedResource.SecretKey, feedResource.UseMachineCredentials)
		if err != nil {
			return nil, err
		}
		feed = s3Feed
	case FeedTypeOCIRegistry:
		ociFeed, err := NewOCIRegistryFeed(feedResource.GetName())
		if err != nil {
			return nil, err
		}
		ociFeed.FeedURI = feedResource.FeedURI
		feed = ociFeed
	default:
		return nil, errors.New("unknown feed type: " + fmt.Sprint(feedResource.GetFeedType()))
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

func ToFeeds(feedResources *resources.Resources[*FeedResource]) *Feeds {
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
		if awsElasticContainerRegistry.OidcAuthentication != nil {
			feedResource.ElasticContainerRegistryOidcAuthentication = &AwsElasticContainerRegistryOidcAuthentication{
				SessionDuration: awsElasticContainerRegistry.OidcAuthentication.SessionDuration,
				Audience:        awsElasticContainerRegistry.OidcAuthentication.Audience,
				SubjectKeys:     awsElasticContainerRegistry.OidcAuthentication.SubjectKeys,
				RoleArn:         awsElasticContainerRegistry.OidcAuthentication.RoleArn,
			}
		}
	case FeedTypeAzureContainerRegistry:
		azureContainerRegistry := feed.(*AzureContainerRegistry)
		feedResource.APIVersion = azureContainerRegistry.APIVersion
		feedResource.FeedURI = azureContainerRegistry.FeedURI
		feedResource.RegistryPath = azureContainerRegistry.RegistryPath
		if azureContainerRegistry.OidcAuthentication != nil {
			feedResource.AzureContainerRegistryOidcAuthentication = &AzureContainerRegistryOidcAuthentication{
				ClientId:    azureContainerRegistry.OidcAuthentication.ClientId,
				TenantId:    azureContainerRegistry.OidcAuthentication.TenantId,
				Audience:    azureContainerRegistry.OidcAuthentication.Audience,
				SubjectKeys: azureContainerRegistry.OidcAuthentication.SubjectKeys,
			}
		}
	case FeedTypeBuiltIn:
		builtInFeed := feed.(*BuiltInFeed)
		feedResource.DeletePackagesAssociatedWithReleases = builtInFeed.DeletePackagesAssociatedWithReleases
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
	case FeedTypeGoogleContainerRegistry:
		googleContainerRegistry := feed.(*AzureContainerRegistry)
		feedResource.APIVersion = googleContainerRegistry.APIVersion
		feedResource.FeedURI = googleContainerRegistry.FeedURI
		feedResource.RegistryPath = googleContainerRegistry.RegistryPath
		if googleContainerRegistry.OidcAuthentication != nil {
			feedResource.AzureContainerRegistryOidcAuthentication = &AzureContainerRegistryOidcAuthentication{
				Audience:    googleContainerRegistry.OidcAuthentication.Audience,
				SubjectKeys: googleContainerRegistry.OidcAuthentication.SubjectKeys,
			}
		}
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
	case FeedTypeArtifactoryGeneric:
		artifactoryGenericFeed := feed.(*ArtifactoryGenericFeed)
		feedResource.Repository = artifactoryGenericFeed.Repository
		feedResource.LayoutRegex = artifactoryGenericFeed.LayoutRegex
		feedResource.FeedURI = artifactoryGenericFeed.FeedURI
	case FeedTypeS3:
		s3Feed := feed.(*S3Feed)
		feedResource.AccessKey = s3Feed.AccessKey
		feedResource.SecretKey = s3Feed.SecretKey
		feedResource.UseMachineCredentials = s3Feed.UseMachineCredentials
	case FeedTypeOCIRegistry:
		ociFeed := feed.(*OCIRegistryFeed)
		feedResource.FeedURI = ociFeed.FeedURI
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
