package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewHelmFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewHelmFeed", "name")
	}

	feed := &Feed{
		FeedType: enum.Helm,
		Name:     name,
	}

	return feed, nil
}
