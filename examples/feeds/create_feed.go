package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func CreateFeedExample() {
	var (
		apiKey         string = "API-YOUR_API_KEY"
		feedName       string = "nuget.org 3"
		feedPassword   string = "" // optional
		feedURI        string = "https://api.nuget.org/v3/index.json"
		feedUsername   string = "" // optional
		octopusURL     string = "https://your_octopus_url"
		spaceName      string = "Default"
		useExtendedAPI bool   = false
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	nuGetFeed := octopusdeploy.NewNuGetFeed(feedName, feedURI)
	nuGetFeed.EnhancedMode = useExtendedAPI

	if len(feedUsername) > 0 {
		nuGetFeed.Username = feedUsername
	}
	if len(feedPassword) > 0 {
		password := octopusdeploy.NewSensitiveValue(feedPassword)
		nuGetFeed.Password = &password
	}

	_, err = client.Feeds.Add(nuGetFeed)
	if err != nil {
		_ = fmt.Errorf("error creating feed: %v", err)
		return
	}
}
