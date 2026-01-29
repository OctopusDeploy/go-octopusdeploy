package feeds

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
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

func TestNpm(t *testing.T) {
	feedResource := FeedResource{
		AccessKey:                         "",
		APIVersion:                        "test",
		DeleteUnreleasedPackagesAfterDays: 10,
		DownloadAttempts:                  5,
		DownloadRetryBackoffSeconds:       10,
		EnhancedMode:                      false,
		FeedType:                          FeedTypeNpm,
		FeedURI:                           "https://registry.npmjs.org",
		IsBuiltInRepoSyncEnabled:          false,
		Name:                              "NPM Feed",
		Password:                          nil,
		PackageAcquisitionLocationOptions: nil,
		Region:                            "",
		RegistryPath:                      "",
		SecretKey:                         nil,
		SpaceID:                           "",
		Username:                          "testuser",
		LayoutRegex:                       "",
		Repository:                        "",
		UseMachineCredentials:             false,
		Resource:                          resources.Resource{},
	}

	feed, err := ToFeed(&feedResource)

	if err != nil {
		t.Fatalf("Error should not have been returned")
	}

	typedFeed := feed.(*NpmFeed)

	if typedFeed.Name != "NPM Feed" {
		t.Fatalf("Name does not match")
	}

	if typedFeed.DownloadAttempts != 5 {
		t.Fatalf("DownloadAttempts does not match")
	}

	if typedFeed.DownloadRetryBackoffSeconds != 10 {
		t.Fatalf("DownloadRetryBackoffSeconds does not match")
	}

	if typedFeed.FeedURI != "https://registry.npmjs.org" {
		t.Fatalf("FeedURI does not match")
	}

	if typedFeed.GetUsername() != "testuser" {
		t.Fatalf("Username does not match")
	}
}

func TestNpmToResource(t *testing.T) {
	feed := NpmFeed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		FeedURI:                     "https://registry.npmjs.org",
		feed:                        *newFeed("NPM Feed", FeedTypeNpm),
	}
	feed.SetUsername("testuser")

	feedResource, err := ToFeedResource(&feed)

	if err != nil {
		t.Fatalf("Error should not have been returned. %s", err)
	}

	if feedResource.FeedType != FeedTypeNpm {
		t.Fatalf("FeedType does not match")
	}

	if feedResource.Name != "NPM Feed" {
		t.Fatalf("Name does not match")
	}

	if feedResource.FeedURI != "https://registry.npmjs.org" {
		t.Fatalf("FeedURI does not match")
	}

	if feedResource.DownloadAttempts != 5 {
		t.Fatalf("DownloadAttempts does not match")
	}

	if feedResource.DownloadRetryBackoffSeconds != 10 {
		t.Fatalf("DownloadRetryBackoffSeconds does not match")
	}

	if feedResource.Username != "testuser" {
		t.Fatalf("Username does not match")
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

func TestOCIRegistryToResource(t *testing.T) {
	feed := OCIRegistryFeed{
		FeedURI: "oci://registry-2.docker.io",
		feed:    *newFeed("Test Registry 2", FeedTypeOCIRegistry),
	}

	feedResource, err := ToFeedResource(&feed)

	if err != nil {
		t.Fatalf("Error should not have been returned. %s", err)
	}

	if feedResource.FeedType != FeedTypeOCIRegistry {
		t.Fatalf("FeedType does not match")
	}

	if feedResource.Name != "Test Registry 2" {
		t.Fatalf("Name does not match")
	}

	if feedResource.FeedURI != "oci://registry-2.docker.io" {
		t.Fatalf("FeedURI does not match")
	}
}
