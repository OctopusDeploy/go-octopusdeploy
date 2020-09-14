package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewDockerFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewDockerFeed", "name")
	}

	feed := &Feed{
		FeedType: enum.Docker,
		Name:     name,
	}

	return feed, nil
}
