package feeds

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"testing"
)

func TestUnexpectedFeed(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "test",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 0,
		DownloadAttempts:                  0,
		DownloadRetryBackoffSeconds:       0,
		EnhancedMode:                      false,
		FeedType:                          "UnexpectedFeedType",
		FeedURI:                           "",
		IsBuiltInRepoSyncEnabled:          false,
		Name:                              "",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	_, err := ToFeed(&feedResource)

	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}
}

func TestECR(t *testing.T) {
	secretKey := "SecretKey"

	feedResource := FeedResource{
		AccessKey:                         "AccessKey",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 0,
		DownloadAttempts:                  0,
		DownloadRetryBackoffSeconds:       0,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeAwsElasticContainerRegistry,
		FeedURI:                           "",
		IsBuiltInRepoSyncEnabled:          false,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "region",
		RegistryPath:                      "",
		SecretKey: &core.SensitiveValue{
			HasValue: true,
			Hint:     nil,
			NewValue: &secretKey,
		},
		SpaceID:               "",
		Username:              "",
		LayoutRegex:           "",
		Repository:            "",
		UseMachineCredentials: false,
		Resource:              resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*AwsElasticContainerRegistry)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.Region != "region" {
		t.Fatalf("Region does not match")
	}

	if typedFeed.AccessKey != feedResource.AccessKey {
		t.Fatalf("AccessKey does not match")
	}

	if *typedFeed.SecretKey.NewValue != secretKey {
		t.Fatalf("SecretKey does not match")
	}

}

func TestBuiltIn(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeBuiltIn,
		FeedURI:                           "",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*BuiltInFeed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.DeleteUnreleasedPackagesAfterDays != 10 {
		t.Fatalf("DeleteUnreleasedPackagesAfterDays does not match")
	}

	if typedFeed.DownloadAttempts != 5 {
		t.Fatalf("DownloadAttempts does not match")
	}

	if typedFeed.DownloadRetryBackoffSeconds != 3 {
		t.Fatalf("DownloadRetryBackoffSeconds does not match")
	}

	if !typedFeed.IsBuiltInRepoSyncEnabled {
		t.Fatalf("IsBuiltInRepoSyncEnabled does not match")
	}

}

func TestDocker(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "APIVersion",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeDocker,
		FeedURI:                           "http://example.com",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "RegistryPath",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*DockerContainerRegistry)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.FeedURI != "http://example.com" {
		t.Fatalf("FeedURI does not match")
	}

	if typedFeed.RegistryPath != "RegistryPath" {
		t.Fatalf("RegistryPath does not match")
	}

	if typedFeed.APIVersion != "APIVersion" {
		t.Fatalf("APIVersion does not match")
	}
}

func TestGitHub(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeGitHub,
		FeedURI:                           "http://example.com",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*GitHubRepositoryFeed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.DownloadAttempts != 5 {
		t.Fatalf("DownloadAttempts does not match")
	}

	if typedFeed.DownloadRetryBackoffSeconds != 3 {
		t.Fatalf("DownloadRetryBackoffSeconds does not match")
	}

	if typedFeed.FeedURI != "http://example.com" {
		t.Fatalf("FeedURI does not match")
	}
}

func TestHelm(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeHelm,
		FeedURI:                           "http://example.com",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*HelmFeed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.FeedURI != "http://example.com" {
		t.Fatalf("FeedURI does not match")
	}
}

func TestMaven(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeMaven,
		FeedURI:                           "http://example.com",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*MavenFeed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.DownloadAttempts != 5 {
		t.Fatalf("DownloadAttempts does not match")
	}

	if typedFeed.DownloadRetryBackoffSeconds != 3 {
		t.Fatalf("DownloadRetryBackoffSeconds does not match")
	}

	if typedFeed.FeedURI != "http://example.com" {
		t.Fatalf("FeedURI does not match")
	}
}

func TestNuget(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      true,
		FeedType:                          FeedTypeNuGet,
		FeedURI:                           "http://example.com",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*NuGetFeed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.DownloadAttempts != 5 {
		t.Fatalf("DownloadAttempts does not match")
	}

	if typedFeed.DownloadRetryBackoffSeconds != 3 {
		t.Fatalf("DownloadRetryBackoffSeconds does not match")
	}

	if !typedFeed.EnhancedMode {
		t.Fatalf("EnhancedMode does not match")
	}

	if typedFeed.FeedURI != "http://example.com" {
		t.Fatalf("FeedURI does not match")
	}

}

func TestOctopusProject(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeOctopusProject,
		FeedURI:                           "",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*OctopusProjectFeed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

}

func TestArtifactory(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeArtifactoryGeneric,
		FeedURI:                           "http://example.com",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "",
		LayoutRegex:                       "LayoutRegex",
		Repository:                        "Repository",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*ArtifactoryGenericFeed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.LayoutRegex != "LayoutRegex" {
		t.Fatalf("LayoutRegex does not match")
	}

	if typedFeed.Repository != "Repository" {
		t.Fatalf("Repository does not match")
	}

	if typedFeed.FeedURI != "http://example.com" {
		t.Fatalf("FeedURI does not match")
	}

}

func TestS3(t *testing.T) {
	secretKey := "SecretKey"

	feedResource := FeedResource{
		AccessKey:                         "AccessKey",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeS3,
		FeedURI:                           "",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "MyFeed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey: &core.SensitiveValue{
			HasValue: true,
			Hint:     nil,
			NewValue: &secretKey,
		},
		SpaceID:               "",
		Username:              "",
		LayoutRegex:           "",
		Repository:            "",
		UseMachineCredentials: true,
		Resource:              resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*S3Feed)

	if typedFeed.Name != "MyFeed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.AccessKey != "AccessKey" {
		t.Fatalf("AccessKey does not match")
	}

	if *typedFeed.SecretKey.NewValue != secretKey {
		t.Fatalf("SecretKey does not match")
	}

	if !typedFeed.UseMachineCredentials {
		t.Fatalf("UseMachineCredentials does not match")
	}
}

func TestOCIRegistry(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       3,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeOCIRegistry,
		FeedURI:                           "oci://test-registry.docker.io",
		IsBuiltInRepoSyncEnabled:          true,
		Name:                              "Test Registry",
		Password:                          core.NewSensitiveValue("test-password"),
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "test-username",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned. %s", err)
	}

	typedFeed := feed.(*OCIRegistryFeed)

	if typedFeed.Name != "Test Registry" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.FeedURI != "oci://test-registry.docker.io" {
		t.Fatalf("FeedURI does not match")
	}

	if typedFeed.Username != "test-username" {
		t.Fatalf("Username does not match")
	}
}
