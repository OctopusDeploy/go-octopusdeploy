package feeds

type FeedType string

const (
	FeedTypeAwsElasticContainerRegistry = FeedType("AwsElasticContainerRegistry")
	FeedTypeAzureContainerRegistry      = FeedType("AzureContainerRegistry")
	FeedTypeBuiltIn                     = FeedType("BuiltIn")
	FeedTypeDocker                      = FeedType("Docker")
	FeedTypeGitHub                      = FeedType("GitHub")
	FeedTypeGoogleContainerRegistry     = FeedType("GoogleContainerRegistry")
	FeedTypeHelm                        = FeedType("Helm")
	FeedTypeMaven                       = FeedType("Maven")
	FeedTypeNuGet                       = FeedType("NuGet")
	FeedTypeOctopusProject              = FeedType("OctopusProject")
	FeedTypeArtifactoryGeneric          = FeedType("ArtifactoryGeneric")
	FeedTypeS3                          = FeedType("S3")
	FeedTypeOCIRegistry                 = FeedType("OciRegistry")
)
