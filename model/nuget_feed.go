package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewNuGetFeed(name string) (*Feed, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewNuGetFeed", "name")
	}

	feed := &Feed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		EnhancedMode:                false,
		FeedType:                    enum.NuGet,
		Name:                        name,
	}

	return feed, nil
}
