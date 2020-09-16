package model

import "github.com/OctopusDeploy/go-octopusdeploy/enum"

func NewNuGetFeed(name string) *Feed {
	return &Feed{
		DownloadAttempts:            5,
		DownloadRetryBackoffSeconds: 10,
		EnhancedMode:                false,
		FeedType:                    enum.NuGet,
		Name:                        name,
	}
}
