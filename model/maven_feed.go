package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewMavenFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewMavenFeed", "name")
	}

	feed := &Feed{
		FeedType: enum.Maven,
		Name:     name,
	}

	return feed, nil
}
