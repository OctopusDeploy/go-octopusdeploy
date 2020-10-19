package examples

import (
	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

func CreateFeedExample() {
	var (
		// Declare working variables
		octopusURL    string = "https://youroctourl"
		octopusAPIKey string = "API-YOURAPIKEY"

		spaceName      string = "Default"
		feedName       string = "nuget.org 3"
		feedURI        string = "https://api.nuget.org/v3/index.json"
		useExtendedAPI bool   = false
		// optional
		feedUsername string = ""
		feedPassword string = ""
	)

	client, err := client.NewClient(nil, octopusURL, octopusAPIKey, spaceName)
	if err != nil {
		// TODO: handle error
	}

	nuGetFeed := model.NewNuGetFeed(feedName, feedURI)
	if err != nil {
		// TODO: handle error
	}

	nuGetFeed.EnhancedMode = useExtendedAPI

	if len(feedUsername) > 0 {
		nuGetFeed.Username = feedUsername
	}
	if len(feedPassword) > 0 {
		password := model.NewSensitiveValue(feedPassword)
		nuGetFeed.Password = &password
	}

	_, err = client.Feeds.Add(nuGetFeed)
	if err != nil {
		// TODO: handle error
	}
}
