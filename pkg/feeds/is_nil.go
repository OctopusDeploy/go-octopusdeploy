package feeds

func IsNil(i interface{}) bool {
	switch v := i.(type) {
	case *AwsElasticContainerRegistry:
		return v == nil
	case *BuiltInFeed:
		return v == nil
	case *DockerContainerRegistry:
		return v == nil
	case *feed:
		return v == nil
	case *FeedResource:
		return v == nil
	case *GitHubRepositoryFeed:
		return v == nil
	case *HelmFeed:
		return v == nil
	case *MavenFeed:
		return v == nil
	case *NuGetFeed:
		return v == nil
	case *OctopusProjectFeed:
		return v == nil
	default:
		return v == nil
	}
}
