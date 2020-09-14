package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewOctopusProjectFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewOctopusProjectFeed", "name")
	}

	feed := &Feed{
		FeedType: enum.OctopusProject,
		Name:     name,
	}

	return feed, nil
}
