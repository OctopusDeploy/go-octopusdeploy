package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewBuiltInFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewBuiltInFeed", "name")
	}

	feed := &Feed{
		FeedType: enum.BuiltIn,
		Name:     name,
	}

	return feed, nil
}
