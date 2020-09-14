package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewNuGetFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewNuGetFeed", "name")
	}

	feed := &Feed{
		Name:     name,
		FeedType: enum.NuGetFeed,
	}

	return feed, nil

}
