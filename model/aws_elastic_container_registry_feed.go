package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewAwsElasticContainerRegistryFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewAwsElasticContainerRegistryFeed", "name")
	}

	feed := &Feed{
		FeedType: enum.AwsElasticContainerRegistry,
		Name:     name,
	}

	return feed, nil
}
