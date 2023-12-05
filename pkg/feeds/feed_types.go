package feeds

type FeedType string

const (
	FeedTypeAwsElasticContainerRegistry = FeedType("AwsElasticContainerRegistry")
	FeedTypeBuiltIn                     = FeedType("BuiltIn")
	FeedTypeDocker                      = FeedType("Docker")
	FeedTypeGitHub                      = FeedType("GitHub")
	FeedTypeHelm                        = FeedType("Helm")
	FeedTypeMaven                       = FeedType("Maven")
	FeedTypeNuGet                       = FeedType("NuGet")
	FeedTypeOctopusProject              = FeedType("OctopusProject")
	FeedTypeArtifactoryGeneric          = FeedType("ArtifactoryGeneric")
)
