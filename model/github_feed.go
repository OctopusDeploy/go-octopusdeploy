package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewGitHubFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewGitHubFeed", "name")
	}

	feed := &Feed{
		FeedType: enum.GitHub,
		Name:     name,
	}

	return feed, nil
}
