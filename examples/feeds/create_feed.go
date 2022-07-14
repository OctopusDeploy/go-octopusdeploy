package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/feeds"
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

	client, err := client.NewClient(nil, apiURL, apiKey, spaceName)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	nuGetFeed, err := feeds.NewNuGetFeed(feedName, feedURI)
	if err != nil {
		_ = fmt.Errorf("error creating NuGet feed: %v", err)
		return
	}

	nuGetFeed.EnhancedMode = useExtendedAPI

	if len(feedUsername) > 0 {
		nuGetFeed.Username = feedUsername
	}
	if len(feedPassword) > 0 {
		nuGetFeed.Password = core.NewSensitiveValue(feedPassword)
	}

	_, err = client.Feeds.Add(nuGetFeed)
	if err != nil {
		_ = fmt.Errorf("error creating feed: %v", err)
		return
	}
}
